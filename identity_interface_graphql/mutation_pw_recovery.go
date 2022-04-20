package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/controllers"
	"github.com/graphql-go/graphql"
)

func InitPasswordResetField(service controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "InitPasswordReset",
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
			err = controllers.InitResetPassword(service, identity, userAgent, ip, origin)
			return err == nil, err
		},
	}
	return &field
}

func ConfirmPasswordResetField(service controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "InitPasswordReset",
		Description: "Initialize the reset of the password of the identity",
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
			// get token of registration
			err = controllers.ResetPassword(service, token, newPassword, userAgent, ip, origin)
			return err == nil, err
		},
	}
	return &field
}
