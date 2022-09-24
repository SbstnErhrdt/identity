package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

type AdminMutations struct {
	Invite *graphql.Field
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
