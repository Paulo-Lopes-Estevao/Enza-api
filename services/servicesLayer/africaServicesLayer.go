package serviceslayer

import (
	"io/ioutil"
	"log"
	"os"
)

func GetListAllCountry(dataJSON string) []byte {
	data := loadData(dataJSON)
	return []byte(data)
}

func GetCountryById(dataJSON string) []byte {
	data := loadData(dataJSON)
	return []byte(data)

}

func loadData(dataJSON string) []byte {
	jsonFile, err := os.Open("utils/data/" + dataJSON)

	if err != nil {
		log.Println("Not found File", err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Println("Not Read File", err.Error())
	}

	return data

}
