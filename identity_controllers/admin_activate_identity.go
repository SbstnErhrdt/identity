package identity_controllers

import "github.com/SbstnErhrdt/identity/identity_models"

// Activate activates the user
func Activate(service IdentityService, user *identity_models.Identity) (err error) {
	user.Active = true
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not activate identity")
	}
	return
}
