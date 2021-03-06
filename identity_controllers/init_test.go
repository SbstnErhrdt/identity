package identity_controllers

import (
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_communication/email"
	"github.com/SbstnErhrdt/identity/services"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/mail"
)

var s *Service

func init() {
	addr, _ := mail.ParseAddress("no-reply@rise-workshop.com")
	s = NewTestService("RISE", *addr)
}

// IdentificationType is the type of the identification
type IdentificationType string

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
type ResolveInvitationEmailTemplate func(mandateUID uuid.UUID, clientUID *uuid.UUID, orgName, firstName, lastName, emailAddress, link string) email.InvitationEmailTemplate

// Service is the identity service
type Service struct {
	Issuer                     string
	Pepper                     string
	Audience                   string
	PrimaryIdentificationType  IdentificationType
	sqlClient                  *gorm.DB
	gqlRootObject              *graphql.Object
	gqlRootMutationObject      *graphql.Object
	senderEmailAddress         mail.Address // default email
	sendEmail                  func(mail.Address, mail.Address, string, string) error
	sendSMS                    func(string, string) error
	emailTemplate              email.GlobalTemplate
	authConfirmationEndpoint   string
	registrationEmailResolver  ResolveRegistrationEmailTemplate
	passwordResetEmailResolver ResolvePasswordResetEmailTemplate
	invitationEmailResolver    ResolveInvitationEmailTemplate
}

// NewTestService inits a new identity service
func NewTestService(issuer string, senderEmailAddress mail.Address) *Service {
	s := Service{
		Issuer:                    issuer,
		Pepper:                    env.FallbackEnvVariable("SECURITY_PEPPER", "PEPPER"),
		senderEmailAddress:        senderEmailAddress,
		PrimaryIdentificationType: EmailIdentificationType,
		sendEmail:                 services.SendEmail,
		// fallback services
		registrationEmailResolver:  email.DefaultRegistrationEmailResolver,
		passwordResetEmailResolver: email.DefaultPasswordResetEmailResolver,
		invitationEmailResolver:    email.DefaultInvitationEmailResolver,
	}
	if s.Pepper == "PEPPER" {
		log.Warn("please change the pepper value in the environment variable SECURITY_PEPPER")
	}
	return &s
}

// SetGraphQLQueryInterface sets the graphql query interface
func (s *Service) SetGraphQLQueryInterface(rootQueryObject *graphql.Object) *Service {
	return s
}

// SetGraphQLMutationInterface sets the graphql mutation interface
func (s *Service) SetGraphQLMutationInterface(rootMutationObject *graphql.Object) *Service {
	return s
}

// SetSQLClient sets the sql client
func (s *Service) SetSQLClient(client *gorm.DB) *Service {
	s.sqlClient = client
	return s
}

// SetAuthConfirmationEndpoint sets the auth confirmation endpoint
func (s *Service) SetAuthConfirmationEndpoint(authConfirmationEndpoint string) *Service {
	s.authConfirmationEndpoint = authConfirmationEndpoint
	return s
}

// SetRegistrationEmailResolver sets the registration email resolver
func (s *Service) SetRegistrationEmailResolver(fn ResolveRegistrationEmailTemplate) *Service {
	s.registrationEmailResolver = fn
	return s
}

// SetInvitationEmailResolver sets the registration email resolver
func (s *Service) SetInvitationEmailResolver(fn ResolveInvitationEmailTemplate) *Service {
	s.invitationEmailResolver = fn
	return s
}

// SetPepper sets the pepper
func (s *Service) SetPepper(pepper string) {
	s.Pepper = pepper
}

// GetSQLClient returns the sql client
func (s *Service) GetSQLClient() *gorm.DB {
	return s.sqlClient
}

// GetPepper returns the pepper
func (s *Service) GetPepper() string {
	return s.Pepper
}

// GetIssuer returns the issuer
func (s *Service) GetIssuer() string {
	return s.Issuer
}

// GetSenderEmailAddress returns the sender email address
func (s *Service) GetSenderEmailAddress() mail.Address {
	return s.senderEmailAddress
}

// GetAudience returns the audience
func (s *Service) GetAudience() string {
	return s.Audience
}

// SendEmail sends an email
func (s *Service) SendEmail(senderAddress mail.Address, receiverAddress mail.Address, subject, content string) error {
	return s.sendEmail(senderAddress, receiverAddress, subject, content)
}

// SendSMS sends an sms
func (s *Service) SendSMS(address string, content string) error {
	// todo
	panic("implement me")
}

// GetEmailTemplate returns the email template
func (s *Service) GetEmailTemplate() email.GlobalTemplate {
	return s.emailTemplate
}

// ResolveRegistrationEmailTemplate returns the registration email template
func (s *Service) ResolveRegistrationEmailTemplate(origin, emailAddress, confirmationUrl string) email.RegistrationEmailTemplate {
	return s.registrationEmailResolver(origin, emailAddress, confirmationUrl)
}

// ResolvePasswordResetEmailTemplate returns the password reset email template
func (s *Service) ResolvePasswordResetEmailTemplate(origin, emailAddress, confirmationUrl string) email.PasswordResetTemplate {
	return s.passwordResetEmailResolver(origin, emailAddress, confirmationUrl)
}

// ResolveInvitationEmailTemplate returns the invitation email template
func (s *Service) ResolveInvitationEmailTemplate(mandateUID uuid.UUID, clientUID *uuid.UUID, orgName, firstName, lastName, emailAddress, link string) email.InvitationEmailTemplate {
	return s.invitationEmailResolver(mandateUID, clientUID, orgName, firstName, lastName, emailAddress, link)
}
