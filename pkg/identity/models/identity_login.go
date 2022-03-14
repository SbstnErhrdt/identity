package models

import (
	"github.com/google/uuid"
)

type IdentityLogin struct {
	Base
	// Relations
	IdentityUID uuid.UUID
	// Attributes
	Email     string
	UserAgent string
	IP        string
}

func (obj *IdentityLogin) TableName() string {
	return "identity_logins"
}
