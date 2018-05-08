package swagger

import (
	"net/http"
	"testing"
)

func init() {
	config := TestCollection{Name: "test"}
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
	config := TestCollection{Name: "test"}
	WriteTestConfiguration(config)
}

// TODO: Write test where body is base64 encoded

func TestCreateTestRunnerConfig(t *testing.T) {
	config := TestCollection{Name: "TestRunnerFirst"}
	config.Tests = []Test{{
		Name:         "BasicGet",
		Method:       "GET",
		Body:         "",
		Url:          "http://localhost:9001/TestTestRunnerFirstBasic",
		ResponseCode: http.StatusOK,
	},
	}

	WriteTestConfiguration(config)
}
