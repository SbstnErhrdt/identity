package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"time"
)

// UpdateIdentity updates the identity
func UpdateIdentity(service IdentityService, newIdentity *identity_models.Identity) (err error) {
	logger := service.GetLogger().With(
		"method", "UpdateIdentity",
		"uid", newIdentity.UID,
	)
	// check if necessary fields are set
	if newIdentity.UID.String() == "" || newIdentity.UID == uuid.Nil {
		return ErrNoIdentityIdentification
	}
	// Get latest version from db
	dbObj, dbErr := GetIdentityByUID(service, newIdentity.UID)
	if dbErr != nil {
		logger.With("err", dbErr).Error("could not get identity")
		err = ErrNoIdentity
		return
	}
	// Update data of database object
	dbObj.FirstName = newIdentity.FirstName
	dbObj.LastName = newIdentity.LastName
	dbObj.Phone = newIdentity.Phone
	dbObj.BackupPhone = newIdentity.BackupPhone
	dbObj.Email = SanitizeEmail(newIdentity.Email)
	dbObj.BackupEmail = SanitizeEmail(newIdentity.BackupEmail)
	dbObj.Salutation = newIdentity.Salutation
	// Set other metadata
	dbObj.UpdatedAt = time.Now().UTC()
	// create new entry
	err = service.GetSQLClient().Save(dbObj).Error
	if err != nil {
		logger.With("err", err).Error("could not update identity")
		return
	}
	newIdentity = dbObj
	return
}
