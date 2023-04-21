package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// ReadAllUsers reads all users
func ReadAllUsers(service IdentityService, adminUID uuid.UUID, keyword string, offset, limit int, orderBy string) (results []*identity_models.Identity, amount int64, err error) {
	// check if admin
	err = IsAdmin(service, adminUID)
	if err != nil {
		return
	}
	// build query
	query := service.GetSQLClient().
		Where("deleted_at IS NULL")

	if len(keyword) > 0 {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}
	// Order by
	if len(orderBy) > 0 {
		query = query.Order(orderBy)
	}
	// Extract amount
	err = query.
		// get multiple results
		// Limit / Offset
		Limit(limit).
		Offset(offset).
		// execute the query
		Find(&results).
		// get the total count
		Limit(-1).
		Offset(0).
		// execute the query
		Count(&amount).Error
	if err != nil {
		log.WithError(err).Error("could not read all users")
	}
	return
}
