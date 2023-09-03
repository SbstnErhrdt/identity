package identity_install

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"gorm.io/gorm"
	"log/slog"
)

// AutoMigrate will auto migrate the database schema
func AutoMigrate(db *gorm.DB) (err error) {
	slog.Info("auto migrate database schema")
	// databases
	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		slog.With("err", err).Error("failed to create extension")
		return err
	}
	err = db.Migrator().AutoMigrate(
		identity_models.IdentityAdmin{},
		identity_models.IdentityApiToken{},
		identity_models.IdentityEmailChange{},
		identity_models.Identity{},
		identity_models.IdentityLogin{},
		identity_models.IdentityResetPassword{},
		identity_models.IdentityRegistrationConfirmation{},
		identity_models.IdentityTokenMeta{},
		identity_models.IdentityRelation{},
	)
	if err != nil {
		slog.With("err", err).Error("failed to create / migrate database")
		return err
	}
	return
}
