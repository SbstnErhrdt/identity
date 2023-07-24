package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// Queries is a struct containing all the graphql queries
type Queries struct {
	CurrentIdentity *graphql.Field
	ApiTokens       *graphql.Field
}

// InitGraphQlQueries initializes the graphql queries
func InitGraphQlQueries(service identity_controllers.IdentityService) *Queries {
	gql := Queries{
		CurrentIdentity: CurrentIdentityField(service),
		ApiTokens:       ApiTokensField(service),
	}
	return &gql
}

// GenerateQueryObjects generates the query objects
func (gql *Queries) GenerateQueryObjects(root *graphql.Object) {
	root.AddFieldConfig(gql.CurrentIdentity.Name, gql.CurrentIdentity)
}
