package repositories

import (
	"fmt"
	"mallekoppie/ChaosGenerator/ChaosMaster/models"
	"testing"

	"github.com/google/uuid"
)

func TestFormatting(t *testing.T) {

	result := fmt.Sprintf("mongodb://%v:%v", "localhost", 27017)

	t.Log(result)

}

func createTestGroup() models.TestGroup {
	testCollections := make([]models.TestCollection, 0)
	testGroupId := uuid.New().String()
	tests := make([]models.Test, 0)
	tests = append(tests, models.Test{ID: "test"})
	colId := uuid.New().String()
	col1 := models.TestCollection{ID: colId,
		Name: "some name", GroupId: testGroupId, Tests: tests}
	testCollections = append(testCollections, col1)

	testGroup := models.TestGroup{

		ID:              testGroupId,
		Name:            "Unit Test Name",
		Description:     "a Nice Description",
		TestCollections: testCollections,
	}
	testGroup.ID = testGroupId

	return testGroup
}

func createTestCollection() models.TestCollection {
	testCollection := models.TestCollection{}
	tests := make([]models.Test, 0)
	tests = append(tests, models.Test{ID: "test", Url: "https://bla:9000"})

	testCollection.Tests = tests

	testCollection.ID = uuid.New().String()
	testCollection.GroupId = uuid.New().String()

	return testCollection
}

func TestInsertTestGroup(t *testing.T) {
	testGroup := createTestGroup()

	err := AddTestGroup(testGroup)
	if err != nil {
		t.Log("Error when inserting testGroup: ", err.Error())
		t.Fail()
	}
}

func TestGetTestGroup(t *testing.T) {
	descriptionToFindAgain := "not really that unique but should be good enough"

	testGroup := createTestGroup()
	testGroup.Description = descriptionToFindAgain

	err := AddTestGroup(testGroup)
	if err != nil {
		t.Log("Error when inserting testGroup: ", err.Error())
		t.Fail()
	}

	result, err := GetTestGroup(testGroup.ID)
	if err != nil {
		t.Fatal("Unable to find test group: ", err.Error())
		t.Fail()
	}

	if result.Description != descriptionToFindAgain {
		t.Fatal("Test group descriptions aren't the same")
		t.Fail()
	}
}

func TestUpdateTestGroup(t *testing.T) {
	testGroup := createTestGroup()
	testGroup.Description = "This is not updated"
	err := AddTestGroup(testGroup)
	if err != nil {
		t.Fatal("Unable to add Test Group: ", err.Error())
		t.FailNow()
	}

	testGroup.Description = "This has been updated"
	testGroup.TestCollections[0].Description = "Also updated for TestUpdateTestGroup"
	err = UpdateTestGroup(testGroup)
	if err != nil {
		t.Fatalf("Error updating record: %v", err.Error())
		t.FailNow()
	}
}

func TestGetAllTestGroups(t *testing.T) {
	testGroups, err := GetAllTestGroups()
	if err != nil {
		t.Fatal("Error retrieving all test groups: ", err.Error())
		t.Fail()
	}

	for i := range testGroups {
		t.Log("Test Group returned: ", testGroups[i])

		for c := range testGroups[i].TestCollections {
			t.Log("Test collection: ", testGroups[i].TestCollections[c])
		}
	}
}

func TestGetTestCollectionForTestGroup(t *testing.T) {
	group := createTestGroup()

	err := AddTestGroup(group)
	if err != nil {
		t.Fatal("Error adding test group: ", err.Error())
		t.FailNow()
	}

	result, err := GetTestCollectionsForGroup(group.ID)
	if err != nil {
		t.Fatal("Error getting test collections for group: ", err.Error())
		t.FailNow()
	}

	if len(result) < 1 {
		t.Log("Incorrect number of test collections returned: ", len(result))
		t.FailNow()
	}

	for i := range result {
		t.Log("Returned Test Collection: ", result[i])
	}

}

func TestAddTestCollection(t *testing.T) {
	collection1 := createTestCollection()
	collection1.Description = "First test collection that we are addiong for the group"
	collection2 := createTestCollection()
	collection2.Description = "Second test collection that we are addiong for the group"
	group := createTestGroup()
	group.Description = "Test Group with seperate TestCollections"
	collection1.GroupId = group.ID
	collection2.GroupId = group.ID

	err := AddTestGroup(group)
	if err != nil {
		t.Fatal("Unable to add test Group")
		t.FailNow()
	}

	err = AddTestCollection(collection1)
	if err != nil {
		t.Log("Error saving testcollection1 collection: ", err.Error())
		t.FailNow()
	}

	err = AddTestCollection(collection2)
	if err != nil {
		t.Log("Error saving testcollection2 collection: ", err.Error())
		t.FailNow()
	}

	resultTestGroup, err := GetTestGroup(group.ID)
	if err != nil {
		t.Fatal("Unable to retrieve TestGroup: ", err.Error())
		t.FailNow()
	}

	if resultTestGroup.ID != group.ID {
		t.Fatal("Wrong test group returned")
		t.FailNow()
	}

	for index := range resultTestGroup.TestCollections {
		if resultTestGroup.TestCollections[index].ID == collection1.ID ||
			resultTestGroup.TestCollections[index].ID == collection2.ID ||
			resultTestGroup.TestCollections[index].ID == group.TestCollections[0].ID {
			//Success continue
		} else {
			t.Fatal("Incorrect TestCollection returned with TestGroup")
			t.FailNow()
		}
	}
}

func TestDeleteTestGroupAndTestCollection(t *testing.T) {
	group := createTestGroup()
	group.Description = "This must be deleted during the test"

	err := AddTestGroup(group)
	if err != nil {
		t.Fatal("Unable to add Test group: ", err.Error())
		t.FailNow()
	}

	err = DeleteTestGroup(group.ID)
	if err != nil {
		t.Fatal("Unable to delete Test Group: ", err.Error())
		t.FailNow()
	}
}

func TestDeleteTestCollection(t *testing.T) {
	group := createTestGroup()
	group.TestCollections[0].Description = "This will be deleted"
	err := AddTestGroup(group)
	if err != nil {
		t.Fatal("Unable to add TestGroup: ", err.Error())
		t.FailNow()
	}

	err = DeleteTestCollection(group.TestCollections[0].ID)
	if err != nil {
		t.Fatal("Unable to delete test collection: ", err.Error())
		t.FailNow()
	}
}

func TestUpdateTestCollection(t *testing.T) {
	group := createTestGroup()

	err := AddTestGroup(group)
	if err != nil {
		t.Fatal("unable to create TestGroup: ", err.Error())
		t.FailNow()
	}

	col := group.TestCollections[0]
	col.Description = "This was updated for TestUpdateTestCollection"

	err = UpdateTestCollection(col)
	if err != nil {
		t.Fatal("Unable to update Test Collection: ", err.Error())
		t.FailNow()
	}
}

// func TestDeleteAllTestGroups(t *testing.T) {
// 	err := DeleteAllTestGroups()
// 	if err != nil {
// 		t.Log("Failed to delete all Test Groups: ", err.Error())
// 		t.Fail()
// 	}
// }
