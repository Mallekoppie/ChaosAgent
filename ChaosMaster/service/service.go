package service

import (
	"encoding/json"
	"io/ioutil"
	"mallekoppie/ChaosGenerator/ChaosMaster/logic"
	"mallekoppie/ChaosGenerator/ChaosMaster/models"
	"mallekoppie/ChaosGenerator/ChaosMaster/util/logger"
	"net/http"

	"github.com/gorilla/mux"
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
	testGroups, err := logic.GetAllTestGroups()
	if err != nil {
		logger.Error("Unable to re")
	}

	data, err := json.Marshal(testGroups)
	if err != nil {
		logger.Error("Error while marshalling test groups: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func AddTestGroup(w http.ResponseWriter, r *http.Request) {
	group := models.TestGroup{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Unable to read body: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Info("Request body received: ", string(data))

	err = json.Unmarshal(data, &group)
	if err != nil {
		logger.Error("AddTestGroup: Unable to unmarshal request: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = logic.CreateTestGroup(group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func UpdateTestGroup(w http.ResponseWriter, r *http.Request) {

}

func DeleteTestGroup(w http.ResponseWriter, r *http.Request) {

}

func AddTestCollection(w http.ResponseWriter, r *http.Request) {

}

func UpdateTestCollection(w http.ResponseWriter, r *http.Request) {

}

func DeleteTestCollection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if len(id) < 1 {
		logger.Error("Invalid ID sent in path url")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Info("Id Received: ", id)

	err := logic.DeleteTestCollection(id)
	if err != nil {
		logger.Error("Unable to delete test collection: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func templateToCopy(w http.ResponseWriter, r *http.Request) {

}
