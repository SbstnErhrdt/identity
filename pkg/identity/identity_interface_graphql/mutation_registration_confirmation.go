package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/pkg/identity/controllers"
	"github.com/graphql-go/graphql"
)

func RegistrationConfirmationField(service controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "RegistrationConfirmation",
		Description: "Confirm the token that was received via email",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"token": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The confirmation token",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			// agent
			userAgent, ok := p.Context.Value("USER_AGENT").(string)
			if !ok {
				err = errors.New("can not extract agent from context")
				return
			}
			// ip
			ip, ok := p.Context.Value("USER_IP").(string)
			if !ok {
				err = errors.New("can not extract ip from context")
				return
			}
			// params
			token := p.Args["token"].(string)
			// login
			err = controllers.RegistrationConfirmation(service, token, userAgent, ip)
			return err == nil, err
		},
	}
	return &field
}
