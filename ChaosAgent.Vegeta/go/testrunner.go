package swagger

import (
	"log"
	"net/http"
	"time"

	cpu "github.com/shirou/gopsutil/cpu"
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
	attacker                *vegeta.Attacker
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

	testStatus.Cpu = GetCPUStatus() //Slow
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

	if IsTestRunning == false {
		go StartVegeta(testName, simulatedUsersInput)
	}
	return true, nil
}

func StartVegeta(testName string, simulatedUsersInput int) {
	rate := vegeta.Rate{Freq: simulatedUsersInput, Per: time.Second}

	duration := 0 * time.Second
	tests, _ := ReadTestConfiguration(testName)
	//targeter := vegeta.NewStaticTargeter(vegeta.Target{
	//	Method: tests.Tests[0].Method,
	//	URL:    tests.Tests[0].Url,
	//	Body:   []byte(tests.Tests[0].Body),
	//})

	targets := make([]vegeta.Target, 0)

	for testIndex := range tests.Tests {
		headers := http.Header{}
		if len(tests.Tests[testIndex].Headers) > 0 {
			for headerIndex := range tests.Tests[testIndex].Headers {
				headers.Add(tests.Tests[testIndex].Headers[headerIndex].Name, tests.Tests[testIndex].Headers[headerIndex].Value)
			}
		}

		targets = append(targets, vegeta.Target{
			Method: tests.Tests[testIndex].Method,
			URL:    tests.Tests[testIndex].Url,
			Body:   []byte(tests.Tests[testIndex].Body),
			Header: headers,
		})
	}

	targeter := vegeta.NewStaticTargeter(targets...)

	attacker = vegeta.NewAttacker()

	var metrics vegeta.Metrics
	IsTestRunning = true
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
		testStats := TestStatistics{
			RequestsExecuted:             1,
			TotalExecutionTimeNanosecond: res.Latency.Nanoseconds(),
			ErrorCount:                   0,
		}

		TestStatisticsChan <- testStats

	}
	metrics.Close()
}

func CoreUpdateTest(simulatedUsersInput int32) error {

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

		if timeSince.Seconds() > 20 {
			RequestsExecuted = 0
			ExecutionTimeNanosecond = 0
			ErrorCount = 0
			lastResetTime = time.Now()
		}

	}
}

func GetCPUStatus() float64 {
	var cpuUsage float64
	data, cpuStatusErr := cpu.Times(true)
	var valueRetrieved bool

	for maxRetries := 0; maxRetries < 5; maxRetries++ {

		for i := range data {
			if data[i].CPU == "_Total" {
				cpuUsage = data[i].User
				valueRetrieved = true
				break
			}
		}

		if cpuStatusErr != nil || len(data) < 1 {
			if cpuStatusErr != nil {
				log.Println("Error retrieving CPU stats:", cpuStatusErr)
			}

			data, cpuStatusErr = cpu.Times(true)
			continue
		}
	}

	if valueRetrieved == false {
		cpuUsage = -1
	}

	return cpuUsage
}
