package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// DeleteApiTokenField is the graphql field for deleting an api token
func DeleteApiTokenField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "deleteApiToken",
		Description: "Delete a api token for the current user",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"UID": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the uid of the api token",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// get identity uid from context
			identityUID, err := GetIdentityUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			// parse token uid
			uidString := p.Args["UID"].(string)
			tokenUID, err := uuid.Parse(uidString)
			if err != nil {
				return nil, errors.New("token uid is not valid")
			}
			// delete api token
			err = identity_controllers.DeleteApiToken(service, identityUID, tokenUID)
			return err == nil, err
		},
	}
	return &field
}
