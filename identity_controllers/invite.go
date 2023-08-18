package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"net/mail"
)

// InviteUser invites a user to the system
func InviteUser(service IdentityService, origin, subject, firstName, lastName, emailAddress, content, link string) (err error) {
	// sanitize email address
	emailAddress = SanitizeEmail(emailAddress)
	// init logger
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
	emailTemplate := service.ResolveInvitationEmailTemplate(origin, firstName, lastName, emailAddress, content, link)
	// generate the content of the email
	emailContent, err := emailTemplate.Content()
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
		emailContent)
	if err != nil {
		logger.With("err", err).Error("could not send email")
		return
	}
	return
}

// InvitationConfirmation confirms an invitation
// - check if accept terms and conditions is true
// - reset password
// - confirm registration
// - accept terms and conditions
func InvitationConfirmation(service IdentityService, token, newPassword, newPasswordConfirmation, userAgent, ip, origin string, acceptTermsAndConditions bool) (userToken string, err error) {
	logger := service.GetLogger().With(
		"method", "InvitationConfirmation",
		"token", token,
		"origin", origin,
	)

	// check if accept terms and conditions is true
	if !acceptTermsAndConditions {
		logger.With("err", err).Error("accept terms and conditions is false")
		err = ErrAcceptTermsAndConditions
		return
	}

	// reset password
	err = ResetPassword(service, token, newPassword, newPasswordConfirmation, userAgent, ip, origin)
	if err != nil {
		logger.With("err", err).Error("could not reset password")
		return
	}

	// confirm registration
	err = RegistrationConfirmation(service, token, userAgent, ip)
	if err != nil {
		logger.With("err", err).Error("could not confirm registration")
		return
	}

	// accept terms and conditions
	err = AcceptTerms(service, token)
	if err != nil {
		logger.With("err", err).Error("could not accept terms and conditions")
		return
	}

	// check if token is in table
	tokenIdentityResult := identity_models.IdentityRegistrationConfirmation{}
	logger.Debug("find registration confirmation")
	err = service.
		GetSQLClient().
		Where("token = ?", token).
		First(&tokenIdentityResult).Error
	if err != nil {
		logger.With("err", err).Error("could not find registration confirmation")
		err = ErrTokenNotFound
		return
	}

	// get the identity
	dbIdentity, err := GetIdentityByUID(service, tokenIdentityResult.IdentityUID)
	if err != nil {
		logger.With("err", err).Error("could not get identity")
		return
	}

	userToken, err = Login(service, dbIdentity.Email, newPassword, userAgent, ip)
	if err != nil {
		logger.With("err", err).Error("could not login")
		return
	}

	return
}

// AcceptTerms accepts the terms and conditions
// uses a token to find the identity
// and sets the accept terms and conditions to true
func AcceptTerms(service IdentityService, token string) (err error) {
	logger := service.GetLogger().With(
		"method", "AcceptTerms",
		"token", token,
	)
	// get the token from the database
	resetPassword := identity_models.IdentityResetPassword{}
	err = service.GetSQLClient().Where("token = ?", token).First(&resetPassword).Error
	if err != nil {
		logger.With("err", err).Error("cam not get reset password from database")
		err = ErrTokenNotFound
		return err
	}

	// get the identity from the database
	identity := identity_models.Identity{}
	err = service.GetSQLClient().Where("uid = ?", resetPassword.IdentityUID).First(&identity).Error
	if err != nil {
		logger.With("err", err).Error("can not get identity from database")
		return
	}

	// update the identity
	identity.AcceptConditionsAndPrivacy = true
	err = service.GetSQLClient().Save(&identity).Error
	if err != nil {
		logger.With("err", err).Error("can not update identity")
		return
	}
	return
}
