package identity_models

import (
	"github.com/google/uuid"
	"time"
)

type IdentityResetPassword struct {
	// Metadata
	Base
	// Relations
	IdentityUID uuid.UUID `gorm:"type:uuid;index" json:"identityUID"`
	// Attributes
	Email                 string     `json:"email"`
	UserAgent             string     `json:"userAgent"`
	IP                    string     `json:"ip"`
	Origin                string     `json:"origin"`
	Expire                time.Time  `json:"expire"`
	Token                 string     `json:"token" gorm:"unique;index"`
	ConfirmationTime      *time.Time `json:"confirmationTime"`
	ConfirmationUserAgent string     `json:"confirmationUserAgent"`
	ConfirmationIP        string     `json:"confirmationIp"`
	ConfirmationOrigin    string     `json:"confirmationOrigin"`
}

func (obj *IdentityResetPassword) TableName() string {
	return "identity_password_resets"
}
