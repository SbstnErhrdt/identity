package identity_install

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Install will create the database schema
func Install(db *gorm.DB) (err error) {
	log.Info("creating identity database schema")
	// databases
	err = db.Migrator().AutoMigrate(
		identity_models.Identity{},
		identity_models.IdentityLogin{},
		identity_models.IdentityTokenMeta{},
		identity_models.IdentityEmailChange{},
		identity_models.IdentityRegistrationConfirmation{},
		identity_models.IdentityResetPassword{},
	)
	if err != nil {
		log.WithError(err).Error("failed to create / migrate database")
	}
	return
}
