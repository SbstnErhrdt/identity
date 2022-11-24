package identity_interface_graphql

import (
	"github.com/SbstnErhrdt/identity/identity_controllers"
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
	DeleteIdentity           *graphql.Field
	AnonymizeIdentity        *graphql.Field
}

func InitMutations(service identity_controllers.IdentityService) *Mutations {
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
	}
	return &gql
}

func (gql *Mutations) GenerateMutationObjects(root *graphql.Object) {
	root.AddFieldConfig(gql.Login.Name, gql.Login)
	root.AddFieldConfig(gql.Register.Name, gql.Register)
	root.AddFieldConfig(gql.RegistrationConfirmation.Name, gql.RegistrationConfirmation)
	root.AddFieldConfig(gql.UpdateIdentity.Name, gql.UpdateIdentity)
	root.AddFieldConfig(gql.ChangePassword.Name, gql.ChangePassword)
	root.AddFieldConfig(gql.InitResetPassword.Name, gql.InitResetPassword)
	root.AddFieldConfig(gql.ConfirmResetPassword.Name, gql.ConfirmResetPassword)
	root.AddFieldConfig(gql.DeleteIdentity.Name, gql.DeleteIdentity)
	root.AddFieldConfig(gql.AnonymizeIdentity.Name, gql.AnonymizeIdentity)
}
