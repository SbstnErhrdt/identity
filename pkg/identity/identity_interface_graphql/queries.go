package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/pkg/identity/controllers"
	"github.com/graphql-go/graphql"
)

type Queries struct {
	CurrentIdentity *graphql.Field
}

func InitGraphQlQueries(service controllers.IdentityService) *Queries {
	gql := Queries{
		CurrentIdentity: CurrentIdentityField(service),
	}
	return &gql
}

func (gql *Queries) GenerateQueryObjects(root *graphql.Object) {
	root.AddFieldConfig("Identity", gql.CurrentIdentity)
}
