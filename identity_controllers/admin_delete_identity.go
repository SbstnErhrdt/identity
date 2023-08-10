package identity_controllers

import (
	"github.com/google/uuid"
	"log/slog"
)

// AdminDeleteIdentity deletes an identity
// - anonymizes the firstname and lastname
// - anonymizes the email and phone
// - deletes the account
func AdminDeleteIdentity(service IdentityService, uid uuid.UUID) (err error) {
	// get the user
	identity, err := GetIdentityByUID(service, uid)
	if err != nil {
		slog.With("err", err).Error("could not get identity")
		return
	}
	return deleteIdentity(service, identity)
}
