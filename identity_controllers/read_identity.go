package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
)

// GetIdentityByEmail retrieves the identity from the database
func GetIdentityByEmail(service IdentityService, email string) (result *identity_models.Identity, err error) {
	email = SanitizeEmail(email)
	// Load identity from database
	identity := identity_models.Identity{}
	err = service.GetSQLClient().
		Limit(1).
		Where("email = ?", email).
		Where("deleted_at is NULL"). // only users that are not deleted
		First(&identity).Error
	if err != nil {
		service.GetLogger().WithError(err).WithField("email", email).Error("could not find user with email")
		return nil, ErrNoUserFound
	}
	return &identity, nil
}

// GetIdentityByUID retrieves the identity from the database
func GetIdentityByUID(service IdentityService, uid uuid.UUID) (result *identity_models.Identity, err error) {
	// Load identity from database
	identity := identity_models.Identity{}
	err = service.GetSQLClient().
		Limit(1).
		Where("uid = ?", uid.String()).
		Where("deleted_at is NULL"). // only users that are not deleted
		First(&identity).Error
	if err != nil {
		service.GetLogger().WithError(err).WithField("uid", uid).Error("could not find user with uid")
		return nil, ErrNoUserFound
	}
	return &identity, nil
}
