package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Rampup struct {
	Time   time.Time `json:"time"`
	Task   string    `json:"task"`
	Status string    `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	runtime, isProvided := os.LookupEnv("RUNTIME")
	if !isProvided {
		runtime = "Localhost"
	}

	goRampup := Rampup{
		Time:   time.Now(),
		Task:   "Re-Learn Go",
		Status: fmt.Sprintf("Deployed on %s", runtime),
	}
	json.NewEncoder(w).Encode(goRampup)
}
