package main

import (
	"net/http"
	"time"
	"encoding/json"
)


type Rampup struct {
	Day time.Time `json:"day"`
	Task string `json:"task"`
	Status string `json:"status"`
}

func main() {
	goRampup := Rampup{
		Day: time.Now(),
		Task: "Re-Learn Go",
		Status: "In Progress",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(goRampup)
	})

	http.ListenAndServe(":8080", nil)
}
