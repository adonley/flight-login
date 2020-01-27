package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// set the json response
func jsonResponse(w http.ResponseWriter, data interface{}, c int) {
	dj, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

// Health returns 200 when service is healthy
func Health(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["status"] = "up"
	jsonResponse(w, response, http.StatusOK)
}
