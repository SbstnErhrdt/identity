package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdminCreateIdentityAndInvite(t *testing.T) {
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

	dbIdentity, err := AdminCreateIdentityAndInvite(testService, "test.local", "Invitation to test platform", "Hi Tristan. Please join.", "Tristan", "Tester", "test@erhardt.net")
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	ass.False(dbIdentity.AcceptConditionsAndPrivacy)

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

}
