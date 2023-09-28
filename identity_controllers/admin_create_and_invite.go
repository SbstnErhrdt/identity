package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/SbstnErhrdt/identity/security"
	"github.com/google/uuid"
	"net/mail"
	"time"
)

// InvitationTimeout is the time after which an invitation expires
var InvitationTimeout = time.Hour * 24 * 30 * 6 // 6 months

// AdminCreateIdentityAndInvite creates an identity and sends an invitation
// - create identity
// - create confirmation
// - reset password
// - send invite
func AdminCreateIdentityAndInvite(service IdentityService, origin, subject, content, firstName, lastName, emailAddress string) (identity identity_models.Identity, err error) {
	// sanitize input
	emailAddress = SanitizeEmail(emailAddress)
	// init logger
	logger := service.GetLogger().With(
		"email", emailAddress,
		"process", "AdminCreateIdentityAndInvite",
	)
	// checks if users can register
	if !service.AllowRegistration(origin) {
		err = ErrRegistrationIsNotAllowed
		logger.With("err", err).Warn("registration is not allowed")
		return
	}
	// check login data
	if len(emailAddress) == 0 {
		err = ErrNoEmail
		logger.With("err", err).Error("no email address")
		return
	}
	// check if user exits already
	logger.Debug("CheckIfEmailIsFree")
	isFree, err := CheckIfEmailIsFree(service, emailAddress)
	if !isFree {
		err = ErrEmailAlreadyExists
		logger.With("err", err).Warn("email already exists")
		// OWASP recommends not to return an in case the account already exists
		return
	}
	if err != nil {
		logger.With("err", err).Error("cannot check if email is free")
		// OWASP recommends not to return an in case the account already exists
		return
	}
	// generate random registration token
	logger.Debug("GenerateRandomString")
	randomToken, err := security.GenerateRandomString(64)
	if err != nil {
		logger.With("err", err).Error("could not generate random string")
		return
	}
	// create user
	identity = identity_models.Identity{
		Email:                      emailAddress,
		AcceptConditionsAndPrivacy: false,
		FirstName:                  firstName,
		LastName:                   lastName,
		Cleared:                    true, // cleared by admin
	}
	identity.UID = uuid.New()
	// set password
	// generate random password
	randomPassword, err := security.GenerateRandomString(64)
	if err != nil {
		logger.With("err", err).Error("could not generate random string")
	}
	logger.Debug("SetNewPassword")
	err = identity.SetNewPassword(service.GetPepper(), randomPassword)
	if err != nil {
		logger.With("err", err).Error("could not set new password")
		err = ErrGenericRegistration
		return
	}

	// init transaction
	tx := service.GetSQLClient().Begin()

	// create user
	logger.Debug("create identity in the database")
	tx.Create(&identity)

	// create confirmation
	confirmation := identity_models.IdentityRegistrationConfirmation{
		Token:                 randomToken, // same token as reset password
		IdentityUID:           identity.UID,
		ExpiredAt:             time.Now().UTC().Add(InvitationTimeout),
		RegistrationIP:        "invited by admin",
		RegistrationUserAgent: "invited by admin",
	}
	logger.Debug("create confirmation token in the database")
	tx.Create(&confirmation)

	// create reset password
	resetPassword := identity_models.IdentityResetPassword{
		IdentityUID: identity.UID,
		Token:       randomToken, // same token as confirmation
		Expire:      time.Now().UTC().Add(InvitationTimeout),
		UserAgent:   "invited by admin",
		IP:          "invited by admin",
	}
	tx.Create(&resetPassword)

	err = tx.Commit().Error
	if err != nil {
		logger.With("err", err).Error("")
		return
	}

	// Build email template
	logger.Debug("Build email template")

	// resolve the email template
	emailTemplate := service.ResolveCreationInvitationEmailTemplate(origin, firstName, lastName, emailAddress, content, randomToken)
	// generate the content of the email
	emailContent, err := emailTemplate.Content()
	if err != nil {
		logger.With("err", err).Error("can not generate template")
		return
	}
	// Send email
	logger.Debug("Send registration email")
	err = service.SendEmail(
		service.GetSenderEmailAddress(),
		mail.Address{
			Name:    emailAddress,
			Address: emailAddress,
		},
		subject,
		emailContent)
	if err != nil {
		logger.With("err", err).Error("can not send email")
		return
	}
	// generate default api token
	_, _ = CreateApiToken(service, identity.UID, "Default API Token", time.Now().Add(time.Hour*24*7*30*12*4))
	return
}
