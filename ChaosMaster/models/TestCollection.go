package models

type TestCollection struct {
	TestCollectionId string `json:"testCollectionId"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Tests            []Test `json:"tests"`
	GroupId          string `json:"groupId"`
}
