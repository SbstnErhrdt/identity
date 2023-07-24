package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDeleteApiToken(t *testing.T) {
	ass := assert.New(t)
	err := EmptyTable(&identity_models.IdentityApiToken{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	testIdentity := uuid.New()
	token, err := CreateApiToken(testService, testIdentity, "test token", time.Now().Add(time.Hour*20).UTC())
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	err = DeleteApiToken(testService, testIdentity, token.UID)
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	err = EmptyTable(&identity_models.IdentityApiToken{})
	ass.NoError(err)

}
