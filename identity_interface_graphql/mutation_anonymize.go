package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// AnonymizeField is the graphql field to anonymize an identity
func AnonymizeField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "anonymizeIdentity",
		Description: "anonymize the current identity",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the current password",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			// agent
			// extract uid
			uid, err := GetIdentityUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			password := p.Args["password"].(string)
			if len(password) == 0 {
				return nil, errors.New("password is required")
			}
			// delete account
			err = identity_controllers.AnonymizeIdentity(service, uid, password)
			return err == nil, err
		},
	}
	return &field
}
