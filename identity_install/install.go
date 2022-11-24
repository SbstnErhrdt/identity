package identity_install

import (
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/sql"
	"github.com/SbstnErhrdt/identity/identity_models"
)

// Install will create the database schema
func Install() {
	client, err := sql.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	// databases
	err = client.Migrator().AutoMigrate(
		identity_models.Identity{},
		identity_models.IdentityLogin{},
		identity_models.IdentityTokenMeta{},
		identity_models.IdentityEmailChange{},
		identity_models.IdentityRegistrationConfirmation{},
		identity_models.IdentityResetPassword{},
	)
	if err != nil {
		panic(err)
	}
}
