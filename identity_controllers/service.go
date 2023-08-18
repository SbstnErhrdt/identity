package identity_controllers

import (
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_communication/email"
	"github.com/SbstnErhrdt/identity/services"
	"gorm.io/gorm"
	"log/slog"
	"net/mail"
	"os"
	"time"
)

// IdentificationType is the type of the identification
type IdentificationType string

func init() {
	if os.Getenv("SECURITY_PEPPER") == "PEPPER" {
		slog.Warn("please change the pepper value in the environment variable SECURITY_PEPPER")
	}
}

const (
	// EmailIdentificationType is the type of identification that is used to identify a user by email
	EmailIdentificationType IdentificationType = "email"
	// PhoneIdentificationType is the type of identification that is used to identify a user by phone number
	PhoneIdentificationType IdentificationType = "phone"
)

// ResolveRegistrationEmailTemplate resolves the registration email template
type ResolveRegistrationEmailTemplate func(origin, emailAddress, token string) email.RegistrationEmailTemplate

// ResolvePasswordResetEmailTemplate resolves the password reset email template
type ResolvePasswordResetEmailTemplate func(origin, emailAddress, token string) email.PasswordResetTemplate

// ResolveInvitationEmailTemplate resolves the invitation email template
// sends a link to the user to register
type ResolveInvitationEmailTemplate func(origin, firstName, lastName, emailAddress, content, link string) email.InvitationEmailTemplate

// ResolveCreationInvitationEmailTemplate resolves the creation invitation email template
// sends a link to the newly created identity to select a password and confirm the email address and the terms of service
type ResolveCreationInvitationEmailTemplate func(origin, firstName, lastName, emailAddress, content, token string) email.CreationInvitationEmailTemplate

// ClearUserFn clears a user
type ClearUserFn func(origin string) bool

// SendMailFn is the function that sends an email
type SendMailFn func(mail.Address, mail.Address, string, string) error

// AllowRegistrationFn checks if users can register
type AllowRegistrationFn func(origin string) bool

// AutoClearUserAfterRegistration automatically clears a user
func AutoClearUserAfterRegistration(origin string) bool {
	return true
}

// AutoBlockUserFn automatically clears a user
func AutoBlockUserFn(origin string) bool {
	return false
}

// ControllerService is the identity service
type ControllerService struct {
	logger                          *slog.Logger
	adminEmail                      string
	Issuer                          string
	Pepper                          string
	Audience                        string
	PrimaryIdentificationType       IdentificationType
	sqlClient                       *gorm.DB
	senderEmailAddress              mail.Address // default email
	sendEmail                       SendMailFn
	sendSMS                         func(string, string) error
	emailTemplate                   email.GlobalTemplate
	authConfirmationEndpoint        string
	registrationEmailResolver       ResolveRegistrationEmailTemplate
	passwordResetEmailResolver      ResolvePasswordResetEmailTemplate
	invitationEmailResolver         ResolveInvitationEmailTemplate
	creationInvitationEmailResolver ResolveCreationInvitationEmailTemplate
	clearUserAfterRegistration      ClearUserFn
	allowRegistration               AllowRegistrationFn
	expirationRegistration          time.Duration
	expirationPasswordReset         time.Duration
}

func (s *ControllerService) GetLogger() *slog.Logger {
	return s.logger
}

// NewService inits a new identity service
func NewService(issuer string, senderEmailAddress mail.Address) *ControllerService {
	s := ControllerService{
		logger:                    slog.Default(),
		Issuer:                    issuer,
		Pepper:                    env.FallbackEnvVariable("SECURITY_PEPPER", "PEPPER"),
		senderEmailAddress:        senderEmailAddress,
		PrimaryIdentificationType: EmailIdentificationType,
		sendEmail:                 services.SendEmail,
		// fallback services
		registrationEmailResolver:       email.DefaultRegistrationEmailResolver,
		passwordResetEmailResolver:      email.DefaultPasswordResetEmailResolver,
		invitationEmailResolver:         email.DefaultInvitationEmailResolver,
		creationInvitationEmailResolver: email.DefaultCreationInvitationEmailResolver,
		clearUserAfterRegistration:      AutoClearUserAfterRegistration,
		allowRegistration:               DefaultAllowRegistration,
		// default values
		expirationRegistration:  24 * time.Hour,
		expirationPasswordReset: 3 * time.Hour,
	}
	return &s
}

// SetLogger sets the logger
func (s *ControllerService) SetLogger(logger *slog.Logger) *ControllerService {
	s.logger = logger
	return s
}

// GetAdminEmail returns the admin email
func (s *ControllerService) GetAdminEmail() string {
	return s.adminEmail
}

// SetAdminEmail sets the admin email
func (s *ControllerService) SetAdminEmail(email string) *ControllerService {
	s.adminEmail = email
	return s
}

func (s ControllerService) GetExpirationRegistration() time.Duration {
	return s.expirationRegistration
}

func (s ControllerService) GetExpirationPasswordReset() time.Duration {
	return s.expirationPasswordReset
}

// SetSQLClient sets the sql client
func (s *ControllerService) SetSQLClient(client *gorm.DB) *ControllerService {
	s.sqlClient = client
	return s
}

// SetClearUserAfterRegistrationResolver sets clear after registration resolver
func (s *ControllerService) SetClearUserAfterRegistrationResolver(fn ClearUserFn) *ControllerService {
	s.clearUserAfterRegistration = fn
	return s
}

// AutoClearUserAfterRegistration checks if a user is automatically cleared after registration
func (s *ControllerService) AutoClearUserAfterRegistration(origin string) bool {
	return s.clearUserAfterRegistration(origin)
}

// SetRegistrationEmailResolver sets the registration email resolver
func (s *ControllerService) SetRegistrationEmailResolver(fn ResolveRegistrationEmailTemplate) *ControllerService {
	s.registrationEmailResolver = fn
	return s
}

// SetPepper sets the pepper
func (s *ControllerService) SetPepper(pepper string) {
	s.Pepper = pepper
}

// GetSQLClient returns the sql client
func (s *ControllerService) GetSQLClient() *gorm.DB {
	return s.sqlClient
}

// GetPepper returns the pepper
func (s *ControllerService) GetPepper() string {
	return s.Pepper
}

// GetIssuer returns the issuer
func (s *ControllerService) GetIssuer() string {
	return s.Issuer
}

// GetSenderEmailAddress returns the sender email address
func (s *ControllerService) GetSenderEmailAddress() mail.Address {
	return s.senderEmailAddress
}

// GetAudience returns the audience
func (s *ControllerService) GetAudience() string {
	return s.Audience
}

// SendEmail sends an email
func (s *ControllerService) SendEmail(senderAddress mail.Address, receiverAddress mail.Address, subject, content string) error {
	return s.sendEmail(senderAddress, receiverAddress, subject, content)
}

// SendSMS sends an sms
func (s *ControllerService) SendSMS(address string, content string) error {
	// todo
	panic("implement me")
}

// GetEmailTemplate returns the email template
func (s *ControllerService) GetEmailTemplate() email.GlobalTemplate {
	return s.emailTemplate
}

// ResolveRegistrationEmailTemplate returns the registration email template
func (s *ControllerService) ResolveRegistrationEmailTemplate(origin, emailAddress, confirmationUrl string) email.RegistrationEmailTemplate {
	return s.registrationEmailResolver(origin, emailAddress, confirmationUrl)
}

// ResolvePasswordResetEmailTemplate returns the password reset email template
func (s *ControllerService) ResolvePasswordResetEmailTemplate(origin, emailAddress, confirmationUrl string) email.PasswordResetTemplate {
	return s.passwordResetEmailResolver(origin, emailAddress, confirmationUrl)
}

// ResolveInvitationEmailTemplate returns the invitation email template
func (s *ControllerService) ResolveInvitationEmailTemplate(origin, firstName, lastName, emailAddress, content, link string) email.InvitationEmailTemplate {
	return s.invitationEmailResolver(origin, firstName, lastName, emailAddress, content, link)
}

func (s *ControllerService) ResolveCreationInvitationEmailTemplate(origin, firstName, lastName, emailAddress, content, token string) email.CreationInvitationEmailTemplate {
	return s.creationInvitationEmailResolver(origin, firstName, lastName, emailAddress, content, token)
}

// AllowRegistration checks if users can register
func (s *ControllerService) AllowRegistration(origin string) bool {
	return s.allowRegistration(origin)
}

func DefaultAllowRegistration(origin string) bool {
	return true
}
