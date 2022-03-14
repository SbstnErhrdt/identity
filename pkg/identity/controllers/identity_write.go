package controllers

import (
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/pkg/identity/models"
	"github.com/SbstnErhrdt/identity/pkg/identity/security"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/mail"
	"reflect"
	"strings"
	"time"
)

var ErrNoEmail = errors.New("no email address")
var ErrNoPassword = errors.New("no password")
var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrRegistrationConfirmationExpired = errors.New("registration confirmation expired. Please re-register.")

// Register registers a new user
func Register(service IdentityService, emailAddress, password string, termAndConditions bool, userAgent, ip, origin string) (err error) {
	// sanitize input
	emailAddress = strings.ToLower(emailAddress)
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
		return err
	}
	if err != nil {
		logger.Error(err)
		return
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
	// todo clear by administrator
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

func Login(service IdentityService, emailAddress, password, userAgent, ip string) (token string, err error) {
	// sanitize input
	emailAddress = strings.ToLower(emailAddress)
	// init logger
	logger := log.WithFields(log.Fields{
		"identity": emailAddress,
		"process":  "Login",
	})
	// check login data
	if len(emailAddress) == 0 {
		logger.Error(ErrNoEmail)
		return "", ErrNoEmail
	}
	if len(password) == 0 {
		logger.Error(ErrNoPassword)
		return "", ErrNoPassword
	}
	// lowercase
	emailAddress = strings.ToLower(emailAddress)
	// track login attempt
	logger.Debug("track login attempt")
	loginAttempt := models.IdentityLogin{
		Email:     emailAddress,
		UserAgent: userAgent,
		IP:        ip,
	}
	err = service.GetSQLClient().Create(&loginAttempt).Error
	if err != nil {
		logger.Error(err)
		return "", err
	}
	// Check if identity exists
	// init identity object
	identity := &models.Identity{}

	// Find the identity in the database
	logger.Debug("find identity")
	err = service.GetSQLClient().
		Where("email = ?", emailAddress).
		First(&identity).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("identity does not exist")
		logger.Error(err)
		// emailAddress is free
		return "", err
	} else if err != nil {
		logger.Error(err)
		return "", err
	}

	// check if identity is not blocked
	logger.Debug("check if identity is not blocked")
	if identity.Blocked {
		err = errors.New("this identity can not login. please contact the support")
		logger.Error(err)
		return "", err
	}

	// check if identity is active
	logger.Debug("check if identity is activated")
	if !identity.Active {
		err = errors.New("this identity's email address is not confirmed yet")
		logger.Error(err)
		return "", err
	}

	// verify the password
	logger.Debug("verify password")
	if !(VerifyPassword(service, identity, password)) {
		err = errors.New("please check your password")
		logger.Error(err)
		return "", err
	}

	// generate token
	log.Debug("generate token")
	audience := env.FallbackEnvVariable("SECURITY_JWT_AUDIENCE", "API")
	token, tokenUID, errToken := security.GenerateJWTToken(identity.UID, audience, map[string]interface{}{})
	if errToken != nil {
		logger.Error(errToken)
		return "", errToken
	}

	// save token id in token table
	tokenMeta := models.IdentityTokenMeta{
		TokenUID:  tokenUID,
		TokenType: models.LoginToken,
	}
	errTokenMet := service.GetSQLClient().Create(&tokenMeta).Error
	if errToken != nil {
		logger.Error(errTokenMet)
		return "", errTokenMet
	}

	// update login attempt
	logger.Debug("update login attempt")
	loginAttempt.IdentityUID = identity.UID
	service.GetSQLClient().Save(&loginAttempt)

	return token, nil
}

func Logout(service IdentityService, token string) (err error) {
	// add the token uuid to the expired database
	return
}

func LogoutAllDevices(service IdentityService, user *models.Identity) (err error) {
	return
}

func InitResetPassword(service IdentityService, emailAddress, origin string) (err error) {
	// init logger
	logger := log.WithFields(log.Fields{
		"email":   emailAddress,
		"process": "InitResetPassword",
	})
	// get identity
	res, err := GetIdentityByEmail(service, emailAddress)
	if err != nil {
		logger.Error(err)
		return err
	}
	// create database entry
	token, err := security.GenerateRandomString(32)
	if err != nil {
		logger.Error(err)
		return err
	}
	resetPassword := models.IdentityResetPassword{
		IdentityUID: res.UID,
		Token:       token,
		Expire:      time.Now().UTC().Add(time.Hour * 24),
	}
	// save in database
	err = service.GetSQLClient().Create(&resetPassword).Error
	if err != nil {
		logger.Error(err)
		return err
	}
	// gen random token
	randomToken, err := security.GenerateRandomString(64)
	if err != nil {
		logger.Error(err)
		return
	}
	// send email
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

func ConfirmResetPassword(service IdentityService, token, password string) (err error) {
	return
}

func InitChangeEmail(service IdentityService, username, password string) (user *models.Identity, err error) {
	return
}

func ConfirmOldEmail(service IdentityService, token string) {
}

func ConfirmNewEmail(service IdentityService, token string) {
}

func DeleteAccount(service IdentityService, username, password string) (user *models.Identity, err error) {
	return
}

func deleteAccount(service IdentityService, identity *models.Identity) (err error) {
	// delete identity
	err = service.GetSQLClient().Delete(&identity).Error
	return
}

func AnonymizeAccount(service IdentityService, user *models.Identity) (err error) {
	// todo: anonymize user
	return
}

// SetNewPassword sets a new password for the user
func SetNewPassword(service IdentityService, user *models.Identity, password string) {
	// Hash pw and salt
	pw, salt := security.HashPassword(env.FallbackEnvVariable("SECURITY_PEPPER", "PEPPER"), password, []byte{})
	user.Password = pw
	user.Salt = salt
	return
}

// Clear clears the user
func Clear(service IdentityService, user *models.Identity) {
	user.Cleared = true
	service.GetSQLClient().Save(user)
	return
}

// Block blocks a user
func Block(service IdentityService, user *models.Identity) {
	user.Blocked = true
	service.GetSQLClient().Save(user)
	return
}

// VerifyPassword verifies the user's password given the user object and the password
func VerifyPassword(service IdentityService, user *models.Identity, password string) bool {
	checkPassword, _ := security.HashPassword(
		env.FallbackEnvVariable("SECURITY_PEPPER", "PEPPER"),
		password,
		user.Salt)
	return reflect.DeepEqual(user.Password, checkPassword)
}

// GenerateJWT generates a Json Web Token from the user object
func GenerateJWT(service IdentityService, user *models.Identity) (result string, err error) {
	// Init the token structure
	payload := map[string]interface{}{}
	payload["userUID"] = user.UID
	// Generate the token
	audience := env.FallbackEnvVariable("SECURITY_JWT_AUDIENCE", "API")
	result, _, err = security.GenerateJWTToken(user.UID, audience, payload)
	if err != nil {
		log.Error(err)
	}
	return
}

// CreatePasswordResetToken creates a password reset token
func CreatePasswordResetToken(service IdentityService, identity *models.Identity) (token string, err error) {
	// Init the token
	userToken := models.IdentityResetPassword{
		IdentityUID: identity.UID,
		Expire:      time.Now().Add(12 * time.Hour),
	}
	// Save the token
	err = service.GetSQLClient().Create(&userToken).Error
	if err != nil {
		log.Error(err)
		err = errors.Wrap(err, "token can not be generated")
		return
	}
	token, err = security.GeneratePasswordResetToken(identity.Email, time.Now().Add(12*time.Hour))
	if err != nil {
		log.Error(err)
	}
	return
}

// SetPasswordOfIdentity set the password of a user by its email
func SetPasswordOfIdentity(service IdentityService, email, newPassword string) (err error) {
	// Find user by name
	user, err := GetIdentityByEmail(service, email)
	if err != nil {
		err = errors.Wrap(err, "can not find user")
		return
	}
	// set password
	SetNewPassword(service, user, newPassword)
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		log.Error(err)
		err = errors.Wrap(err, "can not save user")
		return
	}
	return
}

func UpdateIdentity(service IdentityService, newIdentity *models.Identity) (err error) {
	// check if necessary fields are set
	if newIdentity.UID.String() == "" {
		return errors.New("UpdateIdentity: no uid provided")
	}
	// Get latest version from db
	dbObj, dbErr := GetIdentityByUID(service, newIdentity.UID)
	if dbErr != nil {
		log.Error(dbErr)
		err = errors.New("can not read identity from database")
		return
	}
	// Update data of database object
	dbObj.FirstName = newIdentity.FirstName
	dbObj.LastName = newIdentity.LastName
	dbObj.Phone = newIdentity.Phone
	dbObj.BackupPhone = newIdentity.BackupPhone
	dbObj.Email = strings.ToLower(newIdentity.Email)
	dbObj.BackupEmail = strings.ToLower(newIdentity.BackupEmail)
	dbObj.Salutation = newIdentity.Salutation
	// Set other metadata
	dbObj.UpdatedAt = time.Now().UTC()
	// create new entry
	err = service.GetSQLClient().Save(dbObj).Error
	if err != nil {
		log.Error(err)
		return
	}
	newIdentity = dbObj
	return
}
