package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	ass := assert.New(t)

	// setup
	EmptyIdentityTable()
	EmptyRegistrationConfirmationTable()

	// test

	// login with no email and password
	token, err := Login(testService, "", "", "", "")
	ass.Error(err)
	ass.Empty(token)

	// login user with no password
	token, err = Login(testService, testUserEmail, "", "", "")
	ass.Error(err)
	ass.Empty(token)

	// register user
	err = registerTestUser()
	ass.NoError(err)

	// check if user is confirmed
	dbUser := identity_models.Identity{}
	err = DbConnection.First(&dbUser).Error

	dbUser.Cleared = true
	dbUser.Blocked = false
	dbUser.Active = true
	err = DbConnection.Save(&dbUser).Error
	ass.NoError(err)

	// login user with now password
	token, err = Login(testService, testUserEmail, "wrongPassword", "hello", "123.123.123.134")
	ass.Error(err)
	ass.Empty(token)

	// login user with now password
	token, err = Login(testService, testUserEmail, testPw, "hello", "123.123.123.134")
	ass.NoError(err)
	ass.NotEmpty(token)

	// teardown
	EmptyIdentityTable()
	EmptyRegistrationConfirmationTable()
}
