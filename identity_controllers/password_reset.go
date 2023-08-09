package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/SbstnErhrdt/identity/security"
	"net/mail"
	"time"
)

// InitResetPassword inits the password reset process
func InitResetPassword(service IdentityService, emailAddress, userAgent, ip, origin string) (err error) {
	// init logger
	logger := service.GetLogger().With(
		"email", emailAddress,
		"process", "InitResetPassword",
	)
	// get identity
	res, err := GetIdentityByEmail(service, emailAddress)
	if err != nil {
		logger.With("err", err).Error("could not get identity")
		// OWASP: return no error if the user does not exist
		return nil
	}
	// create database entry
	token, err := security.GenerateRandomString(32)
	if err != nil {
		logger.With("err", err).Error("could not generate token")
		return err
	}
	resetPassword := identity_models.IdentityResetPassword{
		IdentityUID: res.UID,
		Token:       token,
		Expire:      time.Now().UTC().Add(time.Hour * 24),
		UserAgent:   userAgent,
		IP:          ip,
	}
	// save in database
	err = service.GetSQLClient().Create(&resetPassword).Error
	if err != nil {
		logger.With("err", err).Error("cannot save reset password in database")
		return err
	}
	// send email
	// Build email template
	logger.Debug("Build email template")
	// resolve the email template
	emailTemplate := service.ResolvePasswordResetEmailTemplate(origin, emailAddress, token)
	// generate the content of the email
	content, err := emailTemplate.Content()
	if err != nil {
		logger.With("err", err).Error("cannot generate email content")
		return
	}
	// Send email
	logger.Debug("Send reset password email")
	err = service.SendEmail(
		service.GetSenderEmailAddress(),
		mail.Address{
			Name:    emailAddress,
			Address: emailAddress,
		},
		"Password reset",
		content)
	return
}

// ErrTokenExpired is returned when the token is expired
var ErrTokenExpired = errors.New("security token expired. Please initialize a new password reset")

// ErrTokenUsed is returned when the token was used
var ErrTokenUsed = errors.New("security token already used. Please request a new password reset")

// ErrTokenNotFound is returned when the token is not in the database
var ErrTokenNotFound = errors.New("security token not found. Please request a new password reset")

// ResetPassword resets the password
func ResetPassword(service IdentityService, token, newPassword, newPasswordConfirmation, userAgent, ip, origin string) (err error) {
	// init logger
	logger := service.GetLogger().With(
		"token", token,
		"process", "InitResetPassword",
	)

	// check password complexity
	logger.Debug("CheckPasswordComplexity")
	err = security.CheckPasswordComplexity(newPassword)
	if err != nil {
		logger.With("err", err).Error("cannot check password complexity")
		return
	}

	// check if the new password is the same as the confirmed password
	if newPassword != newPasswordConfirmation {
		err = ErrConfirmPassword
		logger.With("err", err).Error("can not confirm password")
		return
	}

	// get the token from the database
	resetPassword := identity_models.IdentityResetPassword{}
	err = service.GetSQLClient().Where("token = ?", token).First(&resetPassword).Error
	if err != nil {
		logger.With("err", err).Error("cam not get reset password from database")
		err = ErrTokenNotFound
		return err
	}

	// check if expiration date is in the past
	if resetPassword.Expire.Before(time.Now().UTC()) {
		err = ErrTokenExpired
		logger.With("err", err).Error("can not reset password")
		return
	}
	// check if the token was used
	if resetPassword.ConfirmationTime != nil {
		err = ErrTokenUsed
		logger.With("err", err).Error("can not reset password")
		return
	}

	// check password complexity
	logger.Debug("CheckPasswordComplexity")
	err = security.CheckPasswordComplexity(newPassword)
	if err != nil {
		logger.With("err", err).Error("cannot check password complexity")
		return
	}

	// get identity
	identity, err := GetIdentityByUID(service, resetPassword.IdentityUID)
	if err != nil {
		logger.With("err", err).Error("cannot get identity from database")
		return
	}
	// change password
	err = SetPasswordOfIdentity(service, identity, newPassword)
	if err != nil {
		logger.With("err", err).Error("cannot set password of identity")
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
		logger.With("err", err).Error("cannot update reset password in database")
		// ignore error
		err = nil
	}
	return
}
