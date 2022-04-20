package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/controllers"
	"github.com/graphql-go/graphql"
)

// ChangePasswordField is the graphql field for change password
func ChangePasswordField(service controllers.IdentityService) *graphql.Field {
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
			newPassword := p.Args["newPassword"].(string)
			newPasswordConfirmation := p.Args["newPasswordConfirmation"].(string)
			// extract uid
			uid, err := GetUserUIDFromContext(&p)
			// update
			err = controllers.ChangePassword(service, uid, oldPassword, newPassword, newPasswordConfirmation)
			return err == nil, err
		},
	}
	return &field
}
