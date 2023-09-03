package identity_models

import "github.com/google/uuid"

type IdentityRelation struct {
	// Metadata
	Base
	// Attributes
	IdentityUID  uuid.UUID `gorm:"type:uuid; primary_key" json:"identityUID"`
	RelationType string    `gorm:"primary_key" json:"relationType"`
	EntityType   string    `gorm:"primary_key" json:"entityType"`
	EntityUID    uuid.UUID `gorm:"type:uuid; primary_key" json:"entityUID"`
}

func (obj *IdentityRelation) TableName() string {
	return "identity_relations"
}

type IdentityRelationEntity interface {
	GetEntityType() string
	GetEntityUID() uuid.UUID
}
