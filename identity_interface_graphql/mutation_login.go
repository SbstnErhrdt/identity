package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/controllers"
	"github.com/graphql-go/graphql"
	"strings"
)

func LoginField(service controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "Login",
		Description: "Submit identity and password to retrieve a token",
		Type:        graphql.String,
		Args: graphql.FieldConfigArgument{
			"identity": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The identity",
			},
			"password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The password",
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
			identity := p.Args["identity"].(string)
			identity = strings.ToLower(identity)
			password := p.Args["password"].(string)
			// login
			token, err := controllers.Login(service, identity, password, userAgent, ip)
			return token, err
		},
	}
	return &field
}
