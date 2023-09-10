package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// CurrentIdentityIsAdminField is the graphql field that returns true if the current identity is an admin
func CurrentIdentityIsAdminField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "CurrentIdentityIsAdmin",
		Description: "returns true of the current identity is an admin",
		Type:        graphql.NewList(ApiTokenGraphQlModel),
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			identityUID, err := GetIdentityUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			err = identity_controllers.IsAdmin(service, identityUID)
			return err == nil, err
		},
	}
	return &field
}
