package identity_interface_graphql

import (
	"github.com/graphql-go/graphql"
)

func RegistrationConfirmationField(service identity_controllers.IdentityService) *graphql.Field {
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
			userAgent, err := GetUserAgentFromContext(&p)
			if err != nil {
				return
			}
			// ip
			ip, err := GetIpFromContext(&p)
			if err != nil {
				return
			}
			// params
			token := p.Args["token"].(string)
			// login
			err = identity_controllers.RegistrationConfirmation(service, token, userAgent, ip)
			return err == nil, err
		},
	}
	return &field
}
