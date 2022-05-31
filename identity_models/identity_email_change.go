package identity_models

import (
	"github.com/google/uuid"
	"time"
)

type IdentityEmailChange struct {
	Base
	// Relations
	IdentityUID uuid.UUID
	// Attributes
	ExpiredAt             time.Time
	ConfirmedOldAt        *time.Time
	ConfirmedNewAt        *time.Time
	ConfirmationOldClient string
	ConfirmationOldIP     string
	ConfirmationNewClient string
	ConfirmationNewIP     string
}

func (obj *IdentityEmailChange) TableName() string {
	return "identity_email_changes"
}
