package identity_interface_graphql

import (
	"encoding/json"
	"github.com/SbstnErhrdt/identity/pkg/identity/controllers"
	"github.com/SbstnErhrdt/identity/pkg/identity/models"
	"github.com/graphql-go/graphql"
)

var IdentityMutationObject = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "IdentityMutationObject",
	Fields: graphql.InputObjectConfigFieldMap{
		"UID": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"salutation": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"firstName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"backupEmail": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"phone": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"backupPhone": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

func UpdateIdentityField(service controllers.IdentityService) *graphql.Field {
	field := graphql.Field{
		Name:        "updateIdentity",
		Description: "Update the identity of the current identity",
		Type:        IdentityGraphQlModel,
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type:        IdentityMutationObject,
				Description: "The identity",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
			// params
			jsonString, err := json.Marshal(p.Args["data"])
			if err != nil {
				return nil, err
			}
			identity := models.Identity{}
			err = json.Unmarshal(jsonString, &identity)
			if err != nil {
				return nil, err
			}
			// extract uid
			uid, err := GetUserUIDFromContext(&p)
			if err != nil {
				return nil, err
			}
			// overwrite uid
			identity.UID = uid
			// update
			err = controllers.UpdateIdentity(service, &identity)
			return &identity, err
		},
	}
	return &field
}
