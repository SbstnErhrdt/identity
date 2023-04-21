package identity_controllers

import "github.com/SbstnErhrdt/identity/identity_models"

// Block blocks a user
func Block(service IdentityService, user *identity_models.Identity) (err error) {
	user.Blocked = true
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		service.GetLogger().WithError(err).Error("could not block identity")
	}
	return
}
