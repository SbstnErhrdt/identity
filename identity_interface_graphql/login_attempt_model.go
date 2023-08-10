package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// LoginAttemptGraphQlModel is the graphql model for the login attempts
var LoginAttemptGraphQlModel = graphql.NewObject(graphql.ObjectConfig{
	Name:        "LoginAttempt",
	Description: "The login attempt of identities",
	Fields: graphql.Fields{
		"UID": &graphql.Field{
			Type:        graphql.String,
			Description: "the unique ID of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityLogin); ok {
					return obj.UID, nil
				}
				return nil, errors.New("can not cast login attempt UID")
			},
		},
		"createdAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "the creation date of the login attempt",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityLogin); ok {
					return obj.CreatedAt, nil
				}
				return nil, errors.New("can not cast login attempt creation date")
			},
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "the name of the login attempt",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityLogin); ok {
					return obj.Email, nil
				}
				return nil, errors.New("can not cast login attempt email")
			},
		},
		"origin": &graphql.Field{
			Type:        graphql.String,
			Description: "the origin of the login attempt",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityLogin); ok {
					return obj.Origin, nil
				}
				return nil, errors.New("can not cast login attempt origin")
			},
		},
		"userAgent": &graphql.Field{
			Type:        graphql.String,
			Description: "the user agent of the login attempt",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityLogin); ok {
					return obj.UserAgent, nil
				}
				return nil, errors.New("can not cast login attempt user agent")
			},
		},
		"ip": &graphql.Field{
			Type:        graphql.String,
			Description: "the ip of the login attempt",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityLogin); ok {
					return obj.IP, nil
				}
				return nil, errors.New("can not cast login attempt ip")
			},
		},
		"identityUID": &graphql.Field{
			Type:        graphql.String,
			Description: "the identity UID of the login attempt",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityLogin); ok {
					return obj.IdentityUID, nil
				}
				return nil, errors.New("can not cast login attempt identity UID")
			},
		},
		"success": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "the success of the login attempt",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityLogin); ok {
					if obj.IdentityUID == nil {
						return false, nil
					}
					return *obj.IdentityUID != uuid.Nil, nil
				}
				return nil, errors.New("can not cast login attempt identity UID")
			},
		},
	},
})
