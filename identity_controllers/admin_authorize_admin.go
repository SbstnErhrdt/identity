package identity_controllers

import (
	"errors"
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ErrExternalNoAdmin is returned if the user is not an admin
var ErrExternalNoAdmin = errors.New("not an admin")

// ErrExternalIsAdmin is returned if the user is an admin
var ErrExternalIsAdmin = errors.New("error while checking if user is admin. Please try again later")

// IsAdmin checks if the user is an admin
func IsAdmin(service IdentityService, identityUID uuid.UUID) (err error) {
	logger := service.GetLogger().WithFields(logrus.Fields{
		"package":     "identity_admin_controllers",
		"func":        "IsAdmin",
		"identityUID": identityUID,
	})
	// build query
	err = service.GetSQLClient().
		Where("deleted_at IS NULL").
		Where("identity_uid = ?", identityUID).
		First(&identity_models.IdentityAdmin{}).Error
	// check if user is admin
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ErrExternalNoAdmin
		logger.Info("user no admin")
		return
	} else if err != nil {
		logger.WithError(err).Error("could not check if user is admin")
		err = ErrExternalIsAdmin
		return
	}
	logger.Info("user is admin")
	return
}
