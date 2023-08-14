package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

type AdminMutations struct {
	AdminInvite           *graphql.Field
	AdminBlockIdentity    *graphql.Field
	AdminUnblockIdentity  *graphql.Field
	AdminLockIdentity     *graphql.Field
	AdminUnlockIdentity   *graphql.Field
	AdminDeleteIdentity   *graphql.Field
	AdminActivateIdentity *graphql.Field
	// TODO: implement
	AdminResetEmail     *graphql.Field
	AdminUpdateIdentity *graphql.Field
	AdminResetPassword  *graphql.Field
}

func InitAdminGraphQlMutations(service identity_controllers.IdentityService) *AdminMutations {
	gql := AdminMutations{
		AdminInvite:           AdminInviteField(service),
		AdminBlockIdentity:    AdminBlockIdentity(service),
		AdminUnblockIdentity:  AdminUnBlockIdentity(service),
		AdminLockIdentity:     AdminLockIdentity(service),
		AdminUnlockIdentity:   AdminUnlockIdentity(service),
		AdminDeleteIdentity:   AdminDeleteIdentity(service),
		AdminActivateIdentity: AdminActivateIdentity(service),
	}
	return &gql
}

func (gql *AdminMutations) GenerateMutationObjects(root *graphql.Object) {
	root.AddFieldConfig(gql.AdminInvite.Name, gql.AdminInvite)
	root.AddFieldConfig(gql.AdminBlockIdentity.Name, gql.AdminBlockIdentity)
	root.AddFieldConfig(gql.AdminUnblockIdentity.Name, gql.AdminUnblockIdentity)
	root.AddFieldConfig(gql.AdminLockIdentity.Name, gql.AdminLockIdentity)
	root.AddFieldConfig(gql.AdminUnlockIdentity.Name, gql.AdminUnlockIdentity)
	root.AddFieldConfig(gql.AdminDeleteIdentity.Name, gql.AdminDeleteIdentity)
	root.AddFieldConfig(gql.AdminActivateIdentity.Name, gql.AdminActivateIdentity)
}
