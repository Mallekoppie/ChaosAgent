package logic

import (
	"mallekoppie/ChaosGenerator/ChaosMaster/manager"
	"mallekoppie/ChaosGenerator/ChaosMaster/models"
	"mallekoppie/ChaosGenerator/ChaosMaster/repositories"
	"mallekoppie/ChaosGenerator/ChaosMaster/util/logger"
)

const (
	consulAgentServiceName string = "ChaosAgent"
	consulAgentMetricsName string = "ChaosAgentMetrics"
)

func GetAllAgents() (agents []models.Agent, err error) {
	logger.Info("Getting all agents")

	agents, err = repositories.GetAllAgents(consulAgentServiceName)
	if err != nil {
		logger.Error("Unable to get agents: ", err.Error())
		return agents, err
	}

	agentmetrics, err := repositories.GetAllAgents(consulAgentMetricsName)
	if err != nil {
		logger.Error("Unable to get agents: ", err.Error())
		return agents, err
	}

	// Add metrics port to agent
	for index := range agents {
		for metricIndex := range agentmetrics {
			if agents[index].Id == agentmetrics[metricIndex].Id {
				agents[index].MetricsPort = agentmetrics[metricIndex].Port
			}
		}
	}

	// Get Agent status
	for index := range agents {
		agent := agents[index]

		chaosAgent, err := manager.GetAgent(agent.Id)
		if err != nil {
			logger.Error("Unable to get agent for Id: ", agent.Id)
			agents[index].Status = "error"
			continue
		}

		alive := chaosAgent.IsAlive()

		logger.Info("Agent IsAlive response: ", alive)

		if alive == true {
			logger.Info("Setting online")
			agents[index].Status = "online"
		} else {
			logger.Info("Setting offline")
			agents[index].Status = "offline"
		}
	}

	logger.Info("Returning agents successfully")
	return agents, nil
}

func UpdateAgent(agent models.Agent) error {
	logger.Info("Adding agent")
	// Register Agent
	err := repositories.UpdateChaosAgent(agent, consulAgentServiceName, agent.Port)
	if err != nil {
		logger.Error("Unable to register normal agent in consul: ", err.Error())
		return err
	}
	// Register Metrics
	err = repositories.UpdateChaosAgent(agent, consulAgentMetricsName, agent.MetricsPort)
	if err != nil {
		logger.Error("Unable to register metrics agent in consul: ", err.Error())
		return err
	}

	logger.Info("Added agent successfully")
	return nil
}

func DeleteAgent(agent models.Agent) error {
	logger.Info("Deleteing agent")

	err := repositories.DeleteChaosAgent(agent, consulAgentServiceName)
	if err != nil {
		logger.Error("Unable to delete normal agent in consul: ", err.Error())
		return err
	}

	err = repositories.DeleteChaosAgent(agent, consulAgentMetricsName)
	if err != nil {
		logger.Error("Unable to delete metric agent in consul: ", err.Error())
		return err
	}

	logger.Info("Agent deleted")
	return nil

}

func GetAllTestGroups() (tests []models.TestGroup, err error) {
	tests, err = repositories.GetAllTestGroups()
	if err != nil {
		return tests, err
	}

	logger.Info("Tests Returned: ", tests)

	return tests, nil
}

func GetTestGroup(id string) (group models.TestGroup, err error) {
	group, err = repositories.GetTestGroup(id)
	if err != nil {
		logger.Error("Unable to get Test Group: ", err.Error())
		return group, err
	}

	return group, err
}

func CreateTestGroup(group models.TestGroup) error {
	err := repositories.AddTestGroup(group)
	if err != nil {
		logger.Error("Unable to create new TestGroup: ", err)
		return err
	}

	return nil
}

func UpdateTestGroup(group models.TestGroup) error {
	err := repositories.UpdateTestGroup(group)
	return err
}

func DeleteTestGroup(id string) error {
	err := repositories.DeleteTestGroup(id)
	return err
}

func AddTestCollection(test models.TestCollection) error {
	err := repositories.AddTestCollection(test)

	return err
}

func UpdateTestCollection(test models.TestCollection) error {
	err := repositories.UpdateTestCollection(test)

	return err
}

func DeleteTestCollection(id string) error {
	err := repositories.DeleteTestCollection(id)
	return err
}
