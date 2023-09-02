package identity_models

import "github.com/google/uuid"

type IdentityTokenMeta struct {
	// Metadata
	Base
	// Attributes
	TokenUID  uuid.UUID `gorm:"type:uuid;index" json:"tokenUID"`
	TokenType TokenType `json:"tokenType"`
}

func (obj *IdentityTokenMeta) TableName() string {
	return "identity_tokens_meta"
}
