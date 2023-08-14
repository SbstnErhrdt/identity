package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// IdentityApiTokensField is the graphql field for the api tokens of an identity
func IdentityApiTokensField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "ApiTokens",
		Description: "Retrieve the api tokens of the current user",
		Type:        graphql.NewList(ApiTokenGraphQlModel),
		Args: graphql.FieldConfigArgument{
			"UID": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "The uid of the user",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// params
			userUID, err := ParseUIDFromArgs(&p, "UID")
			if err != nil {
				return nil, err
			}
			res, err := identity_controllers.GetApiTokensByIdentity(service, userUID)
			return res, err
		},
	}
	return &field
}
