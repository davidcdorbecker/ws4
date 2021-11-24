package controllers

import (
	"encoding/json"
	"net/http"

	"ws4/usecases"

	"github.com/gorilla/mux"
)

var AllowedFilters = map[string]bool {
	"species": true,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filter := params["filter"]
	value := params["toSearch"]

	if _, ok := AllowedFilters[filter]; !ok {
		bytes, _ := json.Marshal(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{http.StatusBadRequest, "filter not allowed"})

		w.Header().Add("Content-Type", "application/json")
		w.Write(bytes)
		return
	}

	characters, err := usecases.FilterCharacter(filter, value)
	if err != nil {
		bytes, _ := json.Marshal(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{http.StatusInternalServerError, err.Error()})

		w.Header().Add("Content-Type", "application/json")
		w.Write(bytes)
		return
	}

	bytes, _ := json.Marshal(characters)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}
