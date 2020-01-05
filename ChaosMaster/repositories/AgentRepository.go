package repositories

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mallekoppie/ChaosGenerator/ChaosMaster/models"
	"net/http"
	"strconv"
	"time"
	//	"github.com/google/uuid"
)

var (
	httpClient                       *http.Client
	DefaultTLSConfig                 = &tls.Config{InsecureSkipVerify: true}
	ConsulUrl                        string
	ConsulToken                      string
	ErrorConsulPortIncorrect         = errors.New("Consul port incorrect")
	ErrorConsulHostIncorrect         = errors.New("Consul host incorrect")
	ErrorConsulResponseCodeIncorrect = errors.New("Consul returned incorrect response code")
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 30
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()

	ServiceConfig, _ = GetConfig()
	ConsulUrl = ServiceConfig.ConsulUrl
	ConsulToken = ServiceConfig.ConsulToken
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
			TLSClientConfig:     DefaultTLSConfig,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func createConsulRequest(port int, host string, enabled bool, serviceName string, id string) (request models.ConsulRequest, err error) {
	consulRequest := models.ConsulRequest{}

	if port < 1024 || port > 65200 {
		log.Println("Bad port value for consul request: ", port)
		return consulRequest, ErrorConsulPortIncorrect
	}

	if len(host) < 1 {
		log.Println("Bad host value for consul request: ", host)
		return consulRequest, ErrorConsulHostIncorrect
	}

	consulRequest.Service.Service = serviceName
	consulRequest.Service.Port = port
	consulRequest.Node = fmt.Sprintf("%v:%v:%v", serviceName, host, port)
	consulRequest.NodeMeta.Enabled = strconv.FormatBool(enabled)
	consulRequest.NodeMeta.Id = id
	consulRequest.Address = host

	return consulRequest, nil
}

func UpdateChaosAgent(agent models.Agent, serviceName string, port int) error {

	requestObject, err := createConsulRequest(port, agent.Host, agent.Enabled, serviceName, agent.Id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(requestObject)
	if err != nil {
		log.Println("Unable to marchall agent to json for update: ", err.Error())
		return err
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/catalog/register", ConsulUrl), bytes.NewBuffer(data))
	if err != nil {
		log.Println("Unable to create new HTTP request for Consul update: ", err.Error())
		return err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		log.Println("Error while sending request to Consul: ", err.Error())
		return err
	}

	if response.StatusCode != http.StatusOK {
		log.Println("Incorrect response code from consul for agent registration: ", response.StatusCode)
		return ErrorConsulResponseCodeIncorrect
	}

	return nil
}

func DeleteChaosAgent(agent models.Agent, serviceName string) error {
	requestObject, err := createConsulRequest(agent.Port, agent.Host, agent.Enabled, serviceName, agent.Id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(requestObject)
	if err != nil {
		log.Println("Unable to marchall agent to json for update: ", err.Error())
		return err
	}

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/catalog/deregister", ConsulUrl), bytes.NewBuffer(data))
	if err != nil {
		log.Println("Unable to create new HTTP request for Consul update: ", err.Error())
		return err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		log.Println("Error while sending request to Consul: ", err.Error())
		return err
	}

	if response.StatusCode != http.StatusOK {
		log.Println("Incorrect response code from consul for agent registration: ", response.StatusCode)
		return ErrorConsulResponseCodeIncorrect
	}

	return nil
}

func GetAllAgents(serviceName string) (agents []models.Agent, err error) {

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/v1/catalog/service/%v", ConsulUrl, serviceName), nil)
	if err != nil {
		log.Println("Error creating request to retrieve all consul agents: ", err.Error())
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		log.Println("Unable to make call to consul: ", err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		log.Println("Consul returned incorrect response code. Expected 200 but received ", response.StatusCode)
		return nil, ErrorConsulResponseCodeIncorrect
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body: ", err.Error())
		return nil, err
	}
	consulAgentResponse := models.ConsulAgentResponse{}
	err = json.Unmarshal(data, &consulAgentResponse)
	if err != nil {
		log.Println("Error unmarshalling consul agents: ", err.Error())
		return nil, err
	}

	agents = make([]models.Agent, 0)
	for i := range consulAgentResponse {

		agent := models.Agent{
			Id:   consulAgentResponse[i].NodeMeta.Id,
			Host: consulAgentResponse[i].Address,
			Port: consulAgentResponse[i].ServicePort,
		}
		enabled, err := strconv.ParseBool(consulAgentResponse[i].NodeMeta.Enabled)
		if err != nil {
			log.Println("unable to parse agent enabled status: ", err.Error())
			enabled = false
		} else {
			agent.Enabled = enabled
		}

		agents = append(agents, agent)
	}

	return agents, nil
}
