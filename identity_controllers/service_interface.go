package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_communication/email"
	"gorm.io/gorm"
	"net/mail"
)

type IdentityService interface {
	GetPepper() string
	GetSQLClient() *gorm.DB
	GetIssuer() string
	GetAudience() string
	GetSenderEmailAddress() mail.Address
	SendEmail(senderAddress mail.Address, receiverAddress mail.Address, subject, content string) error
	SendSMS(receiver string, content string) (err error)
	ResolveRegistrationEmailTemplate(origin, emailAddress, confirmationUrl string) email.RegistrationEmailTemplate
	ResolvePasswordResetEmailTemplate(origin, emailAddress, confirmationUrl string) email.PasswordResetTemplate
	ResolveInvitationEmailTemplate(origin, firstName, lastName, emailAddress, link string) email.InvitationEmailTemplate
	AutoClearUserAfterRegistration(origin string) bool // checks if a user should be automatically cleared after registration
}
