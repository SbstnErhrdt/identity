package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
)

// AddIdentityRelation adds a relation between an identity and an entity
func AddIdentityRelation(
	service IdentityService,
	identityUID uuid.UUID,
	relationType string,
	entity identity_models.IdentityRelationEntity,
) (err error) {
	err = service.GetSQLClient().Create(&identity_models.IdentityRelation{
		IdentityUID:  identityUID,
		RelationType: relationType,
		EntityType:   entity.GetEntityType(),
		EntityUID:    entity.GetEntityUID(),
	}).Error
	if err != nil {
		service.GetLogger().With("error", err).Error("Failed to create relation")
		return err
	}
	return
}
