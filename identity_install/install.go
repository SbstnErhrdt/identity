package identity_install

import (
	"github.com/SbstnErhrdt/go-gorm-all-sql/pkg/sql"
	"github.com/SbstnErhrdt/identity/models"
)

// Install will create the database schema
func Install() {
	client, err := sql.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	// databases
	err = client.Migrator().AutoMigrate(
		models.Identity{},
		models.IdentityLogin{},
		models.IdentityTokenMeta{},
		models.IdentityEmailChange{},
		models.IdentityRegistrationConfirmation{},
		models.IdentityResetPassword{},
	)
	if err != nil {
		panic(err)
	}
}
