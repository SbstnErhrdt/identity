package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"time"
)

// ErrExternalDelete is returned when the email address is already in use
var ErrExternalDelete = errors.New("user could not be deleted. Please try again later")

// DeleteIdentity deletes an identity
// - gets the identity
// - checks if the password is correct
// - anonymizes the firstname and lastname
// - anonymizes the email and phone
// - deletes the account
func DeleteIdentity(service IdentityService, uid uuid.UUID, password string) (err error) {
	logger := service.GetLogger().With(
		"method", "DeleteIdentity",
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
	// anonymize firstname and last name
	err = anonymizeFirstNameLastNameAndSave(service, identity)
	if err != nil {
		logger.With("err", err).Error("could not anonymize firstname and last name")
		err = ErrExternalDelete
		return
	}
	// anonymize the email and phone
	err = anonymizeEmailPhone(service, identity)
	if err != nil {
		logger.With("err", err).Error("")
		err = ErrExternalDelete
		return
	}
	// delete the account
	err = softDeleteAccount(service, identity)
	if err != nil {
		logger.With("err", err).Error("could not delete account")
		err = ErrExternalDelete
		return
	}
	return
}

// softDeleteAccount deletes the account
// - deletes the identity (marks the account as deleted)
func softDeleteAccount(service IdentityService, identity *identity_models.Identity) (err error) {
	// delete identity
	err = service.GetSQLClient().Delete(&identity).Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not delete identity")
	}
	return
}

// anonymizeEmailPhone anonymizes the email and phone
// - hashes the email and phone
// - sets the updated at timestamp
// - saves the identity
func anonymizeEmailPhone(service IdentityService, identity *identity_models.Identity) (err error) {
	// hash emails and phone numbers
	identity.Email = AnonymizeString(identity.Email)
	identity.BackupEmail = AnonymizeString(identity.BackupEmail)
	identity.Phone = AnonymizeString(identity.Phone)
	identity.BackupPhone = AnonymizeString(identity.BackupPhone)
	// save the account
	// Set other metadata
	identity.UpdatedAt = time.Now().UTC()
	// create new entry
	err = service.GetSQLClient().Save(identity).Error
	return
}
