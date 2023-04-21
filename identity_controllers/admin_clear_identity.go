package identity_controllers

import "github.com/SbstnErhrdt/identity/identity_models"

// Clear clears the user
func Clear(service IdentityService, user *identity_models.Identity) (err error) {
	user.Cleared = true
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		service.GetLogger().WithError(err).Error("could not clear identity")
	}
	return
}
