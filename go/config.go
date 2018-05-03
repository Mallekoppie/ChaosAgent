package swagger

import (
	json "encoding/json"
	io "io/ioutil"
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

func WriteTestConfiguration(config TestCollection) error {
	data, err := json.Marshal(config)

	if err != nil {
		log.Println("Failed to marshall config: %v", err)
		return err
	}

	name := config.Name + ".json"

	err = io.WriteFile(name, data, os.ModeExclusive)

	if err != nil {
		log.Println("Failed to marshall config: %v", err)
		return err
	}

	return nil
}

func ReadTestConfiguration(name string) (TestCollection, error) {
	configuration := TestCollection{}
	err := gonfig.GetConf(name+".json", &configuration)

	if err != nil {
		log.Print("Error reading config: %v", err)
		return configuration, err
	}

	return configuration, nil
}
