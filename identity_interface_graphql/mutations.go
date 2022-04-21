package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/controllers"
	"github.com/graphql-go/graphql"
)

type Mutations struct {
	Login                    *graphql.Field
	Register                 *graphql.Field
	RegistrationConfirmation *graphql.Field
	UpdateIdentity           *graphql.Field
	ChangePassword           *graphql.Field
	InitResetPassword        *graphql.Field
	ConfirmResetPassword     *graphql.Field
}

func InitMutations(service controllers.IdentityService) *Mutations {
	gql := Mutations{
		Login:                    LoginField(service),
		Register:                 RegistrationField(service),
		RegistrationConfirmation: RegistrationConfirmationField(service),
		UpdateIdentity:           UpdateIdentityField(service),
		ChangePassword:           ChangePasswordField(service),
		InitResetPassword:        InitResetPasswordField(service),
		ConfirmResetPassword:     ConfirmResetPasswordField(service),
	}
	return &gql
}

func (gql *Mutations) GenerateMutationObjects(root *graphql.Object) {
	root.AddFieldConfig("login", gql.Login)
	root.AddFieldConfig("register", gql.Register)
	root.AddFieldConfig("registrationConfirmation", gql.RegistrationConfirmation)
	root.AddFieldConfig("updateIdentity", gql.UpdateIdentity)
	root.AddFieldConfig("changePassword", gql.ChangePassword)
	root.AddFieldConfig("initResetPassword", gql.InitResetPassword)
	root.AddFieldConfig("confirmResetPassword", gql.ConfirmResetPassword)
}
