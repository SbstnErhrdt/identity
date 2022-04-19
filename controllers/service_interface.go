package controllers

import (
	"github.com/SbstnErhrdt/identity/communication/email"
	"gorm.io/gorm"
	"net/mail"
)

type IdentityService interface {
	GetSQLClient() *gorm.DB
	GetIssuer() string
	GetAudience() string
	GetSenderEmailAddress() mail.Address
	SendEmail(senderAddress mail.Address, receiverAddress mail.Address, subject, content string) error
	SendSMS(receiver string, content string) (err error)
	ResolveRegistrationEmailTemplate(origin, emailAddress, confirmationUrl string) email.RegistrationEmailTemplate
	ResolvePasswordResetEmailTemplate(origin, emailAddress, confirmationUrl string) email.PasswordResetTemplate
}
