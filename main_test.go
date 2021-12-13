package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func TestJSONOutput(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(ApiHandler)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Got code %v want code %v", status, http.StatusOK)
	}

	var rampup Rampup
	json.NewDecoder(response.Body).Decode(&rampup)

	want := "Re-Learn Go"
	if rampup.Task != want {
		t.Errorf("Got %v want %v", rampup.Task, want)
	}

}