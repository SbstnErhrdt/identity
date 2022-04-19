package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/controllers"
	"github.com/graphql-go/graphql"
)

func CurrentIdentityField(service controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "Identity",
		Description: "Retrieve the identity of the current user",
		Type:        IdentityGraphQlModel,
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			userUID, err := GetUserUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			res, err := controllers.GetIdentityByUID(service, userUID)
			return res, err
		},
	}
	return &field
}
