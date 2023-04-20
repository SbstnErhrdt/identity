package identity

import (
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_controllers"
	"github.com/SbstnErhrdt/identity/identity_install"
	"log"
	"os"
	"testing"
)

func init() {
	env.LoadEnvFiles("./test/.env")
}

func TestMain(m *testing.M) {
	log.Println("Start identity package unit test")
	code := m.Run()
	log.Println("Tests finished")
	os.Exit(code)
}

// TestInstall will test the installation of the identity package
func TestInstall(t *testing.T) {

	identity_controllers_test.Connec

	err := identity_install.Install(DbConnection)
	if err != nil {
		t.Fatal(err)
	}
}

func TestServiceSetUp(t *testing.T) {
	controllerService := identity_controllers.NewService("test")

}

func TestUserRegistration(t *testing.T) {

}
