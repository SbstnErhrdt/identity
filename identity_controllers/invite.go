package identity_controllers

import (
	log "github.com/sirupsen/logrus"
	"net/mail"
)

// InviteUser invites a user to the system
func InviteUser(service IdentityService, origin, subject, firstName, lastName, emailAddress, link string) (err error) {
	logger := log.WithFields(
		log.Fields{
			"method":       "InviteUser",
			"origin":       origin,
			"firstName":    firstName,
			"lastName":     lastName,
			"emailAddress": emailAddress,
		},
	)
	// send email
	// Build email template
	logger.Debug("Build email template")
	// resolve the email template
	emailTemplate := service.ResolveInvitationEmailTemplate(origin, firstName, lastName, emailAddress, link)
	// generate the content of the email
	content, err := emailTemplate.Content()
	if err != nil {
		logger.Error(err)
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
	return
}
