package identity_controllers

import (
	"github.com/SbstnErhrdt/identity/identity_models"
	"github.com/google/uuid"
	"log/slog"
	"testing"
)

func createTestIdentity(t *testing.T) {

	testIdentity := &identity_models.Identity{
		Salutation:                 "",
		FirstName:                  "Tristan",
		LastName:                   "Tester",
		Email:                      "test@example.com",
		BackupEmail:                "",
		Phone:                      "",
		BackupPhone:                "",
		AcceptConditionsAndPrivacy: true,
		Active:                     true,
		Cleared:                    true,
		Blocked:                    false,
	}
	err := testIdentity.SetNewPassword(testService.GetPepper(), "securePassword1123%!")
	if err != nil {
		t.Error("can not set pw for test user", err)
	}
	err = TestDbConnection.Create(testIdentity).Error
	if err != nil {
		t.Error("can not create test user", err)
	}

}

func EmptyRelationsTable() {
	err := EmptyTable(&identity_models.IdentityRelation{})
	if err != nil {
		slog.With("err", err).Error("failed to empty identity table")
	}
}

func relationTestSetUp(t *testing.T) (identity *identity_models.Identity, err error) {
	EmptyRelationsTable()
	EmptyIdentityTable()
	createTestIdentity(t)
	identity = &identity_models.Identity{}
	err = TestDbConnection.First(identity).Error
	if err != nil {
		t.Error("can not get first identity", err)
		return
	}
	return
}

func relationTestTearDown() {
	EmptyIdentityTable()
	EmptyRelationsTable()
}

type testEntity struct {
	EntityType string
	EntityUID  uuid.UUID
}

func (e *testEntity) GetEntityType() string {
	return e.EntityType
}

func (e *testEntity) GetEntityUID() uuid.UUID {
	return e.EntityUID
}

func addTestRelation(t *testing.T, identity *identity_models.Identity, entityType, relationType string) {
	// add relation
	testEnt := &testEntity{
		EntityType: entityType,
		EntityUID:  uuid.New(),
	}

	err := AddIdentityRelation(testService, identity.UID, relationType, testEnt)
	if err != nil {
		t.Error(err)
		return
	}

	// check if relation exists
	dbRelation := &identity_models.IdentityRelation{}
	err = TestDbConnection.Order("created_at desc").First(dbRelation).Error
	if err != nil {
		t.Error(err)
		return
	}

	if dbRelation.IdentityUID.String() != identity.UID.String() {
		t.Error("wrong identity uid")
		return
	}

	if dbRelation.RelationType != relationType {
		t.Error("wrong relation type")
		return
	}

	if dbRelation.EntityType != entityType {
		t.Error("wrong entity type")
		return
	}

	if dbRelation.EntityUID.String() != testEnt.EntityUID.String() {
		t.Error("wrong entity uid")
		return
	}

}

func TestAddIdentityRelation(t *testing.T) {
	t.Log("add relation")
	identity, err := relationTestSetUp(t)
	if err != nil {
		t.Error(err)
		return
	}
	addTestRelation(t, identity, "test", "test_relation_type")
	relationTestTearDown()
}
