package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

type AdminMutations struct {
	Invite        *graphql.Field
	BlockUser     *graphql.Field
	UnBlockUser   *graphql.Field
	ActivateUser  *graphql.Field
	ResetEmail    *graphql.Field
	ResetPassword *graphql.Field
	SearchUsers   *graphql.Field
	ReadUser      *graphql.Field
	UpdateUser    *graphql.Field
}

func InitAdminMutations(service identity_controllers.IdentityService) *AdminMutations {
	gql := AdminMutations{
		Invite: InviteField(service),
	}
	return &gql
}

func (gql *AdminMutations) GenerateMutationObjects(root *graphql.Object) {
	root.AddFieldConfig(gql.Invite.Name, gql.Invite)
}
