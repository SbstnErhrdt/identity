package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"gorm.io/gorm"
)

// CheckIfEmailIsFree checks if the username is still free
func CheckIfEmailIsFree(service IdentityService, email string) (isFree bool, err error) {
	email = SanitizeEmail(email)
	err = service.GetSQLClient().
		Limit(1).
		Where("lower(email) = lower(?)", email).
		Where("deleted_at is NULL"). // only users that are not deleted
		First(&identity_models.Identity{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// username is free
		return true, nil
	} else if err != nil {
		service.GetLogger().WithField("email", email).WithError(err).Warn("could not check if email is free")
		return false, ErrEmailIsAlreadyRegistered
	}
	return false, nil
}
