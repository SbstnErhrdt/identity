package identity

import (
	"github.com/SbstnErhrdt/identity/communication/email"
	"github.com/SbstnErhrdt/identity/identity_interface_graphql"
	"github.com/SbstnErhrdt/identity/services"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
	"net/mail"
)

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

// Service is the identity service
type Service struct {
	Issuer                     string
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
}

// NewService inits a new identity service
func NewService(issuer string, senderEmailAddress mail.Address) *Service {
	s := Service{
		Issuer:                    issuer,
		senderEmailAddress:        senderEmailAddress,
		PrimaryIdentificationType: EmailIdentificationType,
		sendEmail:                 services.SendEmail,
		// fallback services
		registrationEmailResolver:  email.DefaultRegistrationEmailResolver,
		passwordResetEmailResolver: email.DefaultPasswordResetEmailResolver,
	}
	return &s
}

// SetGraphQLQueryInterface sets the graphql query interface
func (s *Service) SetGraphQLQueryInterface(rootQueryObject *graphql.Object) *Service {
	s.gqlRootObject = rootQueryObject
	// init queries
	q := identity_interface_graphql.InitGraphQlQueries(s)
	// connect to root query object
	q.GenerateQueryObjects(s.gqlRootObject)
	return s
}

// SetGraphQLMutationInterface sets the graphql mutation interface
func (s *Service) SetGraphQLMutationInterface(rootMutationObject *graphql.Object) *Service {
	s.gqlRootMutationObject = rootMutationObject
	// init mutations
	q := identity_interface_graphql.InitMutations(s)
	// connect to root mutation object
	q.GenerateMutationObjects(rootMutationObject)
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

// GetSQLClient returns the sql client
func (s *Service) GetSQLClient() *gorm.DB {
	return s.sqlClient
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
