package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// ChangePasswordField is the graphql field for change password
func ChangePasswordField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "changePassword",
		Description: "Change the password of the identity of the current identity",
		Type:        graphql.Boolean,
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
			oldPassword := p.Args["oldPassword"].(string)
			if len(oldPassword) == 0 {
				return nil, errors.New("password is required")
			}
			newPassword := p.Args["newPassword"].(string)
			if len(newPassword) == 0 {
				return nil, errors.New("password is required")
			}
			newPasswordConfirmation := p.Args["newPasswordConfirmation"].(string)
			if len(newPasswordConfirmation) == 0 {
				return nil, errors.New("password is required")
			}
			// extract uid
			uid, err := GetUserUIDFromContext(&p)
			// update
			err = identity_controllers.ChangePassword(service, uid, oldPassword, newPassword, newPasswordConfirmation)
			return err == nil, err
		},
	}
	return &field
}
