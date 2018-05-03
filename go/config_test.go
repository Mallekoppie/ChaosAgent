package swagger

import (
	"testing"
)

func init() {
	config := swagger.TestCollection{Name: "test"}
	WriteTestConfiguration(config)
}

func TestReadConfig(t *testing.T) {
	test, err := ReadTestConfiguration("test")

	if err != nil {
		t.Logf("Failed to read config: %v", err)
		t.Fail()
	}

	if len(test.Name) > 0 {
		t.Logf("Test Collection retrieved: %v", test.Name)
	}

}

func TestWriteConfig(t *testing.T) {
	config := swagger.TestCollection{Name: "test"}
	WriteTestConfiguration(config)
}
