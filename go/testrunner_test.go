package swagger

import (
	"testing"
)

func TestGetTestStatus(t *testing.T) {
	status := CoreGetTestStatus()

	t.Log(status.Cpu)
	t.Log(status.ExecutionTime)
	t.Log(status.RequestsExecuted)
	t.Log(status.SimulatedUsers)
	t.Log(status.TestCollectionName)
	t.Log(status.TransactionsPerSecond)
}
