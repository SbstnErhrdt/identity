package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// AdminLockIdentity is a GraphQL mutation to lock an identity
func AdminLockIdentity(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "adminLockIdentity",
		Description: "locks an identity so that the user needs to wait for an admin to unlock it",
		Type:        graphql.Boolean,
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
			err = identity_controllers.Lock(service, identity)
			return err == nil, err
		},
	}
	return &field
}
