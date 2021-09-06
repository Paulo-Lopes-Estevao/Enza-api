package servicesintegration

import (
	"encoding/json"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/apiworld/entities"
	serviceslayer "github.com/Paulo-Lopes-Estevao/apiworld/services/servicesLayer"
	"github.com/gorilla/mux"
)

type Africa struct {
	Africa []entities.Africa
}

func ListAllCountryAfrica(w http.ResponseWriter, r *http.Request) {
	w.Write(serviceslayer.GetListAllCountry("africa.json"))
}

func FindCountryAfrica(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	data := serviceslayer.GetCountryById("africa.json")

	var africa Africa

	json.Unmarshal(data, &africa)

	for _, v := range africa.Africa {

		if v.Name == vars["name"] || v.Keywords[0] == vars["name"] {
			africa, _ := json.Marshal(v)
			w.Write([]byte(africa))
		}

	}
}
