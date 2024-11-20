package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Define the static files directory
	staticDir := "./templates"

	// Check if the directory exists
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Fatalf("Static directory %s does not exist", staticDir)
	}

	// Serve static files
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	// Start the server
	port := "8080"
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
