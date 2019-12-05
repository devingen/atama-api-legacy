package main

import (
	"github.com/devingen/atama-api/api"
	"log"
	"net/http"
)

const PORT = "8080"

func main() {
	log.Println("Server is running on ", PORT)
	http.HandleFunc("/build-score-matrix", api.BuildScoreMatrix)
	http.HandleFunc("/build-pairs", api.BuildPairs)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
