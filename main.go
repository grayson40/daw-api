package main

import (
	"log"
	"net/http"

	"github.com/grayson40/daw-api/handlers"
)

func main() {
	// API routes
	http.HandleFunc("/api/v1/users", handlers.UsersHandler)
	http.HandleFunc("/api/v1/projects", handlers.ProjectsHandler)
	http.HandleFunc("/api/v1/issues", handlers.IssuesHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
