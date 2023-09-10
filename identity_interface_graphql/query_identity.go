package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// CurrentIdentityField is the graphql field that returns the current identity
func CurrentIdentityField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "CurrentIdentity",
		Description: "retrieve the identity information of the current identity",
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
