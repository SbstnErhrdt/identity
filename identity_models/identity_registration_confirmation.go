package identity_models

import (
	"github.com/google/uuid"
	"time"
)

type IdentityRegistrationConfirmation struct {
	Base
	// Relations
	IdentityUID uuid.UUID `gorm:"type:varchar(36);index" json:"identityUID"`
	// Attributes
	Token                 string     `json:"token" gorm:"unique;index"`
	RegistrationUserAgent string     `json:"registrationUserAgent"`
	RegistrationIP        string     `json:"registrationIP"`
	ConfirmationUserAgent string     `json:"confirmationUserAgent"`
	ConfirmationIP        string     `json:"confirmationIP"`
	ExpiredAt             time.Time  `json:"expiredAt"`
	ConfirmedAt           *time.Time `json:"confirmedAt"`
}

func (obj *IdentityRegistrationConfirmation) TableName() string {
	return "identity_registration_confirmations"
}
