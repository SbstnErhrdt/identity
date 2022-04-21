package controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/models"
	"github.com/SbstnErhrdt/identity/security"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// VerifyPassword verifies the user's password given the user object and the password
func VerifyPassword(service IdentityService, user *models.Identity, password string) bool {
	return user.CheckPassword(service.GetPepper(), password)
}

// ErrInvalidPassword is returned when the password is invalid
var ErrInvalidPassword = errors.New("old password is incorrect")

// ErrConfirmPassword is returned when the new password and confirmation do not match
var ErrConfirmPassword = errors.New("new password and confirmation do not match")

// ErrOldPasswordIsSame is returned when the new password is the same as the old password
var ErrOldPasswordIsSame = errors.New("the old password is the same as the new password")

// ChangePassword changes the user's password given the user object and the new password
func ChangePassword(service IdentityService, identityUID uuid.UUID, oldPassword, newPassword, newPasswordConfirmation string) (err error) {
	// init logger
	logger := log.WithFields(log.Fields{
		"process":     "ChangePassword",
		"identityUID": identityUID,
	})

	// get the user
	identity, err := GetIdentityByUID(service, identityUID)
	if err != nil {
		logger.Error(err)
		return err
	}

	// check if the old password is correct
	if !VerifyPassword(service, identity, oldPassword) {
		err = ErrInvalidPassword
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

	// check if the new password is the same as the confirmed password
	if newPassword != newPasswordConfirmation {
		err = ErrConfirmPassword
		logger.Error(err)
		return
	}

	// check if the old password is the same as the new password
	if oldPassword == newPassword {
		err = ErrOldPasswordIsSame
		logger.Error(err)
		return
	}

	// set the new password
	err = identity.SetNewPassword(service.GetPepper(), newPassword)
	if err != nil {
		logger.Error(err)
		return err
	}

	// save the user
	err = service.GetSQLClient().Save(&identity).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return
}

// ErrCanNotChangePassword is returned when the user can not change the password
var ErrCanNotChangePassword = errors.New("can not change password. please try again")

// SetPasswordOfIdentity set the password of a user by its email
func SetPasswordOfIdentity(service IdentityService, user *models.Identity, newPassword string) (err error) {
	// init logger
	logger := log.WithFields(log.Fields{
		"process": "SetPasswordOfIdentity",
	})
	// check password complexity
	logger.Debug("CheckPasswordComplexity")
	err = security.CheckPasswordComplexity(newPassword)
	if err != nil {
		logger.Error(err)
		return
	}
	// set password
	err = user.SetNewPassword(service.GetPepper(), newPassword)
	if err != nil {
		log.Error(err)
		// overwrite error with generic error message
		err = ErrCanNotChangePassword
		return
	}
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		log.Error(err)
		// overwrite err
		err = ErrCanNotChangePassword
		return
	}
	return
}
