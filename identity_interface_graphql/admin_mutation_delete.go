package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// AdminDeleteIdentity is a GraphQL mutation to delete a user
func AdminDeleteIdentity(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "adminDeleteIdentity",
		Description: "deletes an identity",
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
			err = identity_controllers.AdminDeleteIdentity(service, identityUID)
			return err == nil, err
		},
	}
	return &field
}
