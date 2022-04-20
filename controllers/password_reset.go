package controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/models"
	"github.com/SbstnErhrdt/identity/security"
	log "github.com/sirupsen/logrus"
	"net/mail"
	"time"
)

// InitResetPassword inits the password reset process
func InitResetPassword(service IdentityService, emailAddress, userAgent, ip, origin string) (err error) {
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
		UserAgent:   userAgent,
		IP:          ip,
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

// ErrTokenExpired is returned when the token is expired
var ErrTokenExpired = errors.New("security token expired. Please request a new password reset")

// ErrTokenUsed is returned when the token was used
var ErrTokenUsed = errors.New("security token already used. Please request a new password reset")

// ResetPassword resets the password
func ResetPassword(service IdentityService, token string, newPassword string, userAgent, ip, origin string) (err error) {
	// init logger
	logger := log.WithFields(log.Fields{
		"token":   token,
		"process": "InitResetPassword",
	})

	// get the token from the database
	resetPassword := models.IdentityResetPassword{}
	err = service.GetSQLClient().Where("token = ?", token).First(&resetPassword).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	// check if expiration date is in the past
	if resetPassword.Expire.Before(time.Now().UTC()) {
		err = ErrTokenExpired
		logger.Error(err)
		return
	}
	// check if the token was used
	if resetPassword.ConfirmationTime.IsZero() {
		err = ErrTokenUsed
		logger.Error(err)
		return
	}

	// check password complexity
	logger.Debug("CheckPasswordComplexity")
	err = security.CheckPasswordComplexity(newPassword)
	if err != nil {
		logger.Error(err)
		return
	}

	// get identity
	identity, err := GetIdentityByUID(service, resetPassword.IdentityUID)
	if err != nil {
		logger.Error(err)
		return
	}
	// change password
	err = SetPasswordOfIdentity(service, identity, newPassword)
	if err != nil {
		logger.Error(err)
		return
	}
	// update identity reset token
	now := time.Now().UTC()
	resetPassword.ConfirmationTime = &now
	resetPassword.ConfirmationUserAgent = userAgent
	resetPassword.ConfirmationIP = ip
	resetPassword.ConfirmationOrigin = origin

	// update in database
	err = service.GetSQLClient().Save(&resetPassword).Error
	if err != nil {
		logger.Error(err)
		// ignore error
		err = nil
	}
	return
}
