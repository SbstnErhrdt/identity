package communication

import "github.com/SbstnErhrdt/identity/models"

// Communication is the interface for communication (email/sms)
type Communication interface {
	Registration(identity models.Identity, urlWithToken string) error
	RegistrationConfirmation(identity models.Identity) error
	PasswordReset(identity models.Identity, urlWithPasswordResetToken string) error
	PasswordResetConfirmation(identity models.Identity, urlWithPasswordResetToken string) error
	IdentityChange(identity models.Identity, newIdentity string, urlWithConfirmationToken string) error
	IdentityChangeConfirmation(identity models.Identity) error
	IdentityChangeFinalConfirmation(identity models.Identity) error
}
