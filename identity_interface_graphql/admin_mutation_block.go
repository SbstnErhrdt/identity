package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// AdminBlockIdentity is a GraphQL mutation to block a user
func AdminBlockIdentity(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "adminBlockIdentity",
		Description: "blocks an identity so that the user can not login anymore",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"UID": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the uid of the user",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// params
			identityUID, err := ParseUIDFromArgs(&p, "UID")
			if err != nil {
				return nil, err
			}
			// get identity
			identity, err := identity_controllers.GetIdentityByUID(service, identityUID)
			if err != nil {
				return nil, err
			}
			// invite new user
			err = identity_controllers.Block(service, identity)
			return err == nil, err
		},
	}
	return &field
}

// AdminUnBlockIdentity is a GraphQL mutation to unblock a user
func AdminUnBlockIdentity(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "adminUnblockIdentity",
		Description: "unblocks an identity so that the user can login again",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"UID": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the uid of the user",
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
			identity, err := identity_controllers.GetIdentityByUID(service, identityUID)
			if err != nil {
				return nil, err
			}
			// invite new user
			err = identity_controllers.UnBlock(service, identity)
			return err == nil, err
		},
	}
	return &field
}
