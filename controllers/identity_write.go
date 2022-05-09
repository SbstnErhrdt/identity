package controllers

import (
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/models"
	"github.com/SbstnErhrdt/identity/security"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// ErrNoEmail is returned when no email is provided
var ErrNoEmail = errors.New("no email address")

// ErrNoPassword is returned when no password is provided
var ErrNoPassword = errors.New("no password")

// ErrEmailAlreadyExists is returned when the email address is already in use
var ErrEmailAlreadyExists = errors.New("email already exists")

func Logout(service IdentityService, token string) (err error) {
	// add the token uuid to the expired database
	return
}

func LogoutAllDevices(service IdentityService, user *models.Identity) (err error) {
	return
}

func ConfirmResetPassword(service IdentityService, token, password string) (err error) {
	return
}

func InitChangeEmail(service IdentityService, username, password string) (user *models.Identity, err error) {
	return
}

func ConfirmOldEmail(service IdentityService, token string) {
}

func ConfirmNewEmail(service IdentityService, token string) {
}

// DeleteIdentity deletes an identity
func DeleteIdentity(service IdentityService, uid uuid.UUID, password string) (err error) {
	logger := log.WithFields(log.Fields{
		"method": "DeleteIdentity",
		"uid":    uid,
	})
	// get the user
	identity, err := GetIdentityByUID(service, uid)
	if err != nil {
		logger.Error(err)
		return
	}
	// check if the password is correct
	if !VerifyPassword(service, identity, password) {
		err = ErrInvalidPassword
		logger.Error(err)
		return
	}
	// anonymize the pii
	err = anonymize(service, identity)
	if err != nil {
		logger.Error(err)
		return
	}
	// anonymize the email and phone
	err = deleteAnonymize(service, identity)
	if err != nil {
		logger.Error(err)
		return
	}
	//delete the account
	err = deleteAccount(service, identity)
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

func deleteAccount(service IdentityService, identity *models.Identity) (err error) {
	// delete identity
	err = service.GetSQLClient().Delete(&identity).Error
	return
}

func anonymize(service IdentityService, identity *models.Identity) (err error) {
	// hash emails and phone numbers
	identity.FirstName = Hash(identity.FirstName)
	identity.LastName = Hash(identity.LastName)
	identity.Salutation = Hash(identity.Salutation)
	// save the account
	// Set other metadata
	identity.UpdatedAt = time.Now().UTC()
	// create new entry
	err = service.GetSQLClient().Save(identity).Error
	return
}

func deleteAnonymize(service IdentityService, identity *models.Identity) (err error) {
	// hash emails and phone numbers
	identity.Email = Hash(identity.Email)
	identity.BackupEmail = Hash(identity.BackupEmail)
	identity.Phone = Hash(identity.Phone)
	identity.BackupPhone = Hash(identity.BackupPhone)
	// save the account
	// Set other metadata
	identity.UpdatedAt = time.Now().UTC()
	// create new entry
	err = service.GetSQLClient().Save(identity).Error
	return
}

// AnonymizeIdentity anonymize an account
func AnonymizeIdentity(service IdentityService, uid uuid.UUID, password string) (err error) {
	logger := log.WithFields(log.Fields{
		"method": "AnonymizeIdentity",
		"uid":    uid,
	})
	// get the user
	identity, err := GetIdentityByUID(service, uid)
	if err != nil {
		logger.Error(err)
		return
	}
	// check if the password is correct
	if !VerifyPassword(service, identity, password) {
		err = ErrInvalidPassword
		logger.Error(err)
		return
	}
	return anonymize(service, identity)
}

// Clear clears the user
func Clear(service IdentityService, user *models.Identity) {
	user.Cleared = true
	service.GetSQLClient().Save(user)
	return
}

// Block blocks a user
func Block(service IdentityService, user *models.Identity) {
	user.Blocked = true
	service.GetSQLClient().Save(user)
	return
}

// GenerateJWT generates a Json Web Token from the user object
func GenerateJWT(service IdentityService, user *models.Identity) (result string, err error) {
	// Init the token structure
	payload := map[string]interface{}{}
	payload["userUID"] = user.UID
	// Generate the token
	audience := env.FallbackEnvVariable("SECURITY_JWT_AUDIENCE", "API")
	result, _, err = security.GenerateJWTToken(user.UID, audience, payload)
	if err != nil {
		log.Error(err)
	}
	return
}

// UpdateIdentity updates the identity
func UpdateIdentity(service IdentityService, newIdentity *models.Identity) (err error) {
	// check if necessary fields are set
	if newIdentity.UID.String() == "" {
		return errors.New("UpdateIdentity: no uid provided")
	}
	// Get latest version from db
	dbObj, dbErr := GetIdentityByUID(service, newIdentity.UID)
	if dbErr != nil {
		log.Error(dbErr)
		err = errors.New("can not read identity from database")
		return
	}
	// Update data of database object
	dbObj.FirstName = newIdentity.FirstName
	dbObj.LastName = newIdentity.LastName
	dbObj.Phone = newIdentity.Phone
	dbObj.BackupPhone = newIdentity.BackupPhone
	dbObj.Email = strings.ToLower(newIdentity.Email)
	dbObj.BackupEmail = strings.ToLower(newIdentity.BackupEmail)
	dbObj.Salutation = newIdentity.Salutation
	// Set other metadata
	dbObj.UpdatedAt = time.Now().UTC()
	// create new entry
	err = service.GetSQLClient().Save(dbObj).Error
	if err != nil {
		log.Error(err)
		return
	}
	newIdentity = dbObj
	return
}
