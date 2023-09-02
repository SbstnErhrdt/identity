package identity_models

import (
	"github.com/google/uuid"
)

type IdentityAdmin struct {
	// Metadata
	Base
	// Attributes
	IdentityUID uuid.UUID `gorm:"type:uuid; primary_key" json:"identityUID"`
}

func (obj *IdentityAdmin) TableName() string {
	return "identity_admins"
}
