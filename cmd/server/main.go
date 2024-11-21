package main

import (
	"log"
	"net/http"

	"github.com/raviqlahadi/cv-form-backend/internal/db"
	"github.com/raviqlahadi/cv-form-backend/routes"
)

func main() {
	// Initialize the database connection
	db.ConnectDB()

	// Set up routes
	router := routes.InitRoutes()

	// Start the server
	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
