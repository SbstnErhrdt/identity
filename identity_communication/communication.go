package identity_communication

import "github.com/SbstnErhrdt/identity/identity_models"

// Communication is the interface for communication (email/sms)
type Communication interface {
	Registration(identity identity_models.Identity, urlWithToken string) error
	RegistrationConfirmation(identity identity_models.Identity) error
	PasswordReset(identity identity_models.Identity, urlWithPasswordResetToken string) error
	PasswordResetConfirmation(identity identity_models.Identity, urlWithPasswordResetToken string) error
	IdentityChange(identity identity_models.Identity, newIdentity string, urlWithConfirmationToken string) error
	IdentityChangeConfirmation(identity identity_models.Identity) error
	IdentityChangeFinalConfirmation(identity identity_models.Identity) error
}
