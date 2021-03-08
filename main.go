package main

import (
	"log"
	"net/http"
	"rest_api_test/routes"
)

// Main function
func main() {
	// Init router
	r := routes.NewRouter()

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
