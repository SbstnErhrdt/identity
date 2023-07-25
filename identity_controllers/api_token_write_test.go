package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateApiToken(t *testing.T) {
	ass := assert.New(t)
	err := EmptyTable(&identity_models.IdentityApiToken{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	testIdentity := uuid.New()
	tokenDateTime := time.Now().Add(time.Hour * 20).UTC()
	token, err := CreateApiToken(testService, testIdentity, "test token", tokenDateTime)
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	ass.NotEmpty(token)

	dbToken := &identity_models.IdentityApiToken{}
	err = testService.GetSQLClient().Find(&dbToken).Error
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	ass.Equal("test token", dbToken.Name)
	ass.Equal(tokenDateTime, dbToken.ExpirationDate.UTC())

}

func TestDeleteApiToken(t *testing.T) {

	ass := assert.New(t)
	err := EmptyTable(&identity_models.IdentityApiToken{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	testIdentity := uuid.New()
	tokenDateTime := time.Now().Add(time.Hour * 20).UTC()
	token, err := CreateApiToken(testService, testIdentity, "test token", tokenDateTime)
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

	dbDeletedToken := &identity_models.IdentityApiToken{}
	err = testService.GetSQLClient().Unscoped().Find(&dbDeletedToken).Error
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	ass.NotNil(dbDeletedToken.DeletedAt)

	err = EmptyTable(&identity_models.IdentityApiToken{})
	ass.NoError(err)

}
