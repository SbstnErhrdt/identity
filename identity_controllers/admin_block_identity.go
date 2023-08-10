package identity_controllers

import "github.com/SbstnErhrdt/identity/identity_models"

// Block blocks a user
func Block(service IdentityService, user *identity_models.Identity) (err error) {
	user.Blocked = true
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not block identity")
	}
	return
}

// UnBlock unblocks a user
func UnBlock(service IdentityService, user *identity_models.Identity) (err error) {
	user.Blocked = false
	err = service.GetSQLClient().Save(user).Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not unblock identity")
	}
	return
}
