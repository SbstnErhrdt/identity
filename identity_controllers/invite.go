package identity_controllers

import (
	"net/mail"
)

// InviteUser invites a user to the system
func InviteUser(service IdentityService, origin, subject, firstName, lastName, emailAddress, link string) (err error) {
	logger := service.GetLogger().With(
		"method", "InviteUser",
		"origin", origin,
		"firstName", firstName,
		"lastName", lastName,
		"emailAddress", emailAddress,
	)
	// send email
	// Build email template
	logger.Debug("Build email template")
	// resolve the email template
	emailTemplate := service.ResolveInvitationEmailTemplate(origin, firstName, lastName, emailAddress, link)
	// generate the content of the email
	content, err := emailTemplate.Content()
	if err != nil {
		logger.With("err", err).Error("could not generate email content")
		return
	}
	// Send email
	logger.Debug("Send invitation email")
	err = service.SendEmail(
		service.GetSenderEmailAddress(),
		mail.Address{
			Name:    emailAddress,
			Address: emailAddress,
		},
		subject,
		content)
	if err != nil {
		logger.With("err", err).Error("could not send email")
		return
	}
	return
}
