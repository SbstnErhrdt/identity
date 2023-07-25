package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/graphql-go/graphql"
)

// ApiTokenGraphQlModel is the graphql model for the api token
var ApiTokenGraphQlModel = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ApiToken",
	Description: "The API token of the identity",
	Fields: graphql.Fields{
		"UID": &graphql.Field{
			Type:        graphql.String,
			Description: "the unique ID of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityApiToken); ok {
					return obj.UID, nil
				}
				return nil, errors.New("can not cast api token UID")
			},
		},
		"createdAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "the creation date of the api token",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityApiToken); ok {
					return obj.CreatedAt, nil
				}
				return nil, errors.New("can not cast api token creation date")
			},
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "the name of the api token",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityApiToken); ok {
					return obj.Name, nil
				}
				return nil, errors.New("can not cast api token name")
			},
		},
		"token": &graphql.Field{
			Type:        graphql.String,
			Description: "the name of the api token",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityApiToken); ok {
					return obj.Token, nil
				}
				return nil, errors.New("can not cast api token")
			},
		},
		"expirationDate": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "the expiration dat of the api token",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.IdentityApiToken); ok {
					return obj.ExpirationDate, nil
				}
				return nil, errors.New("can not cast api token expiration date")
			},
		},
	},
})
