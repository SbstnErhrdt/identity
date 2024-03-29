package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
	"strings"
)

// LoginField is the graphql field to log in a user
func LoginField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "login",
		Description: "Submit identity and password to retrieve a token",
		Type:        graphql.String,
		Args: graphql.FieldConfigArgument{
			"identity": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the identity",
			},
			"password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "the password",
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
			token, err := identity_controllers.Login(service, identity, password, userAgent, ip)
			return token, err
		},
	}
	return &field
}
