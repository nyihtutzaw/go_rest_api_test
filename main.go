package main

import (
	"log"
	"net/http"
	"rest_api_test/config"
	"rest_api_test/database"
	"rest_api_test/routes"
)

// Main function
func main() {
	// Init router
	r := routes.NewRouter()
	s := http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources/")))
	r.PathPrefix("/").Handler(s)

	// migrate db
	database.MigrateTables()

	// Start server
	log.Fatal(http.ListenAndServe(":"+config.GETEnvVariable("PORT"), r))
}
