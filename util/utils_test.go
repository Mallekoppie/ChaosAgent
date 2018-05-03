package util

import (
	"testing"
)

func TestGetCPU(t *testing.T) {

	result := GetCPUStatus()

	if result == 0 {
		t.Log("Result: ", result)
		t.Fail()
	}
}
