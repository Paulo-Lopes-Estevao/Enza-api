package main

import (
	"fmt"
	"log"
	"net/http"

	servicesintegration "github.com/Paulo-Lopes-Estevao/apiworld/services/servicesIntegration"
	"github.com/gorilla/mux"
)

var Version string = "/v1"

func main() {

	route := mux.NewRouter()

	route.HandleFunc(Version+"/africa", servicesintegration.ListAllCountryAfrica)
	route.HandleFunc(Version+"/africa/{name}", servicesintegration.FindCountryAfrica)

	fmt.Println("Server started at port 2000")
	err := http.ListenAndServe(":2000", route)

	if err != nil {
		log.Println("Not Running Server...", err.Error())
	}

}
