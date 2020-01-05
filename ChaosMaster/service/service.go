package service

import (
	"encoding/json"
	"io/ioutil"
	"mallekoppie/ChaosGenerator/ChaosMaster/logic"
	"mallekoppie/ChaosGenerator/ChaosMaster/models"
	"mallekoppie/ChaosGenerator/ChaosMaster/util/logger"
	"net/http"
)

func GetAllAgents(w http.ResponseWriter, r *http.Request) {
	agents, err := logic.GetAllAgents()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(agents)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func DeleteAgent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Unable to read request body: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	agent := models.Agent{}

	err = json.Unmarshal(data, &agent)
	if err != nil {
		logger.Error("Unable to unmarshal request: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = logic.DeleteAgent(agent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateAgent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Unable to read request body: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	agent := models.Agent{}

	err = json.Unmarshal(data, &agent)
	if err != nil {
		logger.Error("Unable to unmarshal request: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = logic.UpdateAgent(agent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetAllTestGroups(w http.ResponseWriter, r *http.Request) {

}

func AddTestGroup(w http.ResponseWriter, r *http.Request) {

}

func DeleteTestGroup(w http.ResponseWriter, r *http.Request) {

}

func templateToCopy(w http.ResponseWriter, r *http.Request) {

}
