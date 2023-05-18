package identity_interface_graphql

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/graphql-go/graphql"
)

// AdminIdentityGraphQlModel is the identity model for the GraphQL admin interface
var AdminIdentityGraphQlModel = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AdminIdentity",
	Description: "The administration identity",
	Fields: graphql.Fields{
		"UID": &graphql.Field{
			Type:        graphql.String,
			Description: "the unique ID of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.UID, nil
				}
				return nil, errors.New("can not cast UID object")
			},
		},
		"acceptConditionsAndPrivacy": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "accept conditions and privacy",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.AcceptConditionsAndPrivacy, nil
				}
				return nil, errors.New("can not cast UID object")
			},
		},
		"active": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "has activated the account through email",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.Active, nil
				}
				return nil, errors.New("can not cast UID object")
			},
		},
		"cleared": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "has been cleared by an admin or automatically depending on the settings",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.Cleared, nil
				}
				return nil, errors.New("can not cast UID object")
			},
		},
		"blocked": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "has been blocked by an admin",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.Cleared, nil
				}
				return nil, errors.New("can not cast UID object")
			},
		},
		"createdAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "the creation date of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.CreatedAt, nil
				}
				return nil, errors.New("can not cast UID object")
			},
		},
		"deletedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "the deletion date of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.DeletedAt, nil
				}
				return nil, errors.New("can not cast UID object")
			},
		},
		"updatedAt": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "the update date of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.UpdatedAt, nil
				}
				return nil, errors.New("can not cast UID object")
			},
		},
		"salutation": &graphql.Field{
			Type:        graphql.String,
			Description: "the salutation of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.Salutation, nil
				}
				return nil, errors.New("can not cast email object")
			},
		},
		"firstName": &graphql.Field{
			Type:        graphql.String,
			Description: "the first name of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.FirstName, nil
				}
				return nil, errors.New("can not cast email object")
			},
		},
		"lastName": &graphql.Field{
			Type:        graphql.String,
			Description: "the last name of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.LastName, nil
				}
				return nil, errors.New("can not cast email object")
			},
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "the email of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.Email, nil
				}
				return nil, errors.New("can not cast email object")
			},
		},
		"backupEmail": &graphql.Field{
			Type:        graphql.Int,
			Description: "the backup email of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.BackupEmail, nil
				}
				return nil, errors.New("can not cast backupEmail object")
			},
		},
		"phone": &graphql.Field{
			Type:        graphql.String,
			Description: "the phone number of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.Phone, nil
				}
				return nil, errors.New("can not cast phone object")
			},
		},
		"backupPhone": &graphql.Field{
			Type:        graphql.Int,
			Description: "the backup phone of the identity",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(*identity_models.Identity); ok {
					return obj.BackupPhone, nil
				}
				return nil, errors.New("can not cast backupPhone object")
			},
		},
	},
})
