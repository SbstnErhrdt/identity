package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// IdentityField is the graphql field for the identity
func IdentityField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "Identity",
		Description: "Retrieve the identity of a specific user",
		Type:        AdminIdentityGraphQlModel,
		Args: graphql.FieldConfigArgument{
			"UID": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
				Description:  "The uid of the identity",
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
			res, err := identity_controllers.ReadIdentity(service, identityUID)
			if err != nil {
				return nil, err
			}
			return res, nil
		},
	}
	return &field
}
