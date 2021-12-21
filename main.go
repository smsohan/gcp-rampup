package main

import (
	"fmt"
	"net/http"

	"github.com/smsohan/gcp-rampup/api"
)

func main() {
	fmt.Printf("Starting to listen at :8080")
	http.HandleFunc("/", api.Handler)
	http.ListenAndServe(":8080", nil)
}
