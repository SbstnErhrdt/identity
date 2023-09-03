package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
)

// ReadIdentityRelations reads relations between an identity and an entity
func ReadIdentityRelations(service IdentityService, identityUID uuid.UUID, relationType string, entity identity_models.IdentityRelationEntity) (results []*identity_models.IdentityRelation, err error) {
	query := service.GetSQLClient().Where("deleted_at is NULL").Where("identity_uid = ?", identityUID)

	if relationType != "" {
		query = query.Where("relation_type = ?", relationType)
	}

	if entity != nil {
		if entity.GetEntityType() != "" {
			query = query.Where("entity_type = ?", entity.GetEntityType())
		}
		if entity.GetEntityUID() != uuid.Nil {
			query = query.Where("entity_uid = ?", entity.GetEntityUID())
		}
	}

	err = query.Find(&results).Error
	if err != nil {
		service.GetLogger().With("error", err).Error("Failed to read relations")
		return nil, err
	}
	return

}

// ReadIdentityRelation reads a relation between an identity and an entity
func ReadIdentityRelation(service IdentityService, identityUID uuid.UUID, relationType string, entity identity_models.IdentityRelationEntity) (result *identity_models.IdentityRelation, err error) {
	query := service.GetSQLClient().Where("deleted_at is NULL").Where("identity_uid = ?", identityUID)

	if relationType != "" {
		query = query.Where("relation_type = ?", relationType)
	}

	if entity != nil {
		if entity.GetEntityType() != "" {
			query = query.Where("entity_type = ?", entity.GetEntityType())
		}
		if entity.GetEntityUID() != uuid.Nil {
			query = query.Where("entity_uid = ?", entity.GetEntityUID())
		}
	}

	err = query.First(&result).Error
	if err != nil {
		service.GetLogger().With("error", err).Error("Failed to read relations")
		return nil, err
	}
	return
}
