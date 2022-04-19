package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfEmailIsFree(t *testing.T) {

	/*
		assert := assert.New(t)
		env, err := environment.ReadEnvFromJson("../../env.test.json")
		assert.NoError(err)
		sha, err := shared.CreateSharedObject(env)
		assert.NoError(err)

		result, err := CheckIfEmailIsFree(&sha, "sebastwagesian.erhafesfgdt@cdtm.de")
		fmt.Println("Result", result)
		fmt.Println("err", err)
		assert.NoError(err, "There is an error")
		assert.True(result, "Result is not true")
	*/

}

func TestSetPasswordOfUserByName(t *testing.T) {
	/*
		assert := assert.New(t)
		env, err := environment.ReadEnvFromJson("../../env.test.json")
		assert.NoError(err)
		sha, err := shared.CreateSharedObject(env)
		assert.NoError(err)

		err = SetPasswordOfUserByName(&sha, "uGQ5j", "Db1!VnZLD")
		assert.NoError(err)
	*/

}

func TestCheckIfUsernameIsValid(t *testing.T) {
	ass := assert.New(t)

	err := CheckIfUsernameIsValid("uGQ5j")
	ass.NoError(err)

	err = CheckIfUsernameIsValid("helloWorld")
	ass.NoError(err)

	err = CheckIfUsernameIsValid("a")
	ass.Error(err)

	err = CheckIfUsernameIsValid("awfafaww...")
	ass.Error(err)

}
