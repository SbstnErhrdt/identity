package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

func CurrentIdentityField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "CurrentIdentity",
		Description: "Retrieve the identity of the current user",
		Type:        IdentityGraphQlModel,
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			identityUID, err := GetIdentityUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			res, err := identity_controllers.GetIdentityByUID(service, identityUID)
			return res, err
		},
	}
	return &field
}
