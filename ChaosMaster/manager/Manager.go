package manager

import (
	"context"
	"fmt"
	"log"

	"github.com/tkanos/gonfig"

	pb "mallekoppie/ChaosGenerator/Chaos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func GetTest(testName string) (pb.TestCollection, error) {
	configuration := pb.TestCollection{}
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
	Name   string `json:"name"`
	Url    string `json:"url"`
	Client pb.ChaosAgentClient
	Ctx    context.Context
}

func (c *ChaosAgent) Init() error {
	creds, err := credentials.NewClientTLSFromFile("./chaos_agent.cer", "chaos-agent")
	if err != nil {
		log.Println("Error reading certificate: ", err.Error())
		return err
	}

	conn, err := grpc.Dial(c.Url, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Printf("Unable to connect to %v. Error: %v", c.Url, err.Error())
		return err
	}

	c.Client = pb.NewChaosAgentClient(conn)

	resp, err := c.Client.GetVersion(c.Ctx, &pb.Request{})
	if err != nil {
		log.Printf("Error during version check to %v. Error: %v", c.Url, err.Error())
		return err
	}

	log.Println("Agent at %v si online with version %v", c.Url, resp.GetVersion())

	return nil
}

func (c *ChaosAgent) GetStatus() (pb.TestStatus, error) {

	status, err := c.Client.GetTestStatus(c.Ctx, &pb.Request{})

	if err != nil {
		//fmt.Println("Error calling service: ", err)
		return *status, err
	} else {
		return *status, nil
	}

}

func (c *ChaosAgent) IsAlive() bool {
	response, err := c.Client.IsAlive(c.Ctx, &pb.Request{})

	if err != nil {
		//fmt.Printf("Error checking if %v is alive. Error: %v", c.Name, err)
		return false
	} else if response != nil && response.Result == true {
		return true
	}

	return false

}

func (c *ChaosAgent) AddTest(test pb.TestCollection) {
	resp, err := c.Client.AddTests(c.Ctx, &test)

	if err != nil {
		fmt.Printf("Error adding test to %v . Error: %v", c.Name, err)
		return
	}

	if resp != nil && resp.Result != true {
		fmt.Printf("Error adding test for %v . Result: ", c.Name, resp.Result)
	}
}

func (c *ChaosAgent) StartTest(testParameters pb.TestParameters) {
	resp, err := c.Client.StartTestRun(c.Ctx, &testParameters)

	if err != nil {
		fmt.Printf("Error starting test to %v . Error: %v", c.Name, err)
	}

	if resp != nil && resp.Result != true {
		fmt.Printf("Error starting test for %v . Result: ", c.Name, resp.Result)
	}
}

func (c *ChaosAgent) UpdateTest(testParameters pb.TestParameters) {
	resp, err := c.Client.UpdateTestRun(c.Ctx, &testParameters)

	if err != nil {
		fmt.Printf("Error updating test to %v . Error: %v", c.Name, err)
	}

	if resp != nil && resp.Result != true {
		fmt.Printf("Error updating test for %v . Result: ", c.Name, resp.Result)
	}
}

func (c *ChaosAgent) StopTest() {
	if c != nil && c.Client != nil && c.Client != nil {
		resp, err := c.Client.StopTestRun(c.Ctx, &pb.StopTestRequest{})

		if err != nil {
			fmt.Printf("Error stopping test to %v . Error: %v", c.Name, err)
		}

		if resp != nil && resp.Result != true {
			fmt.Printf("Error stopping test for %v . Result: ", c.Name, resp.Result)
		}
	}
}
