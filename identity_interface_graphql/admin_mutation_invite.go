package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// AdminInviteField is the graphql field to invite a new user
func AdminInviteField(service identity_controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "adminInvite",
		Description: "invites a new user",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The email of the user",
			},
			"firstName": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The first name of the user",
			},
			"lastName": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The last name of the user",
			},
			"subject": &graphql.ArgumentConfig{
				Type:         graphql.NewNonNull(graphql.String),
				Description:  "The subject of the email of the invitation",
				DefaultValue: "You have been invited to join",
			},
			"link": &graphql.ArgumentConfig{
				Type:         graphql.String,
				Description:  "The link to registration form",
				DefaultValue: "/identity/register",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			err = CheckAdmin(service, &p)
			if err != nil {
				return nil, err
			}
			// from context
			// origin
			origin, err := GetOriginFromContext(&p)
			if err != nil {
				return
			}
			// params
			email := p.Args["email"].(string)
			if len(email) == 0 {
				return nil, errors.New("email is required")
			}
			firstName := p.Args["firstName"].(string)
			if len(email) == 0 {
				return nil, errors.New("first name is required")
			}
			lastName := p.Args["lastName"].(string)
			if len(email) == 0 {
				return nil, errors.New("last name is required")
			}
			subject := p.Args["subject"].(string)
			if len(email) == 0 {
				return nil, errors.New("last subject is required")
			}
			// build link
			// check if first rune is /
			link := p.Args["link"].(string)
			if len(link) == 0 {
				link = "https://" + origin + "/identity/register"
			}
			if link[0] == '/' {
				// combine origin with link
				// remove last rune if it is /
				if origin[len(origin)-1] == '/' {
					origin = origin[:len(origin)-1]
				}
				link = "https://" + origin + "/" + link
			}
			// invite new user
			err = identity_controllers.InviteUser(service, origin, subject, firstName, lastName, email, link)
			return err == nil, err
		},
	}
	return &field
}
