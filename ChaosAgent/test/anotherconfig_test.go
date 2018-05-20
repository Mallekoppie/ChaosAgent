package test

import (
	"mallekoppie/ChaosGenerator/ChaosAgent/go"
	"net/http"
	"testing"
)

func init() {
	config := swagger.TestCollection{Name: "test"}
	swagger.WriteTestConfiguration(config)
}

func TestReadConfig(t *testing.T) {
	test, err := swagger.ReadTestConfiguration("test")

	if err != nil {
		t.Logf("Failed to read config: %v", err)
		t.Fail()
	}

	if len(test.Name) > 0 {
		t.Logf("Test Collection retrieved: %v", test.Name)
	}

}

func TestWriteConfig(t *testing.T) {
	config := swagger.TestCollection{Name: "base64bodyTest"}
	config.Tests = []swagger.Test{{
		Name:         "BasicGet",
		Method:       "GET",
		Body:         "",
		Url:          "http://localhost:9001/TestTestRunnerFirstBasic",
		ResponseCode: http.StatusOK,
	},
	}

	swagger.WriteTestConfiguration(config)
}

// TODO: Write test where body is base64 encoded
func TestWriteConfigWithBase64Body(t *testing.T) {
	config := swagger.TestCollection{Name: "test"}

	swagger.WriteTestConfiguration(config)
}

func TestCreateTestRunnerConfig(t *testing.T) {
	config := swagger.TestCollection{Name: "TestRunnerFirst"}
	config.Tests = []swagger.Test{{
		Name:         "BasicGet",
		Method:       "GET",
		Body:         "",
		Url:          "http://localhost:9001/TestTestRunnerFirstBasic",
		ResponseCode: http.StatusOK,
	},
	}

	swagger.WriteTestConfiguration(config)
}
