package main

import (
	"github.com/atkinsonbg/go-gmux-proper-unit-testing/handlers"
	"github.com/atkinsonbg/go-gmux-proper-unit-testing/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/healthcheck", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/timezones", handlers.TimezonesHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", r))
}