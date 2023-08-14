package identity_controllers

import "github.com/SbstnErhrdt/identity/identity_models"

// Lock locks an identity
func Lock(service IdentityService, user *identity_models.Identity) (err error) {
	user.Active = true
	user.Cleared = false
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not lock identity")
	}
	return
}
