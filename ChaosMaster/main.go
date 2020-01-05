package main

import (
	//"mallekoppie/ChaosGenerator/ChaosMaster/manager"
	"mallekoppie/ChaosGenerator/ChaosMaster/routes"

	"net/http"
)

func main() {
	//manager.RunUI()

	router := routes.NewRouter()

	http.ListenAndServe("0.0.0.0:9000", router)
}
