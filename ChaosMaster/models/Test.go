package models

type Test struct {
	TestId       string   `json:"testId"`
	Name         string   `json:"name"`
	Method       string   `json:"method"`
	Url          string   `json:"url"`
	Body         string   `json:"body"`
	Headers      []Header `json:"headers"`
	ResponseCode int      `json:"responseCode"`
	ResponseBody string   `json:"responseBody"`
}

type Header struct {
	HeaderId string `json:"headerId"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}
