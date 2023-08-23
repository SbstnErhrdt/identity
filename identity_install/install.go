package identity_install

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"gorm.io/gorm"
	"log/slog"
)

// Install will create the database schema
func Install(db *gorm.DB) (err error) {
	slog.Info("creating identity database schema")
	// databases
	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		slog.With("err", err).Error("failed to create extension")
		return err
	}

	err = db.Migrator().AutoMigrate(
		identity_models.Identity{},
		identity_models.IdentityLogin{},
		identity_models.IdentityTokenMeta{},
		identity_models.IdentityEmailChange{},
		identity_models.IdentityRegistrationConfirmation{},
		identity_models.IdentityResetPassword{},
		identity_models.IdentityApiToken{},
		identity_models.IdentityAdmin{},
	)
	if err != nil {
		slog.With("err", err).Error("failed to create / migrate database")
		return err
	}
	return
}
