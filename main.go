package main

import (
	"github.com/devingen/atama-api/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/build-score-matrix", api.BuildScoreMatrix)
	http.HandleFunc("/build-pairs", api.BuildPairs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
