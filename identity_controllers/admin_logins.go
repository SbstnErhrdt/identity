package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
)

// ReadIdentityLogins reads the identity logins
func ReadIdentityLogins(service IdentityService, keyword string, offset, limit int, orderBy string) (results []*identity_models.IdentityLogin, amount int64, err error) {
	// build query
	query := service.GetSQLClient().
		Where("deleted_at IS NULL")

	if len(keyword) > 0 {
		query = query.Where("email LIKE ?", "%"+keyword+"%")
	}
	// Order by
	if len(orderBy) > 0 {
		query = query.Order(orderBy)
	} else {
		query = query.Order("created_at DESC")
	}

	results = []*identity_models.IdentityLogin{}

	// Extract amount
	err = query.
		Where("deleted_at IS NULL").
		// get multiple results
		// Limit / Offset
		Limit(limit).
		Offset(offset).
		// execute the query
		Find(&results).
		// get the total count
		Limit(-1).
		Offset(-1).
		// execute the query
		Count(&amount).Error
	if err != nil {
		service.GetLogger().With("err", err).Error("could not read all login attempts")
	}
	return
}
