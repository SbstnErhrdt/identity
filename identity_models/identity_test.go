package identity_models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_VerifyPassword(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	testPepper := "THIS_IS_A_TEST_PEPPER_12345"
	pw := "Hello World"

	user := Identity{}
	err := user.SetNewPassword(testPepper, pw)
	ass.NoError(err)
	if err != nil {
		t.Error(err)
	}

	ass.True(user.CheckPassword(testPepper, pw))
	ass.True(user.CheckPassword(testPepper, string([]byte(pw))))
	ass.False(user.CheckPassword(testPepper, "stupidWrongPassword"))

	err = user.SetNewPassword(testPepper, "pw2")
	ass.NoError(err)
	if err != nil {
		t.Error(err)
	}

	ass.False(user.CheckPassword(testPepper, pw))
	ass.False(user.CheckPassword(testPepper, string([]byte(pw))))
	ass.False(user.CheckPassword(testPepper, "stupidWrongPassword"))
	ass.True(user.CheckPassword(testPepper, "pw2"))

	otherTestPepper := "THIS_IS_A_DIFFERENT_TEST_PEPPER_12345"
	ass.False(user.CheckPassword(otherTestPepper, "pw2"))
}
