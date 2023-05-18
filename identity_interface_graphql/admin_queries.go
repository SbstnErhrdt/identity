package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

type AdminQueries struct {
	SearchUsers *graphql.Field
	User        *graphql.Field
}

func InitAdminGraphQlQueries(service identity_controllers.IdentityService) *AdminQueries {
	gql := AdminQueries{
		SearchUsers: IdentitiesSearchField(service),
		User:        Identity(service),
	}
	return &gql
}

func (gql *AdminQueries) GenerateQueryObjects(root *graphql.Object) {
	root.AddFieldConfig(gql.SearchUsers.Name, gql.SearchUsers)
	root.AddFieldConfig(gql.User.Name, gql.User)
}
