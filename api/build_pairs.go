package api

import (
	"encoding/json"
	"github.com/devingen/atama-api/atama"
	"github.com/devingen/atama-api/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type BuildPairsBody struct {
	Rules       []model.ConditionalComparisonRule `json:"rules"`
	List1       []atama.MatchItem                 `json:"list1"`
	List1Fields []model.GenericField              `json:"list1Fields"`
	List2       []atama.MatchItem                 `json:"list2"`
	List2Fields []model.GenericField              `json:"list2Fields"`
}

type BuildPairsResponseBody atama.CalculationResult

func BuildPairs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Session-Token, Cache-Control, If-Modified-Since, ETag, X-Requested-With, Client")

	if r.Method == "OPTIONS" {
		return
	}

	var body BuildPairsBody
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(reqBody, &body)
	if err != nil {
		log.Fatal(err.Error())
	}

	start := time.Now()
	log.Printf("")
	log.Printf("%d %d", len(body.List1), len(body.List2))

	scoreMatrix := atama.GenerateScoreMatrix(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	log.Printf("GenerateScoreMatrix took %s", time.Since(start))

	result := atama.CalculateList(scoreMatrix, nil, 0, false)
	log.Printf("CalculateList took %s", time.Since(start))

	_ = json.NewEncoder(w).Encode(result)
}
