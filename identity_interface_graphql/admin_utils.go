package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
	"log/slog"
)

// CheckAdmin checks if the user is admin
func CheckAdmin(service identity_controllers.IdentityService, p *graphql.ResolveParams) (err error) {
	// from context
	identityUID, err := GetIdentityUIDFromContext(p)
	if err != nil {
		slog.With("err", err).Error("Could not get user from context")
		return err
	}
	// check if user is admin
	err = identity_controllers.IsAdmin(service, identityUID)
	if err != nil {
		slog.With("err", err).With("securityIncident", 2).Error("Identity is not admin")
		return err
	}
	return
}
