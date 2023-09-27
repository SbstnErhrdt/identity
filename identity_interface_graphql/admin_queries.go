package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/gql_auto"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

type AdminQueries struct {
	SearchIdentities  *graphql.Field
	Identity          *graphql.Field
	LoginAttempts     *graphql.Field
	IdentityApiTokens *graphql.Field
}

func InitAdminGraphQlQueries(service identity_controllers.IdentityService) *AdminQueries {
	gql := AdminQueries{
		SearchIdentities:  IdentitiesSearchField(service),
		Identity:          IdentityField(service),
		LoginAttempts:     LoginAttemptsField(service),
		IdentityApiTokens: IdentityApiTokensField(service),
	}
	return &gql
}

func (gql *AdminQueries) GenerateQueryObjects(root *graphql.Object) {
	gql_auto.AddField(root, gql.SearchIdentities)
	gql_auto.AddField(root, gql.Identity)
	gql_auto.AddField(root, gql.LoginAttempts)
	gql_auto.AddField(root, gql.IdentityApiTokens)
}
