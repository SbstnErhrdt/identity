package identity_controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfEmailIsFree(t *testing.T) {
	// setup
	EmptyIdentityTable()
	// test
	ass := assert.New(t)
	result, err := CheckIfEmailIsFree(testService, testUserEmail)
	ass.NoError(err, "there is an error")
	ass.True(result, "email should be free")

	// register user
	err = Register(
		testService,
		testUserEmail,
		testPw,
		true,
		"test-agent",
		"0.0.0.0",
		"dev.local",
	)
	ass.NoError(err)

	// check again
	result, err = CheckIfEmailIsFree(testService, testUserEmail)
	ass.NoError(err, "there is an error")
	ass.False(result, "email should be not free anymore")

	// tear down
	EmptyIdentityTable()
}
