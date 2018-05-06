package swagger

import (
	util "mallekoppie/ChaosAgent/util"
)

var (
	ExecutionTime         int32
	RequestsExecuted      int32
	SimulatedUsers        int32
	TransactionsPerSecond int32
	TestCollectionName    string
	IsTestRunning         bool
	RunningSimulatedUsers map[int]bool
)

func init() {
	RunningSimulatedUsers = make(map[int]bool)
}

func CoreGetTestStatus() TestStatus {
	testStatus := TestStatus{}

	testStatus.Cpu = util.GetCPUStatus()
	testStatus.ExecutionTime = ExecutionTime
	testStatus.RequestsExecuted = RequestsExecuted
	testStatus.SimulatedUsers = SimulatedUsers
	testStatus.TransactionsPerSecond = TransactionsPerSecond
	testStatus.TestCollectionName = TestCollectionName

	return testStatus
}

func CoreRunTest(testName string, simulatedUsersInput int) (bool, error) {

	testCollection, configError := ReadTestConfiguration(testName)

	if configError != nil {
		IsTestRunning = false
		return IsTestRunning, configError
	}

	IsTestRunning = true

	go RunTest(testCollection)

	return IsTestRunning, nil
}

func RunTest(config TestCollection) {

	for IsTestRunning == true {
		for testIndex := range config.Tests {
			if IsTestRunning == false {
				break
			}

			item := config.Tests[testIndex]
			headers := make(map[string]string)

			if len(item.Headers) > 0 {
				for h := range item.Headers {
					headers[item.Headers[h].Name] = item.Headers[h].Value
				}
			}

			responseCode, responseBody, err := util.MakeHttpCall(item.Url, item.Method, headers, item.Body)

			if err != nil {
				// log error and change statistics
			}

			if responseCode != item.ResponseCode {
				// log error and change statistics
			}

			if responseBody != item.ResponseBody {
				// log error and change statistics
			}

		}
	}

}
