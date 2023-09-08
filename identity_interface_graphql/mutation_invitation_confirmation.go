package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// InvitationConfirmationField is the graphql field to confirm an invitation
func InvitationConfirmationField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "invitationConfirmation",
		Description: "confirms an invitation",
		Type:        graphql.String,
		Args: graphql.FieldConfigArgument{
			"token": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the token of the invitation",
			},
			"password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the password",
			},
			"passwordConfirmation": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the password confirmation",
			},
			"acceptTermsAndConditions": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Boolean),
				Description: "the confirmation of the acceptance of the terms and conditions",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
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
			// params
			token := p.Args["token"].(string)
			if len(token) == 0 {
				return nil, errors.New("token is required")
			}
			password := p.Args["password"].(string)
			if len(password) == 0 {
				return nil, errors.New("password is required")
			}
			passwordConfirmation := p.Args["passwordConfirmation"].(string)
			if len(passwordConfirmation) == 0 {
				return nil, errors.New("password confirmation is required")
			}
			// check if passwords match
			if password != passwordConfirmation {
				err = errors.New("passwords do not match")
				return
			}
			termsAndConditions := p.Args["acceptTermsAndConditions"].(bool)
			if !termsAndConditions {
				err = errors.New("please accept the terms and conditions")
				return
			}
			// confirm invitation
			token, err = identity_controllers.InvitationConfirmation(service, token, password, passwordConfirmation, userAgent, ip, origin, termsAndConditions)
			return token, err
		},
	}
	return &field
}
