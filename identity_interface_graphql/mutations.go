package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/gql_auto"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/graphql-go/graphql"
)

// Mutations is a struct that holds all the mutations
type Mutations struct {
	Login                    *graphql.Field
	Register                 *graphql.Field
	RegistrationConfirmation *graphql.Field
	UpdateIdentity           *graphql.Field
	ChangePassword           *graphql.Field
	InitResetPassword        *graphql.Field
	ConfirmResetPassword     *graphql.Field
	DeleteIdentity           *graphql.Field
	AnonymizeIdentity        *graphql.Field
	CreateApiToken           *graphql.Field
	DeleteApiToken           *graphql.Field
	InvitationConfirmation   *graphql.Field
}

func InitGraphQlMutations(service identity_controllers.IdentityService) *Mutations {
	gql := Mutations{
		Login:                    LoginField(service),
		Register:                 RegisterField(service),
		RegistrationConfirmation: RegistrationConfirmationField(service),
		UpdateIdentity:           UpdateIdentityField(service),
		ChangePassword:           ChangePasswordField(service),
		InitResetPassword:        InitResetPasswordField(service),
		ConfirmResetPassword:     ConfirmResetPasswordField(service),
		DeleteIdentity:           DeleteField(service),
		AnonymizeIdentity:        AnonymizeField(service),
		CreateApiToken:           CreateApiTokenField(service),
		DeleteApiToken:           DeleteApiTokenField(service),
		InvitationConfirmation:   InvitationConfirmationField(service),
	}
	return &gql
}

func (gql *Mutations) GenerateMutationObjects(root *graphql.Object) {
	gql_auto.AddField(root, gql.Login)
	gql_auto.AddField(root, gql.Register)
	gql_auto.AddField(root, gql.RegistrationConfirmation)
	gql_auto.AddField(root, gql.UpdateIdentity)
	gql_auto.AddField(root, gql.ChangePassword)
	gql_auto.AddField(root, gql.InitResetPassword)
	gql_auto.AddField(root, gql.ConfirmResetPassword)
	gql_auto.AddField(root, gql.DeleteIdentity)
	gql_auto.AddField(root, gql.AnonymizeIdentity)
	gql_auto.AddField(root, gql.CreateApiToken)
	gql_auto.AddField(root, gql.DeleteApiToken)
	gql_auto.AddField(root, gql.InvitationConfirmation)
}
