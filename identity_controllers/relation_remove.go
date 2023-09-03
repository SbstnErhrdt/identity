package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
)

var ErrNilIdentityUID = errors.New("identity uid is nil")

var ErrNilEntity = errors.New("entity is nil")

var ErrNilRelationType = errors.New("relation type is nil")

// RemoveIdentityRelation removes a relation between an identity and an entity
func RemoveIdentityRelation(
	service IdentityService,
	identityUID uuid.UUID,
	relationType string,
	entity identity_models.IdentityRelationEntity,
) (err error) {
	if identityUID == uuid.Nil {
		return ErrNilIdentityUID
	}
	if entity == nil {
		return ErrNilEntity
	}
	if relationType == "" {
		return ErrNilRelationType
	}

	if entity.GetEntityType() == "" {
		return ErrNilEntity
	}

	if entity.GetEntityUID() == uuid.Nil {
		return ErrNilEntity
	}

	err = service.GetSQLClient().
		Where("deleted_at is NULL").
		Where("identity_uid = ?", identityUID).
		Where("relation_type = ?", relationType).
		Where("entity_type = ?", entity.GetEntityType()).
		Where("entity_uid = ?", entity.GetEntityUID()).
		Delete(&identity_models.IdentityRelation{}).Error
	if err != nil {
		service.GetLogger().With("error", err).Error("Failed to delete relation")
		return err
	}
	return
}
