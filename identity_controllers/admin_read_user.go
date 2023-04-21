package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// ReadIdentity reads a specific user
func ReadIdentity(service IdentityService, adminUID, identityUID uuid.UUID) (result *identity_models.Identity, err error) {
	// check if admin
	err = IsAdmin(service, adminUID)
	if err != nil {
		return
	}
	// build query
	err = service.GetSQLClient().
		Where("deleted_at IS NULL").
		Where("uid = ?", identityUID).
		// execute the query
		First(&result).
		Error
	if err != nil {
		log.WithError(err).Error("could not read users")
	}
	return
}
