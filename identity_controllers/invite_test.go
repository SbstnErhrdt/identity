package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInviteUser(t *testing.T) {
	ass := assert.New(t)
	err := InviteUser(testService, "erhardt.net", "Invitation", "Sebastian", "Erhardt", "test@erhardt.net", "Please join my friend", "https://erhardt.net")
	ass.NoError(err)
}

func TestInvitationConfirmation(t *testing.T) {
	ass := assert.New(t)
	// empty identities
	err := EmptyTable(&identity_models.Identity{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	// empty registration confirmation
	err = EmptyTable(&identity_models.IdentityRegistrationConfirmation{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	// empty reset pw tokens
	err = EmptyTable(&identity_models.IdentityResetPassword{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("create and invite user")
	dbIdentity, err := AdminCreateIdentityAndInvite(testService, "test.local", "Test Invitation to test platform", "Hi Tristan. Please join. And confirm", "Tristan", "Tester", "test@erhardt.net")
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	dbRegToken := &identity_models.IdentityRegistrationConfirmation{}
	err = testService.GetSQLClient().Find(&dbRegToken).Error
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	ass.NotEqual(dbIdentity.UID.String(), uuid.Nil.String())
	ass.Equal(dbIdentity.UID.String(), dbRegToken.IdentityUID.String())

	dbResetToken := &identity_models.IdentityResetPassword{}
	err = testService.GetSQLClient().Find(&dbResetToken).Error
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	ass.Equal(dbIdentity.UID.String(), dbResetToken.IdentityUID.String())
	ass.Equal(dbRegToken.Token, dbResetToken.Token)

	agent := "confirm invite test agent"

	t.Log("confirm invitation")
	err = InvitationConfirmation(testService, dbRegToken.Token, testPw, testPw, agent, "127.0.0.1", "test.local", true)
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	updatedDbIdentity := identity_models.Identity{}

	err = testService.GetSQLClient().Where("email = ?", "test@erhardt.net").Find(&updatedDbIdentity).Error
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	ass.True(updatedDbIdentity.AcceptConditionsAndPrivacy)
	t.Log("identity has accepted terms and conditions")

	updatedDbRegToken := &identity_models.IdentityRegistrationConfirmation{}
	err = testService.GetSQLClient().Find(&updatedDbRegToken).Error
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	ass.NotNil(updatedDbRegToken.UpdatedAt)
	ass.Equal(agent, updatedDbRegToken.ConfirmationUserAgent)

	updatedDbResetToken := &identity_models.IdentityResetPassword{}
	err = testService.GetSQLClient().Find(&updatedDbResetToken).Error
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	ass.NotNil(updatedDbResetToken.UpdatedAt)
	ass.Equal(agent, updatedDbResetToken.ConfirmationUserAgent)
	t.Log("confirmed invitation user agent")

	err = EmptyTable(&identity_models.Identity{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	// empty registration confirmation
	err = EmptyTable(&identity_models.IdentityRegistrationConfirmation{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

	// empty reset pw tokens
	err = EmptyTable(&identity_models.IdentityResetPassword{})
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}

}
