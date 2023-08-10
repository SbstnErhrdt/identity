package identity_models

import (
	"github.com/google/uuid"
	"time"
)

type IdentityApiToken struct {
	// Metadata
	Base
	// Attributes
	IdentityUID    uuid.UUID `json:"identityUID" gorm:"type:varchar(36); primaryKey;"`
	TokenUID       uuid.UUID `json:"tokenUID" gorm:"type:varchar(36); primaryKey;"`
	Name           string    `json:"name"`
	Token          string    `json:"token"`
	ExpirationDate time.Time `json:"expirationDate"`
}

func (obj *IdentityApiToken) TableName() string {
	return "identity_api_tokens"
}
