package swagger

import (
	"net/http"
	"testing"
	"time"
	//json "encoding/json"
	//io "io/ioutil"
)

/*
func TestGetTestStatus(t *testing.T) {
	status := CoreGetTestStatus()

	t.Log(status.Cpu)
	t.Log(status.ExecutionTime)
	t.Log(status.RequestsExecuted)
	t.Log(status.SimulatedUsers)
	t.Log(status.TestCollectionName)
	t.Log(status.TransactionsPerSecond)
}

*/

type TestRoute struct {
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var (
	TestCountChen chan int
	TestRoutes    map[int]TestRoute
	srv           http.Server
)

const (
	BasePath string = "http://localhost:9001"
)

func MakeLittleServerToCall() {

	myMux := http.NewServeMux()

	for r := range TestRoutes {
		myMux.HandleFunc(TestRoutes[r].Pattern, TestRoutes[r].HandlerFunc)
	}

	srv.Handler = myMux

	srv.ListenAndServe()
}

var TestCount int

func init() {
	TestRoutes = make(map[int]TestRoute)
	srv = http.Server{Addr: ":9001"}
	TestRoutes[0] = TestRoute{Pattern: "/TestTestRunnerFirstBasic", HandlerFunc: func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		TestCount++
	}}

	go MakeLittleServerToCall()
}

func TestTestRunnerFirstBasic(t *testing.T) {
	CoreRunTest("TestRunnerFirst", 20)

	time.Sleep(time.Second * 20)

	CoreStopTest()

	t.Log("Tests completed: ", TestCount)

}

func TestKillServer(t *testing.T) {
	srv.Close()
}
