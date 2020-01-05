package logic

import (
	"mallekoppie/ChaosGenerator/ChaosMaster/repositories"
	"mallekoppie/ChaosGenerator/ChaosMaster/util/logger"
)

const (
	ConsulAgentServiceName string = "ChaosAgent"
)

func GetAllAgents() {
	logger.Info("Getting all agents")
	repositories.GetAllAgents(ConsulAgentServiceName)
}
