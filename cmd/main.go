package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/angelotinho/ldsdemo/api"
)

const (
	// Port ...
	Port = ":8080"
)

func main() {
	log.Println("Starting Leidos Assignment server...")
	log.Printf("Go version %s", runtime.Version())
	mux := http.NewServeMux()

	appServer := api.NewServer()
	mux.Handle("/api/hello", appServer)
	mux.Handle("/api/add", appServer)
	mux.Handle("/api/time", appServer)
	err := http.ListenAndServe(Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
