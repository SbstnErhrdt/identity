package identity_models

import (
	"github.com/google/uuid"
	"time"
)

type IdentityEmailChange struct {
	Base
	// Relations
	IdentityUID uuid.UUID `gorm:"type:varchar(36);index" json:"identityUID"`
	// Attributes
	ExpiredAt             time.Time  `gorm:"index" json:"expiredAt"`
	ConfirmedOldAt        *time.Time `json:"confirmedOldAt"`
	ConfirmedNewAt        *time.Time `json:"confirmedNewAt"`
	ConfirmationOldClient string     `json:"confirmationOldClient"`
	ConfirmationOldIP     string     `json:"confirmationOldIP"`
	ConfirmationNewClient string     `json:"confirmationNewClient"`
	ConfirmationNewIP     string     `json:"confirmationNewIP"`
}

func (obj *IdentityEmailChange) TableName() string {
	return "identity_email_changes"
}
