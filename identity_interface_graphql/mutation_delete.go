package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// DeleteField is the graphql field for deleting an identity
func DeleteField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "deleteIdentity",
		Description: "Delete the current identity",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The current password",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			// agent
			// extract uid
			uid, err := GetUserUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			password := p.Args["password"].(string)
			if len(password) == 0 {
				return nil, errors.New("password is required")
			}
			// delete account
			err = identity_controllers.DeleteIdentity(service, uid, password)
			return err == nil, err
		},
	}
	return &field
}
