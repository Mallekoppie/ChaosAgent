package util

import (
	"log"
	"net/http"
	"testing"
	//"time"
)

/*
// Function is slow. Commenting out while testing http functions
func TestGetCPU(t *testing.T) {

	result := GetCPUStatus()

	if result == 0 {
		t.Log("Result: ", result)
		t.Fail()
	}
}
*/

type Route struct {
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var (
	HeaderTestChan chan bool
	Routes         map[int]Route
	srv            http.Server
)

const (
	BasePath string = "http://localhost:9999"
)

func init() {
	HeaderTestChan = make(chan bool, 10)
	Routes = make(map[int]Route)
	srv = http.Server{Addr: ":9999"}
	Routes[0] = Route{Pattern: "/TestHttpCallForBasicGet", HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Println("Request received")
	}}
	Routes[1] = Route{Pattern: "/TestHttpCallForToSendHeaders", HandlerFunc: func(w http.ResponseWriter, r *http.Request) {

		log.Println("HeaderTestRequest received")

		first := false
		second := false
		for header := range r.Header {

			if header == "First" {
				if r.Header.Get(header) == "value" {
					first = true
				}
			} else if header == "Second" {
				if r.Header.Get(header) == "secondvalue" {
					second = true
				}
			}
		}

		if first == true && second == true {
			HeaderTestChan <- true
		} else {
			HeaderTestChan <- false
		}

		w.WriteHeader(http.StatusOK)
		log.Println("Handler Done")
	}}

}

func MakeLittleServerToCall(t *testing.T) {

	myMux := http.NewServeMux()

	for r := range Routes {
		myMux.HandleFunc(Routes[r].Pattern, Routes[r].HandlerFunc)
	}

	srv.Handler = myMux

	srv.ListenAndServe()
}

func TestHttpCallForBasicGet(t *testing.T) {
	go MakeLittleServerToCall(t)

	headers := make(map[string]string)
	responseCode, _, _ := MakeHttpCall(BasePath+"/TestHttpCallForBasicGet", http.MethodGet, headers, "")

	if responseCode != http.StatusOK {
		t.Log("Incorrect response code received")
		t.Fail()
	}

	t.Log("Received correct response code")
}

func TestHttpCallForToSendHeaders(t *testing.T) {
	go MakeLittleServerToCall(t)

	headers := make(map[string]string)
	headers["first"] = "value"
	headers["second"] = "secondvalue"

	responseCode, _, _ := MakeHttpCall(BasePath+"/TestHttpCallForToSendHeaders", http.MethodGet, headers, "")

	if responseCode != http.StatusOK {
		t.Log("Incorrect response code received")
		t.Fail()
	}

	t.Log("Waiting for Channel")
	correctheadersReceived := <-HeaderTestChan

	if correctheadersReceived != true {
		t.Log("Incorrect headers received")
		t.Fail()
	}
}

/*
func TestExternalTest(t *testing.T) {
	headers := make(map[string]string)
	headers["first"] = "value"
	headers["test"] = "secondvalue"

	responseCode, _, _ := MakeHttpCall("http://localhost:80/ConnectionTest", http.MethodGet, headers, "")

	if responseCode != http.StatusOK {
		t.Log("Incorrect response code received")
		t.Fail()
	}
}
*/
func TestKillServer(t *testing.T) {
	srv.Close()
}
