package identity_controllers

import "testing"

func TestReadIdentityRelation(t *testing.T) {
	identity, err := relationTestSetUp(t)
	if err != nil {
		t.Error(err)
	}
	addTestRelation(t, identity, "test", "test_relation")

	dbRelation, err := ReadIdentityRelation(testService, identity.UID, "test_relation", &testEntity{
		EntityType: "test",
	})
	if err != nil {
		t.Error(err)
	}

	if dbRelation.IdentityUID.String() != identity.UID.String() {
		t.Error("wrong identity uid")
	}

	if dbRelation.RelationType != "test_relation" {
		t.Error("wrong relation type")
	}

	if dbRelation.EntityType != "test" {
		t.Error("wrong entity type")
	}

	relationTestTearDown()
}

func TestReadIdentityRelations(t *testing.T) {
	identity, err := relationTestSetUp(t)
	if err != nil {
		t.Error(err)
	}
	addTestRelation(t, identity, "test", "test_relations")
	addTestRelation(t, identity, "test", "test_relations")

	dbRelations, err := ReadIdentityRelations(testService, identity.UID, "test_relations", &testEntity{
		EntityType: "test",
	})
	if err != nil {
		t.Error(err)
	}

	for _, dbRelation := range dbRelations {
		if dbRelation.IdentityUID.String() != identity.UID.String() {
			t.Error("wrong identity uid")
		}

		if dbRelation.RelationType != "test_relations" {
			t.Error("wrong relation type")
		}

		if dbRelation.EntityType != "test" {
			t.Error("wrong entity type")
		}
	}

	relationTestTearDown()
}
