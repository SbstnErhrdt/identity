package identity_controllers

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/mail"
)

// InviteUser invites a user to the system
func InviteUser(service IdentityService, mandateUID uuid.UUID, clientUID *uuid.UUID, orgName, subject, firstName, lastName, emailAddress, link string) (err error) {
	logger := log.WithFields(
		log.Fields{
			"method":       "InviteUser",
			"mandateUID":   mandateUID,
			"clientUID":    clientUID,
			"firstName":    firstName,
			"lastName":     lastName,
			"emailAddress": emailAddress,
		},
	)
	// send email
	// Build email template
	logger.Debug("Build email template")
	// resolve the email template
	emailTemplate := service.ResolveInvitationEmailTemplate(mandateUID, clientUID, orgName, firstName, lastName, emailAddress, link)
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
