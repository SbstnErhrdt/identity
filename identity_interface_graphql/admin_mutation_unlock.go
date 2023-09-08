package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// AdminUnlockIdentity is a GraphQL mutation to unlock an identity
func AdminUnlockIdentity(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "adminUnlockIdentity",
		Type:        graphql.Boolean,
		Description: "unlocks an identity so that the user does not need to confirm his email anymore, is allowed to login",
		Args: graphql.FieldConfigArgument{
			"UID": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the uid of the user",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			err = CheckAdmin(service, &p)
			if err != nil {
				return nil, err
			}
			// params
			identityUID, err := ParseUIDFromArgs(&p, "UID")
			if err != nil {
				return nil, err
			}
			// get identity
			identity, err := identity_controllers.GetIdentityByUID(service, identityUID)
			if err != nil {
				return nil, err
			}
			// invite new user
			err = identity_controllers.Unlock(service, identity)
			return err == nil, err
		},
	}
	return &field
}
