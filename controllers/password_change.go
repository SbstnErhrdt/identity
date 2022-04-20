package controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/models"
	log "github.com/sirupsen/logrus"
)

// ErrCanNotChangePassword is returned when the user can not change the password
var ErrCanNotChangePassword = errors.New("can not change password. please try again")

// SetPasswordOfIdentity set the password of a user by its email
func SetPasswordOfIdentity(service IdentityService, user *models.Identity, newPassword string) (err error) {
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
