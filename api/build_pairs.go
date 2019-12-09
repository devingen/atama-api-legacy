package api

import (
	"encoding/json"
	"github.com/devingen/atama-api/dto"
	"github.com/devingen/atama-api/service"
	"io/ioutil"
	"log"
	"net/http"
)

func BuildPairs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Session-Token, Cache-Control, If-Modified-Since, ETag, X-Requested-With, Client")

	if r.Method == "OPTIONS" {
		return
	}

	var body dto.BuildPairsBody
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		log.Fatal(err.Error())
	}

	_ = json.NewEncoder(w).Encode(service.BuildPairs(body))

}
