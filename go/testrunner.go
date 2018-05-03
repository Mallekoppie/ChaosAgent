package swagger

import (
	util "mallekoppie/ChaosAgent/util"
	"net/http"
)

// Test statistics variables
var (
	executionTime         int32
	requestsExecuted      int32
	simulatedusers        int32
	transactionsPerSecond int32
	testCollectionName    string
)

var (
	httpClient *http.Client
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func MakeCall() {
	var endPoint string = "http://localhost:80/ConnectionTest"

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		log.Fatalf("Error Occured. %+v", err)
	}

	// use httpClient to send request
	response, err := httpClient.Do(req)
	if err != nil && response == nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	} else {
		// Close the connection to reuse it
		defer response.Body.Close()

		// Let's check if the work actually is done
		// We have seen inconsistencies even when we get 200 OK response
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Couldn't parse response body. %+v", err)
		}

		//log.Println("Response Body:", string(body))
	}
}

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
