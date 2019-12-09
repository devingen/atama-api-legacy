package main

import (
	"github.com/devingen/atama-api/api"
	"log"
	"net/http"
	"net/http/pprof"
)

const PORT = "8080"

func main() {
	log.Println("Server is running on ", PORT)
	http.HandleFunc("/build-score-matrix", api.BuildScoreMatrix)
	http.HandleFunc("/build-pairs", api.BuildPairs)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))

	// Register pprof handlers
	http.HandleFunc("/debug/pprof/", pprof.Index)
	http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	http.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	http.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	http.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	http.Handle("/debug/pprof/block", pprof.Handler("block"))
}
