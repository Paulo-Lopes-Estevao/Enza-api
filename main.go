package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	servicesintegration "github.com/Paulo-Lopes-Estevao/apiworld/services/servicesIntegration"
	"github.com/gorilla/mux"
)

var Version string = "/v1"

func main() {

	port := os.Getenv("PORT")

	route := mux.NewRouter()

	route.HandleFunc(Version+"/africa", servicesintegration.ListAllCountryAfrica)
	route.HandleFunc(Version+"/africa/{name}", servicesintegration.FindCountryAfrica)

	fmt.Println("Server started at port 2000")
	err := http.ListenAndServe(":"+port, route)

	if err != nil {
		log.Println("Not Running Server...", err.Error())
	}

}
