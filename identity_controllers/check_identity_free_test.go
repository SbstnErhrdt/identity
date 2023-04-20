package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func deleteAllIdentities() {
	i := identity_models.Identity{}
	DbConnection.Exec("DELETE FROM ?", i.TableName())
}

func TestCheckIfEmailIsFree(t *testing.T) {

	ass := assert.New(t)
	result, err := CheckIfEmailIsFree(testService, testUserEmail)
	ass.NoError(err, "There is an error")
	ass.True(result, "Result is not true")
}
