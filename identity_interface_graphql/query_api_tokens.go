package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// ApiTokensField is the graphql field for the api tokens
func ApiTokensField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "ApiTokens",
		Description: "Retrieve the api tokens of the current user",
		Type:        graphql.NewList(ApiTokenGraphQlModel),
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			identityUID, err := GetIdentityUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			res, err := identity_controllers.GetApiTokensByIdentity(service, identityUID)
			return res, err
		},
	}
	return &field
}
