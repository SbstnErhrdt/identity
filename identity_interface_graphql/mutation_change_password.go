package identity_interface_graphql

import (
	"encoding/json"
	"github.com/SbstnErhrdt/identity/controllers"
	"github.com/SbstnErhrdt/identity/models"
	"github.com/graphql-go/graphql"
)

// ChangePasswordField is the graphql field for change password
func ChangePasswordField(service controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "changePassword",
		Description: "Change the password of the identity of the current identity",
		Type:        IdentityGraphQlModel,
		Args: graphql.FieldConfigArgument{
			"oldPassword": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the old password",
			},
			"newPassword": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the new password",
			},
			"newPasswordConfirmation": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the new password confirmation",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// params
			jsonString, err := json.Marshal(p.Args["data"])
			if err != nil {
				return nil, err
			}
			identity := models.Identity{}
			err = json.Unmarshal(jsonString, &identity)
			if err != nil {
				return nil, err
			}
			// extract uid
			uid, err := GetUserUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			// overwrite uid
			identity.UID = uid
			// update
			err = controllers.UpdateIdentity(service, &identity)
			return &identity, err
		},
	}
	return &field
}
