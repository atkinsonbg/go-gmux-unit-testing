package handlers

import (
	"net/http"

	"github.com/atkinsonbg/go-gmux-proper-unit-testing/database"
)

func TimezonesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	database.GetAllTimezones()
}
