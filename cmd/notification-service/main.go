package main

import (
	"log"
	"notification-service/internal/config"
	"notification-service/internal/database"
	"notification-service/internal/server"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	lis := server.NewGRPCServer()

	database.LoadDB()

	// Start the server
	if err := lis.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
