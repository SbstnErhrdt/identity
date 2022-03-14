package identity

import (
	"github.com/SbstnErhrdt/identity/pkg/identity/communication/email"
	"github.com/SbstnErhrdt/identity/pkg/identity/identity_interface_graphql"
	"github.com/SbstnErhrdt/identity/pkg/identity/services"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
	"net/mail"
)

type IdentificationType string

const (
	EmailIdentificationType IdentificationType = "email"
	PhoneIdentificationType IdentificationType = "phone"
)

type ResolveRegistrationEmailTemplate func(origin, emailAddress, token string) email.RegistrationEmailTemplate
type ResolvePasswordResetEmailTemplate func(origin, emailAddress, token string) email.PasswordResetTemplate

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

func (s *Service) SetGraphQLQueryInterface(rootQueryObject *graphql.Object) *Service {
	s.gqlRootObject = rootQueryObject
	// init queries
	q := identity_interface_graphql.InitGraphQlQueries(s)
	// connect to root query object
	q.GenerateQueryObjects(s.gqlRootObject)
	return s
}

func (s *Service) SetGraphQLMutationInterface(rootMutationObject *graphql.Object) *Service {
	s.gqlRootMutationObject = rootMutationObject
	// init mutations
	q := identity_interface_graphql.InitMutations(s)
	// connect to root mutation object
	q.GenerateMutationObjects(rootMutationObject)
	return s
}

func (s *Service) SetSQLClient(client *gorm.DB) *Service {
	s.sqlClient = client
	return s
}

func (s *Service) SetAuthConfirmationEndpoint(authConfirmationEndpoint string) *Service {
	s.authConfirmationEndpoint = authConfirmationEndpoint
	return s
}

func (s *Service) SetRegistrationEmailResolver(fn ResolveRegistrationEmailTemplate) *Service {
	s.registrationEmailResolver = fn
	return s
}

func (s *Service) GetSQLClient() *gorm.DB {
	return s.sqlClient
}

func (s *Service) GetIssuer() string {
	return s.Issuer
}

func (s *Service) GetSenderEmailAddress() mail.Address {
	return s.senderEmailAddress
}

func (s *Service) GetAudience() string {
	return s.Audience
}

func (s *Service) SendEmail(senderAddress mail.Address, receiverAddress mail.Address, subject, content string) error {
	return s.sendEmail(senderAddress, receiverAddress, subject, content)
}

func (s *Service) SendSMS(address string, content string) error {
	panic("implement me")
}

func (s *Service) GetEmailTemplate() email.GlobalTemplate {
	return s.emailTemplate
}

func (s *Service) ResolveRegistrationEmailTemplate(origin, emailAddress, confirmationUrl string) email.RegistrationEmailTemplate {
	return s.registrationEmailResolver(origin, emailAddress, confirmationUrl)
}

func (s *Service) ResolvePasswordResetEmailTemplate(origin, emailAddress, confirmationUrl string) email.PasswordResetTemplate {
	return s.passwordResetEmailResolver(origin, emailAddress, confirmationUrl)
}
