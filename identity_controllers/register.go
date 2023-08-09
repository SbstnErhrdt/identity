package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/SbstnErhrdt/identity/security"
	"github.com/google/uuid"
	"net/mail"
	"time"
)

// ErrRegistrationConfirmationExpired is returned when the registration confirmation has expired
var ErrRegistrationConfirmationExpired = errors.New("registration confirmation expired. Please re-register")

// ErrAcceptTermsAndConditions is returned when the user did not accept the terms and conditions
var ErrAcceptTermsAndConditions = errors.New("please accept the terms and conditions")

// ErrGenericRegistration message
var ErrGenericRegistration = errors.New("there was a problem during the registration process. Please try again")

// ErrRegistrationIsNotAllowed is returned when the system does not allow registration
var ErrRegistrationIsNotAllowed = errors.New("registration is not allowed. Please contact the administrator")

// Register registers a new user
func Register(service IdentityService, emailAddress, password string, termAndConditions bool, userAgent, ip, origin string) (err error) {
	// sanitize input
	emailAddress = SanitizeEmail(emailAddress)
	// init logger
	logger := service.GetLogger().With(
		"email", emailAddress,
		"process", "Register",
	)
	// checks if users can register
	if !service.AllowRegistration(origin) {
		err = ErrRegistrationIsNotAllowed
		logger.With("err", err).Warn("registration is not allowed")
		return
	}
	// check login data
	if len(emailAddress) == 0 {
		err = ErrNoEmail
		logger.With("err", err).Error("no email address")
		return
	}
	if len(password) == 0 {
		err = ErrNoPassword
		logger.With("err", err).Error("no password")
		return
	}
	// check terms and conditions
	if !termAndConditions {
		err = ErrAcceptTermsAndConditions
		return
	}
	// check if user exits already
	logger.Debug("CheckIfEmailIsFree")
	isFree, err := CheckIfEmailIsFree(service, emailAddress)
	if !isFree {
		err = ErrEmailAlreadyExists
		logger.With("err", err).Warn("email already exists")
		// OWASP recommends not to return an in case the account already exists
		return nil
	}
	if err != nil {
		logger.With("err", err).Error("cannot check if email is free")
		// OWASP recommends not to return an in case the account already exists
		return nil
	}
	// check password complexity
	logger.Debug("CheckPasswordComplexity")
	err = security.CheckPasswordComplexity(password)
	if err != nil {
		logger.With("err", err).Error("password does not meet the complexity requirements")
		return
	}
	// generate random registration token
	logger.Debug("GenerateRandomString")
	randomToken, err := security.GenerateRandomString(64)
	if err != nil {
		logger.With("err", err).Error("could not generate random string")
		return
	}
	// create user
	identity := identity_models.Identity{
		Email:                      emailAddress,
		AcceptConditionsAndPrivacy: termAndConditions,
	}
	identity.UID = uuid.New()
	// set password
	logger.Debug("SetNewPassword")
	err = identity.SetNewPassword(service.GetPepper(), password)
	if err != nil {
		logger.With("err", err).Error("could not set new password")
		err = ErrGenericRegistration
		return
	}
	// Is the user is automatically cleared after the registration
	if service.AutoClearUserAfterRegistration(origin) {
		// user can directly log in after the registration confirmation
		identity.Cleared = true
	} else {
		// user needs to be cleared manually by an administrator
		identity.Cleared = false
	}

	// init transaction
	tx := service.GetSQLClient().Begin()
	// create user
	logger.Debug("create identity in the database")
	tx.Create(&identity)
	// create confirmation
	confirmation := identity_models.IdentityRegistrationConfirmation{
		Token:                 randomToken,
		IdentityUID:           identity.UID,
		ExpiredAt:             time.Now().UTC().Add(time.Hour * 12),
		RegistrationIP:        ip,
		RegistrationUserAgent: userAgent,
	}
	logger.Debug("create confirmation token in the database")
	tx.Create(&confirmation)
	err = tx.Commit().Error
	if err != nil {
		logger.With("err", err).Error("")
		return
	}
	// Build email template
	logger.Debug("Build email template")
	// resolve the email template
	emailTemplate := service.ResolveRegistrationEmailTemplate(origin, emailAddress, randomToken)
	// generate the content of the email
	content, err := emailTemplate.Content()
	if err != nil {
		logger.With("err", err).Error("")
		return
	}
	// Send email
	logger.Debug("Send registration email")
	err = service.SendEmail(
		service.GetSenderEmailAddress(),
		mail.Address{
			Name:    emailAddress,
			Address: emailAddress,
		},
		// todo: make subject dynamic based on mandate or client
		"Registration Confirmation",
		content)
	return
}

var ErrExternalNoRegistrationToken = errors.New("no registration token found in the database. Please re-register")

// RegistrationConfirmation confirms a registration
func RegistrationConfirmation(service IdentityService, token, userAgent, ip string) (err error) {
	logger := service.GetLogger().With(
		"token", token,
		"process", "ConfirmEmail",
	)
	// check if token is in table
	tokenIdentityResult := identity_models.IdentityRegistrationConfirmation{}
	logger.Debug("find registration confirmation")
	err = service.
		GetSQLClient().
		Where("token = ?", token).
		First(&tokenIdentityResult).Error
	if err != nil {
		logger.With("err", err).Error("")
		err = ErrExternalNoRegistrationToken
		return
	}
	// check if identity was already confirmed
	if tokenIdentityResult.ConfirmedAt != nil {
		logger.Warn("already confirmed")
		return nil
	}
	// check if confirmation is already expired
	// if the identity is expired, delete the identity
	if time.Now().UTC().After(tokenIdentityResult.ExpiredAt) {
		logger.Warn("confirmation expired")
		identity, errIdentity := GetIdentityByUID(service, tokenIdentityResult.IdentityUID)
		if errIdentity != nil {
			err = errIdentity
			logger.With("err", err).Error("")
			return
		}
		// delete account
		errAccount := softDeleteAccount(service, identity)
		if errAccount != nil {
			err = errAccount
			logger.With("err", err).Error("")
			return
		}
		return ErrRegistrationConfirmationExpired
	}
	// get identity
	identity := identity_models.Identity{}
	logger.Debug("find identity")
	err = service.
		GetSQLClient().
		Where("uid = ?", tokenIdentityResult.IdentityUID).
		First(&identity).Error
	if err != nil {
		logger.With("err", err).Error("")
		return
	}
	// confirm identity
	// and activate user
	identity.Active = true
	// save identity
	logger.Debug("updated identity")
	err = service.
		GetSQLClient().
		Save(&identity).Error
	if err != nil {
		logger.With("err", err).Error("")
		return
	}
	// update tokenIdentityResult
	logger.Debug("update registration confirmation")
	now := time.Now().UTC()
	tokenIdentityResult.ConfirmedAt = &now
	tokenIdentityResult.ConfirmationUserAgent = userAgent
	tokenIdentityResult.ConfirmationIP = ip
	err = service.
		GetSQLClient().
		Save(&tokenIdentityResult).Error
	if err != nil {
		logger.With("err", err).Error("")
		return
	}
	return
}
