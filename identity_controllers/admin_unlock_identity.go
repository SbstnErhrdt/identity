package identity_controllers

import "github.com/SbstnErhrdt/identity/identity_models"

// Unlock unlocks an identity
func Unlock(service IdentityService, user *identity_models.Identity) (err error) {
	user.Active = true
	user.Cleared = true
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not unlock identity")
	}
	return
}
