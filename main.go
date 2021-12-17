package main

import (
	"net/http"

	"github.com/smsohan/gcp-rampup/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	http.ListenAndServe(":8080", nil)
}
