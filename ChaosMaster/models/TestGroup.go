package models

type TestGroup struct {
	TestGroupId     string           `json:"testGroupId"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	TestCollections []TestCollection `json:"testCollections"`
}
