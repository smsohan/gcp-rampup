package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Rampup struct {
	Time    time.Time `json:"time"`
	Task   string    `json:"task"`
	Status string    `json:"status"`
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	goRampup := Rampup{
		Time:    time.Now(),
		Task:   "Re-Learn Go",
		Status: "Deployed on CloudRun",
	}
	json.NewEncoder(w).Encode(goRampup)
}

func main() {
	http.HandleFunc("/", ApiHandler)
	http.ListenAndServe(":8080", nil)
}
