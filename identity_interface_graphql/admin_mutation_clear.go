package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// AdminClearIdentity is a GraphQL mutation to clear a user
func AdminClearIdentity(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "adminClearIdentity",
		Description: "clears an identity so that does not need to confirm his email anymore",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"UID": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The uid of the user",
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
			err = identity_controllers.Clear(service, identity)
			return err == nil, err
		},
	}
	return &field
}
