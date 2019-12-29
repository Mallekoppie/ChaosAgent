package repositories

import (
	"fmt"
	"mallekoppie/ChaosGenerator/ChaosMaster/models"
	"testing"
)

func TestFormatting(t *testing.T) {

	result := fmt.Sprintf("mongodb://%v:%v", "localhost", 27017)

	t.Log(result)

}

func createTestGroup() models.TestGroup {
	testCollections := make([]models.TestCollection, 0)

	tests := make([]models.Test, 0)
	tests = append(tests, models.Test{TestId: "test"})

	col1 := models.TestCollection{TestCollectionId: "bla",
		Name: "some name"}
	testCollections = append(testCollections, col1)

	testGroupId := "Some unique Id. Ideally a guid"
	testGroup := models.TestGroup{
		TestGroupId:     testGroupId,
		Name:            "Unit Test Name",
		Description:     "a Nice Description",
		TestCollections: testCollections,
	}

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

func TestDeleteTestCollection(t *testing.T) {
	idToDelete := "IDTODELETE"
	testGroup := createTestGroup()
	testGroup.TestGroupId = idToDelete
	col1 := models.TestCollection{TestCollectionId: idToDelete,
		Name: "some name"}

	testGroup.TestCollections = append(testGroup.TestCollections, col1)

	err := AddTestGroup(testGroup)
	if err != nil {
		t.Log("Unable to add test group for delete test: ", err.Error())
		t.Fail()
	}

	err = DeleteTestCollection(idToDelete, idToDelete)
	if err != nil {
		t.Log("Unable to delete test collection: ", err.Error())
		t.Fail()
	}
}

// func TestDeleteAllTestGroups(t *testing.T) {
// 	err := DeleteAllTestGroups()
// 	if err != nil {
// 		t.Log("Failed to delete all Test Groups: ", err.Error())
// 		t.Fail()
// 	}
// }
