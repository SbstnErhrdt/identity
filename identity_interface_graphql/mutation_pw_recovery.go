package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// InitResetPasswordField initializes the reset password flow
func InitResetPasswordField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "initPasswordReset",
		Description: "Initialize the reset of the password of the identity",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"identity": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The identity",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			// agent
			userAgent, err := GetUserAgentFromContext(&p)
			if err != nil {
				return
			}
			// ip
			ip, err := GetIpFromContext(&p)
			if err != nil {
				return
			}
			// origin
			origin, err := GetOriginFromContext(&p)
			if err != nil {
				return
			}
			// parameters
			identity := p.Args["identity"].(string)
			// get token of registration
			err = identity_controllers.InitResetPassword(service, identity, userAgent, ip, origin)
			return err == nil, err
		},
	}
	return &field
}

// ConfirmResetPasswordField confirms the reset password flow
func ConfirmResetPasswordField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "confirmResetPassword",
		Description: "Uses the token and sets an new password",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"token": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the reset password token",
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
			// from context
			// agent
			userAgent, err := GetUserAgentFromContext(&p)
			if err != nil {
				return
			}
			// ip
			ip, err := GetIpFromContext(&p)
			if err != nil {
				return
			}
			// origin
			origin, err := GetOriginFromContext(&p)
			if err != nil {
				return
			}
			// parameters
			token := p.Args["token"].(string)
			newPassword := p.Args["newPassword"].(string)
			newPasswordConfirmation := p.Args["newPasswordConfirmation"].(string)
			// get token of registration
			err = identity_controllers.ResetPassword(service, token, newPassword, newPasswordConfirmation, userAgent, ip, origin)
			return err == nil, err
		},
	}
	return &field
}
