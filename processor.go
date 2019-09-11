package main

import (
	"encoding/json"
	"net/http"

	jq "gopkg.in/thedevsaddam/gojsonq.v2"
)

const (
	selectQueryHeaderParam = "Select-query"
)

func findHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	jqSelectPath := r.Header.Get(selectQueryHeaderParam)

	items := jq.New().Reader(r.Body).Pluck(jqSelectPath)
	logger.Printf("Select selector: '%s' found: '%v'", jqSelectPath, items)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		logger.Printf("Error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func selectHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	jqSelectPath := r.Header.Get(selectQueryHeaderParam)

	items := jq.New().Reader(r.Body).Find(jqSelectPath)
	logger.Printf("Selector %s found %+v", jqSelectPath, items)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		logger.Printf("Error encoding response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
