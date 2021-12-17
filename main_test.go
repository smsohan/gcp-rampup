package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smsohan/gcp-rampup/api"
)

func TestJSONOutput(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Handler)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Got code %v want code %v", status, http.StatusOK)
	}

	var rampup api.Rampup
	json.NewDecoder(response.Body).Decode(&rampup)

	want := "Re-Learn Go"
	if rampup.Task != want {
		t.Errorf("Got %v want %v", rampup.Task, want)
	}

}
