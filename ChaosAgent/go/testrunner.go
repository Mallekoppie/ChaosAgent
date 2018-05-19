package swagger

import (
	"errors"
	util "mallekoppie/ChaosAgent/util"
	"time"
)

var (
	ExecutionTimeNanosecond int64
	RequestsExecuted        int64
	SimulatedUsers          int32
	ErrorCount              int32
	TestCollectionName      string
	IsTestRunning           bool
	RunningSimulatedUsers   map[int32]bool
	SampleIntervalSecond    int32 = 10
	TestStatisticsChan      chan TestStatistics
)

func init() {
	RunningSimulatedUsers = make(map[int32]bool)
	TestStatisticsChan = make(chan TestStatistics, 1000)
}

type TestStatistics struct {
	RequestsExecuted             int64
	TotalExecutionTimeNanosecond int64
	ErrorCount                   int32
}

func CoreGetTestStatus() TestStatus {
	testStatus := TestStatus{}

	//testStatus.Cpu = util.GetCPUStatus() //Slow
	if ExecutionTimeNanosecond > 0 {
		testStatus.ExecutionTime = ExecutionTimeNanosecond / 1000000000
	}
	testStatus.RequestsExecuted = RequestsExecuted
	testStatus.SimulatedUsers = SimulatedUsers
	if RequestsExecuted > 0 {
		testStatus.AverageExecutionTime = ExecutionTimeNanosecond / RequestsExecuted / 1000000
	}

	if testStatus.ExecutionTime > 0 {
		testStatus.TransactionsPerSecond = RequestsExecuted / testStatus.ExecutionTime
	}
	testStatus.TestCollectionName = TestCollectionName

	return testStatus
}

func CoreStopTest() {
	if IsTestRunning == true {

		for SimulatedUsers > 0 {
			RunningSimulatedUsers[SimulatedUsers] = false

			SimulatedUsers--
		}

		IsTestRunning = false
	}
}

func CoreRunTest(testName string, simulatedUsersInput int) (bool, error) {

	if IsTestRunning == true {
		return IsTestRunning, errors.New("Test is already running")
	}

	testCollection, configError := ReadTestConfiguration(testName)

	if configError != nil {
		IsTestRunning = false
		return IsTestRunning, configError
	}

	IsTestRunning = true
	SimulatedUsers = 0
	ExecutionTimeNanosecond = 0
	RequestsExecuted = 0
	ErrorCount = 0
	go MonitorAndUpdateStatistics()
	for i := 0; i < simulatedUsersInput; i++ {
		RunningSimulatedUsers[SimulatedUsers] = true
		go RunTest(testCollection, SimulatedUsers)
		SimulatedUsers++
	}

	return IsTestRunning, nil
}

func CoreUpdateTest(simulatedUsersInput int32) error {
	if IsTestRunning == false {
		return errors.New("No Test is running")
	}

	if simulatedUsersInput < SimulatedUsers-1 {
		return nil
	}

	testCollection, configError := ReadTestConfiguration(TestCollectionName)

	if configError != nil {
		return errors.New("Could not retrieve test config")
	}

	for SimulatedUsers < simulatedUsersInput {
		RunningSimulatedUsers[SimulatedUsers] = true
		go RunTest(testCollection, SimulatedUsers)
		SimulatedUsers++
	}

	return nil
}

func MonitorAndUpdateStatistics() {
	for IsTestRunning == true {
		testStats := <-TestStatisticsChan

		RequestsExecuted = RequestsExecuted + testStats.RequestsExecuted
		ExecutionTimeNanosecond = ExecutionTimeNanosecond + testStats.TotalExecutionTimeNanosecond

	}
}

func RunTest(config TestCollection, index int32) {

	for RunningSimulatedUsers[index] == true {

		for testIndex := range config.Tests {
			var testErrors int32
			var testTimeNanosecond, testsCompleted int64
			if RunningSimulatedUsers[index] == false {
				break
			}

			item := config.Tests[testIndex]
			headers := make(map[string]string)

			if len(item.Headers) > 0 {
				for h := range item.Headers {
					headers[item.Headers[h].Name] = item.Headers[h].Value
				}
			}

			startTime := time.Now()
			responseCode, responseBody, err := util.MakeHttpCall(item.Url, item.Method, headers, item.Body)
			result := time.Since(startTime)
			testTimeNanosecond = result.Nanoseconds()
			testsCompleted++

			if err != nil || responseCode != item.ResponseCode || responseBody != item.ResponseBody {
				testErrors++
			}

			testStats := TestStatistics{
				RequestsExecuted:             testsCompleted,
				TotalExecutionTimeNanosecond: testTimeNanosecond,
				ErrorCount:                   testErrors,
			}

			TestStatisticsChan <- testStats
		}
	}
}
