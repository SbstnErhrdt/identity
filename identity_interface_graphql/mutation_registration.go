package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/controllers"
	"github.com/graphql-go/graphql"
)

func RegistrationField(service controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "Registration",
		Description: "Submit identity and password to retrieve a token",
		Type:        graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"identity": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The identity",
			},
			"password": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The password",
			},
			"passwordConfirmation": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The password confirmation",
			},
			"acceptTermsAndConditions": &graphql.ArgumentConfig{
				Type:        graphql.NewNonNull(graphql.Boolean),
				Description: "The confirmation of the acceptance of the terms and conditions",
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
			// ip
			origin, ok := p.Context.Value("ORIGIN").(string)
			if !ok {
				err = errors.New("can not extract origin from context")
				return
			}
			// parameters
			identity := p.Args["identity"].(string)
			password := p.Args["password"].(string)
			passwordConfirmation := p.Args["passwordConfirmation"].(string)
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
			// get token of registration
			err = controllers.Register(service, identity, password, termsAndConditions, userAgent, ip, origin)
			return err == nil, err
		},
	}
	return &field
}