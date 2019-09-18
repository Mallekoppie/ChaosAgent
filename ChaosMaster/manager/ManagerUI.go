package manager

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"mallekoppie/ChaosGenerator/ChaosMaster/util"
	"os"
	"os/exec"
	"strconv"
	"strings"

	pb "mallekoppie/ChaosGenerator/Chaos"
)

var (
	reader bufio.Scanner
	Config ChaosMasterConfig
)

func init() {
	reader = *bufio.NewScanner(os.Stdin)
	var err error
	Config, err = GetConfig()

	if err != nil {
		fmt.Println("Error retrieving config: ", err)
		return
	}

	for i := 0; i < len(Config.Agents); i++ {
		Config.Agents[i].Init()
	}
}

func Prnt(s string) {
	fmt.Println(s)
}

func OptionClearOutput() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func DisplayManagerUI() {

	Prnt("Menu:")
	Prnt("a       = Check if Agents are alive")
	Prnt("start   = Start tests")
	Prnt("stop    = Stop tests")
	Prnt("stopall = Stop tests for all Agents")
	Prnt("s       = Status of runners")
	Prnt("ua      = Update Test Parameters for all")
	Prnt("u       = Update Test Parameters for specific Agent")
	Prnt("deploy  = Deploy Test Collection for all Agents")
	Prnt("d       = Display Configured servers")
	Prnt("l       = Display test collection files")
	Prnt("c       = Clear")
	Prnt("convert = Convert contents of files in conversions")
	Prnt("q       = Quit")
	Prnt("")
}

func ReadInput() string {
	reader.Scan()

	if reader.Err() != nil {
		return "Error reading input"
	}

	return reader.Text()
}

func RunUI() {
	quit := false

	for quit == false {
		DisplayManagerUI()
		result := ReadInput()

		switch strings.ToLower(result) {
		case "a":
			OptionIsAgentsAlive()
		case "start":
			OptionStartTest()
		case "stop":
			OptionStopOne()
		case "stopall":
			OptionStopAllTests()
		case "s":
			OptionAgentStatusAll()
		case "ua":
			OptionUpdateTestParametersForAllAgents()
		case "u":
			OptionUpdateTestParametersForOneAgent()
		case "deploy":
			OptionDeployTestToAllAgents()
		case "d":
			OptionDisplayConfigration()
		case "l":
			OptionListTestCollectionFiles()
		case "c":
			OptionClearOutput()
		case "convert":
			util.ConvertFileContentsToBase64()
		case "q":
			quit = true
		default:
			Prnt("Command not supported")
		}

	}
}

func OptionListTestCollectionFiles() {
	files, err := ioutil.ReadDir("tests")

	if err != nil {
		fmt.Println("Error reading test collection folder: ", err)
		return
	}

	for i := 0; i < len(files); i++ {
		filename := files[i].Name()
		result := strings.Index(filename, ".json")
		fmt.Println(filename[0:result])
	}
}

func OptionUpdateTestParametersForOneAgent() {

	OptionClearOutput()
	fmt.Println("Enter the Chaos Agent to update:")
	for i := 0; i < len(Config.Agents); i++ {
		fmt.Println(Config.Agents[i].Name)
	}

	reader.Scan()
	agentName := reader.Text()

	fmt.Println("Enter the number of users to simulate:")
	reader.Scan()
	numberOfUsers := reader.Text()

	numberOfUsersInt, convertErr := strconv.Atoi(numberOfUsers)

	if convertErr != nil {
		fmt.Println("Error converting your option to int32:", convertErr)
		return
	}

	test := pb.TestParameters{Simulatedusers: int32(numberOfUsersInt)}

	for i := 0; i < len(Config.Agents); i++ {

		if Config.Agents[i].Name == agentName {
			Config.Agents[i].UpdateTest(test)

			fmt.Printf("Agent %v test pushed \n", Config.Agents[i].Name)
		}
	}
}

func OptionUpdateTestParametersForAllAgents() {

	fmt.Println("Enter the number of users to simulate:")
	reader.Scan()
	numberOfUsers := reader.Text()

	numberOfUsersInt, convertErr := strconv.Atoi(numberOfUsers)

	if convertErr != nil {
		fmt.Println("Error converting your option to int32:", convertErr)
		return
	}

	test := pb.TestParameters{Simulatedusers: int32(numberOfUsersInt)}

	for i := 0; i < len(Config.Agents); i++ {
		Config.Agents[i].UpdateTest(test)

		fmt.Printf("Agent %v test pushed \n", Config.Agents[i].Name)
	}

}

func OptionStartTest() {

	fmt.Println("Enter the test collection name:")
	reader.Scan()
	testName := reader.Text()

	_, getTestErr := GetTest(testName)

	if getTestErr != nil {
		fmt.Println("The test does not exist locally: ", getTestErr)
		return
	}

	fmt.Println("Enter the number of users to simulate:")
	reader.Scan()
	numberOfUsers := reader.Text()

	numberOfUsersInt, convertErr := strconv.Atoi(numberOfUsers)

	if convertErr != nil {
		fmt.Println("Error converting your option to int32:", convertErr)
		return
	}

	test := pb.TestParameters{Simulatedusers: int32(numberOfUsersInt), TestCollectionName: testName}

	for i := 0; i < len(Config.Agents); i++ {
		Config.Agents[i].StartTest(test)

		fmt.Printf("Agent %v test pushed \n", Config.Agents[i].Name)
	}
}

func OptionDeployTestToAllAgents() {
	var name string
	for len(name) < 1 {
		fmt.Println("Enter the name of the test that you want to distribute: ")
		reader.Scan()
		name = reader.Text()
	}

	test, configErr := GetTest(name)

	if configErr != nil {
		// It is already displayed so we can just stop
		return
	}

	for i := 0; i < len(Config.Agents); i++ {
		Config.Agents[i].AddTest(test)

		fmt.Printf("Agent %v test pushed \n", Config.Agents[i].Name)
	}
}

func OptionAgentStatusAll() {
	for i := 0; i < len(Config.Agents); i++ {
		status, err := Config.Agents[i].GetStatus()

		if err != nil {
			//fmt.Printf("Error getting status for agent %v . Error: %v \n", Config.Agents[i].Name, err)
			fmt.Printf("Agent %v is offline\n", Config.Agents[i].Name)
		} else {
			fmt.Printf("Agent %v Test: %v \tTPS: %v \tEPS: %v \tUsers: %v \tTR: %v \tErrors: %v \tExecutionTime: %v \tCPU: %v  \n", Config.Agents[i].Name, status.TestCollectionName, status.TransactionsPerSecond, status.ErrorsPerSecond, status.SimulatedUsers, status.RequestsExecuted, status.ErrorsRaised, status.AverageExecutionTime, status.Cpu)
		}
	}
}

func OptionIsAgentsAlive() {
	for i := 0; i < len(Config.Agents); i++ {
		result := Config.Agents[i].IsAlive()

		fmt.Printf("Agent %v alive status: %v \n", Config.Agents[i].Name, result)
	}
}

func OptionDisplayConfigration() {
	for i := 0; i < len(Config.Agents); i++ {
		fmt.Println(Config.Agents[i])
	}
}

func OptionStopOne() {

	OptionClearOutput()
	fmt.Println("Enter the Chaos Agent to update:")
	for i := 0; i < len(Config.Agents); i++ {
		fmt.Println(Config.Agents[i].Name)
	}

	reader.Scan()
	name := reader.Text()

	agentFound := false

	for i := 0; i < len(Config.Agents); i++ {
		if Config.Agents[i].Name == name {
			fmt.Println("Stopping test for Agent: " + Config.Agents[i].Name)
			Config.Agents[i].StopTest()
			agentFound = true
		}
	}

	if agentFound == false {
		fmt.Println("Agent not found: " + name)
	}
}

func OptionStopAllTests() {
	for i := 0; i < len(Config.Agents); i++ {
		fmt.Println("Stopping test for Agent: " + Config.Agents[i].Name)
		Config.Agents[i].StopTest()
	}
}
