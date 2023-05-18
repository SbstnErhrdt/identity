package identity_install

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	log "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
)

// Install will create the database schema
func Install(db *gorm.DB) (err error) {
	log.Info("creating identity database schema")
	// databases
	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		log.WithError(err).Error("failed to create extension")
		return err
	}

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
		return err
	}
	return
}
