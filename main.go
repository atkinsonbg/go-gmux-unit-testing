package main

import (
	"github.com/atkinsonbg/go-gmux-proper-unit-testing/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	http.ListenAndServe(":80", r)
}
