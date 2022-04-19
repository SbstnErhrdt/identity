package controllers

import (
	"github.com/SbstnErhrdt/identity/communication/email"
	"github.com/SbstnErhrdt/identity/services"
	"gorm.io/gorm"
	"net/mail"
)

type TestIdentityService struct{}

func (t TestIdentityService) ResolvePasswordResetEmailTemplate(origin, emailAddress, confirmationUrl string) email.PasswordResetTemplate {
	//TODO implement me
	panic("implement me")
}

func (t TestIdentityService) ResolveRegistrationEmailTemplate(origin, emailAddress, confirmationUrl string) email.RegistrationEmailTemplate {
	//TODO implement me
	panic("implement me")
}

func (t TestIdentityService) GetAuthConfirmationEndpoint() string {
	panic("implement me")
}

func (t TestIdentityService) GetSenderEmailAddress() mail.Address {
	panic("implement me")
}

func (t TestIdentityService) SendEmail(senderAddress mail.Address, receiverAddress mail.Address, subject, content string) error {
	return services.SendEmail(senderAddress, receiverAddress, subject, content)
}

func (t TestIdentityService) GetEmailTemplate() email.GlobalTemplate {
	return email.GlobalTemplate{
		PrimaryColor:        "#333",
		PrimaryBorderColor:  "#000",
		HeaderLogoUrl:       "",
		IntroText:           "Helloo",
		OutroText:           "Goodbye",
		FooterCopyrightText: "Test Copyright",
		UnsubscribeLink:     "https://erhardt.net",
	}
}

func (t TestIdentityService) GetSQLClient() *gorm.DB {
	panic("implement me")
}

func (t TestIdentityService) GetIssuer() string {
	panic("implement me")
}

func (t TestIdentityService) GetAudience() string {
	panic("implement me")
}

func (t TestIdentityService) SendSMS(address string, content string) error {
	panic("implement me")
}
