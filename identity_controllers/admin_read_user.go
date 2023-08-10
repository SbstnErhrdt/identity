package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
)

// ReadIdentity reads a specific user
func ReadIdentity(service IdentityService, identityUID uuid.UUID) (result *identity_models.Identity, err error) {
	// build query
	err = service.GetSQLClient().
		Where("deleted_at IS NULL").
		Where("uid = ?", identityUID).
		// execute the query
		First(&result).
		Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not read users")
	}
	return
}
