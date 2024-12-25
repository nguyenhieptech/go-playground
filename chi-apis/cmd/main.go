package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nguyenhieptech/go-playground/chi-apis/internal/db"
	"github.com/nguyenhieptech/go-playground/chi-apis/internal/routes"
)

func main() {
	// Load .env file
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	// Initialize database connection
	mongoClient, collection := db.InitMongoDB()
	defer mongoClient.Disconnect(db.Ctx)

	// Set up routes
	r := routes.SetupRoutes(collection)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
