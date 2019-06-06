package manager

import (
	"context"
	"fmt"
	"log"
	"mallekoppie/ChaosGenerator/ChaosMaster/swagger"

	"github.com/tkanos/gonfig"
)

func GetConfig() (ChaosMasterConfig, error) {
	configuration := ChaosMasterConfig{}
	err := gonfig.GetConf("ChaosMasterConfig.json", &configuration)

	if err != nil {
		log.Print("Error reading config: %v", err)
		return configuration, err
	}

	return configuration, nil
}

func GetTest(testName string) (swagger.TestCollection, error) {
	configuration := swagger.TestCollection{}
	err := gonfig.GetConf("./tests/"+testName+".json", &configuration)

	if err != nil {
		log.Print("Error reading config: %v", err)
		return configuration, err
	}

	return configuration, nil
}

type ChaosMasterConfig struct {
	Agents []ChaosAgent `json:"agents,omitempty"`
}

type ChaosAgent struct {
	Name   string             `json:"name"`
	Url    string             `json:"url"`
	Client *swagger.APIClient `json:"client,omitempty"`
	Ctx    context.Context    `json:"ctx,omitempty"`
}

func (c *ChaosAgent) Init() {
	config := swagger.NewConfiguration()

	config.BasePath = c.Url
	c.Client = swagger.NewAPIClient(config)
	c.Ctx = context.TODO()
}

func (c *ChaosAgent) GetStatus() (swagger.TestStatus, error) {
	status, _, err := c.Client.DefaultApi.GetTestStatus(c.Ctx)

	if err != nil {
		//fmt.Println("Error calling service: ", err)
		return status, err
	} else {
		return status, nil
	}

}

func (c *ChaosAgent) IsAlive() bool {
	resp, err := c.Client.DefaultApi.IsAlive(c.Ctx)

	if err != nil {
		//fmt.Printf("Error checking if %v is alive. Error: %v", c.Name, err)
		return false
	} else if resp != nil && resp.StatusCode == 200 {
		return true
	}

	return false

}

func (c *ChaosAgent) AddTest(test swagger.TestCollection) {
	resp, err := c.Client.DefaultApi.AddTests(c.Ctx, test)

	if err != nil {
		fmt.Printf("Error adding test to %v . Error: %v", c.Name, err)
		return
	}

	if resp != nil && resp.StatusCode != 200 {
		fmt.Printf("Error adding test for %v . ResponseCode: ", c.Name, resp.StatusCode)
	}
}

func (c *ChaosAgent) StartTest(testParameters swagger.TestParameters) {
	resp, err := c.Client.DefaultApi.StartTestRun(c.Ctx, testParameters)

	if err != nil {
		fmt.Printf("Error starting test to %v . Error: %v", c.Name, err)
	}

	if resp != nil && resp.StatusCode != 200 {
		fmt.Printf("Error starting test for %v . ResponseCode: ", c.Name, resp.StatusCode)
	}
}

func (c *ChaosAgent) UpdateTest(testParameters swagger.TestParameters) {
	resp, err := c.Client.DefaultApi.UpdateTestRun(c.Ctx, testParameters)

	if err != nil {
		fmt.Printf("Error updating test to %v . Error: %v", c.Name, err)
	}

	if resp != nil && resp.StatusCode != 200 {
		fmt.Printf("Error updating test for %v . ResponseCode: ", c.Name, resp.StatusCode)
	}
}

func (c *ChaosAgent) StopTest() {
	if c != nil && c.Client != nil && c.Client.DefaultApi != nil {
		resp, err := c.Client.DefaultApi.StopTestRun(c.Ctx, "")

		if err != nil {
			fmt.Printf("Error stopping test to %v . Error: %v", c.Name, err)
		}

		if resp != nil && resp.StatusCode != 200 {
			fmt.Printf("Error stopping test for %v . ResponseCode: ", c.Name, resp.StatusCode)
		}
	}
}
