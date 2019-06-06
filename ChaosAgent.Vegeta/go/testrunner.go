package swagger

import (
	"errors"
	util "mallekoppie/ChaosGenerator/ChaosAgent/util"

	"log"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
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
	attacker                vegeta.Attacker
)

func init() {
	RunningSimulatedUsers = make(map[int32]bool)
	TestStatisticsChan = make(chan TestStatistics, 2000)
	go MonitorAndUpdateStatistics()
}

type TestStatistics struct {
	RequestsExecuted             int64
	TotalExecutionTimeNanosecond int64
	ErrorCount                   int32
}

func CoreGetTestStatus() TestStatus {
	testStatus := TestStatus{}

	testStatus.Cpu = util.GetCPUStatus() //Slow
	if ExecutionTimeNanosecond > 0 {
		testStatus.ExecutionTime = ExecutionTimeNanosecond / 1000000000
	}
	testStatus.RequestsExecuted = RequestsExecuted
	testStatus.SimulatedUsers = SimulatedUsers
	if RequestsExecuted > 0 {
		testStatus.AverageExecutionTime = ExecutionTimeNanosecond / RequestsExecuted / int64(SimulatedUsers) / 1000000
	}

	if testStatus.ExecutionTime > 0 {
		testStatus.TransactionsPerSecond = RequestsExecuted / testStatus.ExecutionTime
		testStatus.TransactionsPerSecond = testStatus.TransactionsPerSecond * int64(SimulatedUsers)
	}
	testStatus.TestCollectionName = TestCollectionName

	testStatus.ErrorsRaised = ErrorCount

	if ErrorCount > 0 && testStatus.ExecutionTime > 0 {
		testStatus.ErrorsPerSecond = int64(ErrorCount) / testStatus.ExecutionTime
	}

	return testStatus
}

func CoreStopTest() {
	if IsTestRunning == true {

		attacker.Stop()

		IsTestRunning = false
	}
}

func CoreRunTest(testName string, simulatedUsersInput int) (bool, error) {

	go StartVegeta(testName, simulatedUsersInput)

	return true, nil
}

func StartVegeta(testName string, simulatedUsersInput int) {
	rate := vegeta.Rate{Freq: simulatedUsersInput, Per: time.Second}

	duration := 0 * time.Second
	tests, _ := ReadTestConfiguration(testName)
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: tests.Tests[0].Method,
		URL:    tests.Tests[0].Url,
		Body:   []byte(tests.Tests[0].Body),
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
		testStats := TestStatistics{
			RequestsExecuted:             1,
			TotalExecutionTimeNanosecond: res.Latency.Nanoseconds(),
		}

		TestStatisticsChan <- testStats

	}
	metrics.Close()
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
	lastResetTime := time.Now()
	for true {
		testStats := <-TestStatisticsChan

		RequestsExecuted = RequestsExecuted + testStats.RequestsExecuted
		ExecutionTimeNanosecond = ExecutionTimeNanosecond + testStats.TotalExecutionTimeNanosecond
		ErrorCount = ErrorCount + testStats.ErrorCount

		timeSince := time.Since(lastResetTime)

		if timeSince.Seconds() > 10 {
			RequestsExecuted = 0
			ExecutionTimeNanosecond = 0
			ErrorCount = 0
			lastResetTime = time.Now()
		}

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

			if err != nil || responseCode != item.ResponseCode {
				testErrors++
				log.Printf("Error. Expected Code: %v but received: %v", item.ResponseCode, responseCode)
			} else if len(item.ResponseBody) > 0 && responseBody != item.ResponseBody {
				testErrors++
				log.Printf("Error. Expected body: %v but received: %v", item.ResponseBody, responseBody)
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
