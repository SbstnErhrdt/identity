package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
	"time"
)

// CreateApiTokenField is the graphql field to create a new api token
func CreateApiTokenField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "createApiToken",
		Description: "create a new api token for the current user",
		Type:        ApiTokenGraphQlModel,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type:         graphql.NewNonNull(graphql.String),
				Description:  "the name of the api token",
				DefaultValue: "default api token",
			},
			"expirationDate": &graphql.ArgumentConfig{
				Type:         graphql.DateTime,
				Description:  "the expiration date of the api token",
				DefaultValue: nil,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// from context
			// agent
			// extract uid
			uid, err := GetIdentityUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			name := p.Args["name"].(string)
			if len(name) == 0 {
				return nil, errors.New("name is required")
			}
			// expiration date
			expirationDate, ok := p.Args["expirationDate"].(time.Time)
			if !ok {
				expirationDate = time.Now().Add(time.Hour * 24 * 7 * 30 * 12 * 4) // 4 years
			}
			// parse date

			// delete account
			token, err := identity_controllers.CreateApiToken(service, uid, name, expirationDate)
			if err != nil {
				return nil, err
			}
			return token, nil
		},
	}
	return &field
}
