package logic

import (
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
