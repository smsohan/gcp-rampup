package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Rampup struct {
	Day    time.Time `json:"day"`
	Task   string    `json:"task"`
	Status string    `json:"status"`
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	goRampup := Rampup{
		Day:    time.Now(),
		Task:   "Re-Learn Go",
		Status: "In Progress",
	}
	json.NewEncoder(w).Encode(goRampup)
}

func main() {
	http.HandleFunc("/", ApiHandler)
	http.ListenAndServe(":8080", nil)
}
