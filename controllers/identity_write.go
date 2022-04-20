package controllers

import (
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/models"
	"github.com/SbstnErhrdt/identity/security"
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

func DeleteAccount(service IdentityService, username, password string) (user *models.Identity, err error) {
	return
}

func deleteAccount(service IdentityService, identity *models.Identity) (err error) {
	// delete identity
	err = service.GetSQLClient().Delete(&identity).Error
	return
}

func AnonymizeAccount(service IdentityService, user *models.Identity) (err error) {
	// todo: anonymize user
	return
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

// VerifyPassword verifies the user's password given the user object and the password
func VerifyPassword(service IdentityService, user *models.Identity, password string) bool {
	return user.CheckPassword(service.GetPepper(), password)
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
