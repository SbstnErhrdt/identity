package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/gql_auto"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// AdminMutations is a struct that holds all the admin mutations
type AdminMutations struct {
	AdminInvite                  *graphql.Field
	AdminCreateIdentityAndInvite *graphql.Field
	AdminBlockIdentity           *graphql.Field
	AdminUnblockIdentity         *graphql.Field
	AdminLockIdentity            *graphql.Field
	AdminUnlockIdentity          *graphql.Field
	AdminDeleteIdentity          *graphql.Field
	AdminActivateIdentity        *graphql.Field
	// TODO: implement
	AdminResetEmail     *graphql.Field
	AdminUpdateIdentity *graphql.Field
	AdminResetPassword  *graphql.Field
}

// InitAdminGraphQlMutations initializes the admin mutations
func InitAdminGraphQlMutations(service identity_controllers.IdentityService) *AdminMutations {
	gql := AdminMutations{
		AdminInvite:                  AdminInviteField(service),
		AdminCreateIdentityAndInvite: AdminCreateIdentityAndInviteField(service),
		AdminBlockIdentity:           AdminBlockIdentity(service),
		AdminUnblockIdentity:         AdminUnBlockIdentity(service),
		AdminLockIdentity:            AdminLockIdentity(service),
		AdminUnlockIdentity:          AdminUnlockIdentity(service),
		AdminDeleteIdentity:          AdminDeleteIdentity(service),
		AdminActivateIdentity:        AdminActivateIdentity(service),
	}
	return &gql
}

// GenerateMutationObjects generates the mutation objects
func (gql *AdminMutations) GenerateMutationObjects(root *graphql.Object) {
	gql_auto.AddField(root, gql.AdminInvite)
	gql_auto.AddField(root, gql.AdminCreateIdentityAndInvite)
	gql_auto.AddField(root, gql.AdminBlockIdentity)
	gql_auto.AddField(root, gql.AdminUnblockIdentity)
	gql_auto.AddField(root, gql.AdminLockIdentity)
	gql_auto.AddField(root, gql.AdminUnlockIdentity)
	gql_auto.AddField(root, gql.AdminDeleteIdentity)
	gql_auto.AddField(root, gql.AdminActivateIdentity)
}
