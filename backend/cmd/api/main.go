package main

import (
	"backend/internal/db"
	"backend/internal/store"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get MONGO_URI from environment
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in environment")
	}

	if err := db.InitDB(mongoURI); err != nil {
		log.Println(err)
	}

	defer db.CloseDB()

	mongoClient := db.GetMongoClient()
	store := store.NewStorage(mongoClient, "ems-brevo")

	fmt.Println("Mongo client initialized")

	config := config{
		addr: ":8080",
	}

	app := &application{
		Config: config,
		Store:  store,
	}

	fmt.Println("Server started")
	app.run(app.mount())
}
