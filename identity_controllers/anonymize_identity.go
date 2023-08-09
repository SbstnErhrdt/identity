package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"time"
)

// anonymizeFirstNameLastNameAndSave anonymizes an identity
// - hashes the firstname and lastname
// - sets the updated at timestamp
// - saves the identity
func anonymizeFirstNameLastNameAndSave(service IdentityService, identity *identity_models.Identity) (err error) {
	// hash firstname and lastname
	identity.FirstName = AnonymizeString(identity.FirstName)
	identity.LastName = AnonymizeString(identity.LastName)
	// save the account
	// Set other metadata
	identity.UpdatedAt = time.Now().UTC()
	// create new entry
	err = service.GetSQLClient().Save(identity).Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not anonymizeFirstNameLastNameAndSave identity")
	}
	return
}

// AnonymizeIdentity anonymizes the first name, the last NameAndSave an account
// - gets the identity by uid
// - checks if the password is correct
// - anonymizes the identity
// - saves the identity
func AnonymizeIdentity(service IdentityService, uid uuid.UUID, password string) (err error) {
	logger := service.GetLogger().With(
		"method", "AnonymizeIdentity",
		"uid", uid,
	)
	// get the user
	identity, err := GetIdentityByUID(service, uid)
	if err != nil {
		logger.With("err", err).Error("could not get identity")
		return
	}
	// check if the password is correct
	if !VerifyPassword(service, identity, password) {
		err = ErrInvalidPassword
		logger.With("err", err).Error("could not verify password")
		return
	}
	return anonymizeFirstNameLastNameAndSave(service, identity)
}
