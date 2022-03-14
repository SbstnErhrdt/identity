package models

import (
	"github.com/google/uuid"
	"time"
)

type IdentityResetPassword struct {
	// Metadata
	Base
	// Relations
	IdentityUID uuid.UUID `json:"identityUID"`
	// Attributes
	Expire time.Time `json:"expire"`
	Token  string    `json:"token" gorm:"unique;index"`
}

func (obj *IdentityResetPassword) TableName() string {
	return "identity_password_resets"
}
