package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
)

// ErrExternalNoTokensFound is returned when no tokens are found
var ErrExternalNoTokensFound = errors.New("no tokens found")

// GetApiTokensByIdentity retrieves the api tokens of an identity from the database
func GetApiTokensByIdentity(service IdentityService, identityUID uuid.UUID) (results []*identity_models.IdentityApiToken, err error) {
	// Load identity from database
	results = []*identity_models.IdentityApiToken{}
	err = service.GetSQLClient().
		Where("identity_uid = ?", identityUID.String()).
		Where("deleted_at is NULL").
		Order("created_at desc").
		Find(&results).Error
	if err != nil {
		service.GetLogger().WithError(err).WithField("identity_uid", identityUID).Error("could not find tokens with uid")
		return nil, ErrExternalNoTokensFound
	}
	return results, nil
}
