package identity_controllers

import "testing"

func TestRemoveIdentityRelation(t *testing.T) {

	identity, err := relationTestSetUp(t)
	if err != nil {
		t.Error(err)
	}
	addTestRelation(t, identity, "test", "test_relations")
	addTestRelation(t, identity, "test", "test_relations")
	addTestRelation(t, identity, "testDel", "test_relations")

	// get

	dbDelRelation, err := ReadIdentityRelation(testService, identity.UID, "test_relations", &testEntity{
		EntityType: "testDel",
	})
	if err != nil {
		t.Error(err)
	}

	err = RemoveIdentityRelation(testService, identity.UID, "test_relations", &testEntity{
		EntityType: "testDel",
		EntityUID:  dbDelRelation.EntityUID,
	})
	if err != nil {
		t.Error(err)
	}

	dbRelations, err := ReadIdentityRelations(testService, identity.UID, "test_relations", &testEntity{
		EntityType: "test",
	})
	if err != nil {
		t.Error(err)
	}

	if len(dbRelations) != 2 {
		t.Error("wrong number of relations", len(dbRelations))
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

	// check if deleted
	dbRelations, err = ReadIdentityRelations(testService, identity.UID, "test_relations", nil)
	if err != nil {
		t.Error(err)
	}

	if len(dbRelations) != 2 {
		t.Error("wrong number of relations", len(dbRelations))
		for _, dbRelation := range dbRelations {
			t.Log(*dbRelation)
		}
	}

	relationTestTearDown()

}
