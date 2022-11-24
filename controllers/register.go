package controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/models"
	"github.com/SbstnErhrdt/identity/security"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/mail"
	"time"
)

// ErrRegistrationConfirmationExpired is returned when the registration confirmation has expired
var ErrRegistrationConfirmationExpired = errors.New("registration confirmation expired. Please re-register")

// Register registers a new user
func Register(service IdentityService, emailAddress, password string, termAndConditions bool, userAgent, ip, origin string) (err error) {
	// sanitize input
	emailAddress = SanitizeEmail(emailAddress)
	// init logger
	logger := log.WithFields(log.Fields{
		"email":   emailAddress,
		"process": "Register",
	})
	// check login data
	if len(emailAddress) == 0 {
		logger.Error(ErrNoEmail)
		return ErrNoEmail
	}
	if len(password) == 0 {
		logger.Error(ErrNoPassword)
		return ErrNoPassword
	}
	// check terms and conditions
	if !termAndConditions {
		err = errors.New("please accept the terms and conditions")
		return
	}
	// check if user exits already
	logger.Debug("CheckIfEmailIsFree")
	isFree, err := CheckIfEmailIsFree(service, emailAddress)
	if !isFree {
		err = ErrEmailAlreadyExists
		logger.Warn(err)
		// OWASP recommends not to return an in case the account already exists
		return nil
	}
	if err != nil {
		logger.Error(err)
		// OWASP recommends not to return an in case the account already exists
		return nil
	}
	// check password complexity
	logger.Debug("CheckPasswordComplexity")
	err = security.CheckPasswordComplexity(password)
	if err != nil {
		logger.Error(err)
		return
	}
	// generate random registration token
	logger.Debug("GenerateRandomString")
	randomToken, err := security.GenerateRandomString(64)
	if err != nil {
		logger.Error(err)
		return
	}
	// create user
	identity := models.Identity{
		Email:                      emailAddress,
		AcceptConditionsAndPrivacy: termAndConditions,
	}
	identity.UID = uuid.New()
	// set password
	logger.Debug("SetNewPassword")
	SetNewPassword(service, &identity, password)
	// init transaction
	tx := service.GetSQLClient().Begin()
	// create user
	logger.Debug("create identity in the database")
	tx.Create(&identity)
	// create confirmation
	confirmation := models.IdentityRegistrationConfirmation{
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
		logger.Error(err)
		return
	}
	// Build email template
	logger.Debug("Build email template")
	// resolve the email template
	emailTemplate := service.ResolveRegistrationEmailTemplate(origin, emailAddress, randomToken)
	// generate the content of the email
	content, err := emailTemplate.Content()
	if err != nil {
		logger.Error(err)
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
		"Registration",
		content)
	return
}

// RegistrationConfirmation confirms a registration
func RegistrationConfirmation(service IdentityService, token, userAgent, ip string) (err error) {
	logger := log.WithFields(log.Fields{
		"token":   token,
		"process": "ConfirmEmail",
	})
	// check if token is in table
	tokenIdentityResult := models.IdentityRegistrationConfirmation{}
	logger.Debug("find registration confirmation")
	err = service.
		GetSQLClient().
		Where("token = ?", token).
		First(&tokenIdentityResult).Error
	if err != nil {
		logger.Error(err)
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
			logger.Error(err)
			return
		}
		// delete account
		errAccount := deleteAccount(service, identity)
		if errAccount != nil {
			err = errAccount
			logger.Error(err)
			return
		}
		return ErrRegistrationConfirmationExpired
	}
	// get identity
	identity := models.Identity{}
	logger.Debug("find identity")
	err = service.
		GetSQLClient().
		Where("uid = ?", tokenIdentityResult.IdentityUID).
		First(&identity).Error
	if err != nil {
		logger.Error(err)
		return
	}
	// confirm identity
	// and activate user
	identity.Active = true
	// todo clear by administrator+
	// the idea here is, that the user can be activated by the administrator
	// or, depending on the initial setting, can be activated directly
	identity.Cleared = true
	// save identity
	logger.Debug("updated identity")
	err = service.
		GetSQLClient().
		Save(&identity).Error
	if err != nil {
		logger.Error(err)
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
		logger.Error(err)
		return
	}
	return
}