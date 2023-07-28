package identity_controllers

import (
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/SbstnErhrdt/identity/security"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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

func LogoutAllDevices(service IdentityService, user *identity_models.Identity) (err error) {
	return
}

func ConfirmResetPassword(service IdentityService, token, password string) (err error) {
	return
}

func InitChangeEmail(service IdentityService, username, password string) (user *identity_models.Identity, err error) {
	return
}

func ConfirmOldEmail(service IdentityService, token string) {
}

func ConfirmNewEmail(service IdentityService, token string) {
}

// GenerateJWT generates a Json Web Token from the user object
func GenerateJWT(service IdentityService, user *identity_models.Identity) (result string, err error) {
	// Init the token structure
	payload := map[string]interface{}{}
	payload["userUID"] = user.UID
	// Generate the token
	audience := env.FallbackEnvVariable("SECURITY_JWT_AUDIENCE", "APP")
	result, _, err = security.GenerateJWTToken(user.UID, audience, payload)
	if err != nil {
		log.Error(err)
	}
	return
}

// ErrNoIdentity is returned if no identity is found
var ErrNoIdentity = errors.New("no identity found")

// ErrNoIdentityIdentification is returned if no identification is provided
var ErrNoIdentityIdentification = errors.New("no identity identification provided. Please contact the administrator")
