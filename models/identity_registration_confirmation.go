package models

import (
	"github.com/google/uuid"
	"time"
)

type IdentityRegistrationConfirmation struct {
	Base
	// Relations
	IdentityUID uuid.UUID
	// Attributes
	Token                 string `json:"token" gorm:"unique;index"`
	RegistrationUserAgent string
	RegistrationIP        string
	ConfirmationUserAgent string
	ConfirmationIP        string
	ExpiredAt             time.Time
	ConfirmedAt           *time.Time
}

func (obj *IdentityRegistrationConfirmation) TableName() string {
	return "identity_registration_confirmations"
}
