package services

import (
	"encoding/json"
	"gin-go-api/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GetObjectFromWebService() models.Object {

	response, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject models.Object
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
