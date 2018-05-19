package manager

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	Prnt("a      = Check if Agents are alive")
	Prnt("start  = Start tests")
	Prnt("stop   = Stop tests")
	Prnt("s      = Status of runners")
	Prnt("w      = Watch status of runners. Update 5s")
	Prnt("deploy = Deploy config to all servers")
	Prnt("d      = Display Configured servers")
	Prnt("c      = Clear")
	Prnt("q      = Quit")
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
		case "start":
		case "stop":
		case "s":
		case "w":
		case "deploy":
		case "d":
			OptionDisplayConfigration()
		case "c":
			OptionClearOutput()
		case "q":
			quit = true
		default:
			Prnt("Command not supported")
		}

	}
}

func OptionDisplayConfigration() {
	for i := 0; i < len(Config.Agents); i++ {
		fmt.Println(Config.Agents[i])
	}
}
