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

	tests := make([]models.Test, 0)
	tests = append(tests, models.Test{ID: "test"})

	col1 := models.TestCollection{ID: "bla",
		Name: "some name"}
	testCollections = append(testCollections, col1)

	testGroupId := uuid.New().String()
	testGroup := models.TestGroup{

		ID:              testGroupId,
		Name:            "Unit Test Name",
		Description:     "a Nice Description",
		TestCollections: testCollections,
	}
	uid := uuid.New()
	testGroup.ID = uid.String()

	return testGroup
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
	AddTestGroup(testGroup)

	testGroup.Description = "This has been updated"
	err := UpdateTestGroup(testGroup)
	if err != nil {
		t.Fatalf("Error updating record: %v", err.Error())
		t.Fail()
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
	}
}

// func TestDeleteAllTestGroups(t *testing.T) {
// 	err := DeleteAllTestGroups()
// 	if err != nil {
// 		t.Log("Failed to delete all Test Groups: ", err.Error())
// 		t.Fail()
// 	}
// }
