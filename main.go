package main

import (
	"github.com/atkinsonbg/go-gmux-proper-unit-testing/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/healthcheck", handlers.HealthHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", r))
}
