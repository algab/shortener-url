package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type message struct {
	Message string
}

// Redirect Route
func Redirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data, result := GetCache(params["id"])
	short, _ := data.(Post)
	if result == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&message{Message: "Link not found"})
	} else {
		http.Redirect(w, r, short.URL, http.StatusMovedPermanently)
	}
}
