package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testPw = "MaggieThatcherIs110%Sexy!"

func TestRegister(t *testing.T) {
	// setup
	EmptyIdentityTable()
	// test
	ass := assert.New(t)
	err := registerTestUser()
	ass.NoError(err)
	dbIdentity := identity_models.Identity{}
	err = DbConnection.First(&dbIdentity).Error
	ass.NoError(err)
	ass.Equal(testUserEmail, dbIdentity.Email)
	// teardown
	EmptyIdentityTable()
}

func registerTestUser() error {
	return Register(
		testService,
		testUserEmail,
		testPw,
		true,
		"test-agent",
		"0.0.0.0",
		"dev.local",
	)
}

func TestRegistrationConfirmation(t *testing.T) {
	ass := assert.New(t)

	// setup
	EmptyIdentityTable()
	EmptyRegistrationConfirmationTable()

	// test
	err := registerTestUser()
	ass.NoError(err)

	// get the registration confirmation
	initialConfDbEntry := identity_models.IdentityRegistrationConfirmation{}
	err = DbConnection.First(&initialConfDbEntry).Error
	ass.NoError(err)
	ass.Empty(initialConfDbEntry.ConfirmedAt)

	// confirm registration with wrong token
	err = RegistrationConfirmation(testService, "wrong token", "test-agent", "0.0.0.0")
	ass.Error(err)

	// confirm with right token
	err = RegistrationConfirmation(testService, initialConfDbEntry.Token, "test-agent", "0.0.0.0")
	ass.NoError(err)

	// get the registration confirmation again
	confDbEntry := identity_models.IdentityRegistrationConfirmation{}
	err = DbConnection.First(&confDbEntry).Error
	ass.NoError(err)
	ass.NotEmpty(confDbEntry.ConfirmedAt)

	// check if user is confirmed
	dbUser := identity_models.Identity{}
	err = DbConnection.First(&dbUser).Error
	ass.True(dbUser.Cleared)

	// teardown
	EmptyIdentityTable()
	EmptyRegistrationConfirmationTable()
}
