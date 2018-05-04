package swagger

import (
	util "mallekoppie/ChaosAgent/util"
)

// Test statistics variables
var (
	executionTime         int32
	requestsExecuted      int32
	simulatedusers        int32
	transactionsPerSecond int32
	testCollectionName    string
)

func CoreGetTestStatus() TestStatus {
	testStatus := TestStatus{}

	testStatus.Cpu = util.GetCPUStatus()
	testStatus.ExecutionTime = executionTime
	testStatus.RequestsExecuted = requestsExecuted
	testStatus.SimulatedUsers = simulatedusers
	testStatus.TransactionsPerSecond = transactionsPerSecond
	testStatus.TestCollectionName = testCollectionName

	return testStatus
}
