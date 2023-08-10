package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

type AdminMutations struct {
	AdminInvite          *graphql.Field
	AdminBlockIdentity   *graphql.Field
	AdminUnBlockIdentity *graphql.Field
	AdminClearIdentity   *graphql.Field
	AdminDeleteIdentity  *graphql.Field
	// TODO: implement
	AdminResetEmail     *graphql.Field
	AdminUpdateIdentity *graphql.Field
	AdminResetPassword  *graphql.Field
}

func InitAdminGraphQlMutations(service identity_controllers.IdentityService) *AdminMutations {
	gql := AdminMutations{
		AdminInvite:          AdminInviteField(service),
		AdminBlockIdentity:   AdminBlockIdentity(service),
		AdminUnBlockIdentity: AdminUnBlockIdentity(service),
		AdminClearIdentity:   AdminClearIdentity(service),
		AdminDeleteIdentity:  AdminDeleteIdentity(service),
	}
	return &gql
}

func (gql *AdminMutations) GenerateMutationObjects(root *graphql.Object) {
	root.AddFieldConfig(gql.AdminInvite.Name, gql.AdminInvite)
	root.AddFieldConfig(gql.AdminBlockIdentity.Name, gql.AdminBlockIdentity)
	root.AddFieldConfig(gql.AdminUnBlockIdentity.Name, gql.AdminUnBlockIdentity)
	root.AddFieldConfig(gql.AdminClearIdentity.Name, gql.AdminClearIdentity)
	root.AddFieldConfig(gql.AdminDeleteIdentity.Name, gql.AdminDeleteIdentity)
}
