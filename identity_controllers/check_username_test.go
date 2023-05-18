package identity_controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfUsernameIsValid(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	err := CheckIfUsernameIsValid("uGQ5j")
	ass.NoError(err)

	err = CheckIfUsernameIsValid("helloWorld")
	ass.NoError(err)

	err = CheckIfUsernameIsValid("a")
	ass.Error(err)

	err = CheckIfUsernameIsValid("awfafaww...")
	ass.Error(err)

	err = CheckIfUsernameIsValid("awfa_!§faww")
	ass.Error(err)
}
