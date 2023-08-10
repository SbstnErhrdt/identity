package identity_models

import (
	"github.com/google/uuid"
)

type IdentityLogin struct {
	Base
	// Relations
	IdentityUID *uuid.UUID `gorm:"type:varchar(36);index" json:"identityUID"`
	// Attributes
	Email     string `json:"email"`
	UserAgent string `json:"userAgent"`
	IP        string `json:"ip"`
	Origin    string `json:"origin"`
}

func (obj *IdentityLogin) TableName() string {
	return "identity_logins"
}
