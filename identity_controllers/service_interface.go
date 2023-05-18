package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_communication/email"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/mail"
	"time"
)

type IdentityService interface {
	GetLogger() *logrus.Logger
	// SQL
	GetSQLClient() *gorm.DB
	// Security
	GetPepper() string
	GetIssuer() string
	GetAudience() string
	// Admin
	GetAdminEmail() string
	// Communication
	GetSenderEmailAddress() mail.Address
	SendEmail(senderAddress mail.Address, receiverAddress mail.Address, subject, content string) error
	SendSMS(receiver string, content string) (err error)
	ResolveRegistrationEmailTemplate(origin, emailAddress, confirmationUrl string) email.RegistrationEmailTemplate
	ResolvePasswordResetEmailTemplate(origin, emailAddress, confirmationUrl string) email.PasswordResetTemplate
	ResolveInvitationEmailTemplate(origin, firstName, lastName, emailAddress, link string) email.InvitationEmailTemplate
	// Registration
	AutoClearUserAfterRegistration(origin string) bool // checks if a user should be automatically cleared after registration
	AllowRegistration(origin string) bool              // checks if a user should be automatically cleared after registration
	GetExpirationRegistration() time.Duration
	GetExpirationPasswordReset() time.Duration
}
