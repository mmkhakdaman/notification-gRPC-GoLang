package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"notification-service/internal/config"
	"notification-service/internal/database"
	"notification-service/internal/handler"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"notification-service/proto"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	serverPort := config.GetEnv("SERVER_PORT")

	database.LoadDB()

	loadGrpcServer(serverPort)
}

func runServer(serverPort string) net.Listener {
	// Create a listener on the specified port
	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}

func loadGrpcServer(serverPort string) {
	lis := runServer(serverPort)

	// Initialize the repository
	repo := repository.NewNotificationRepository()

	// Initialize the service
	svc := service.NewNotificationService(repo)

	// Initialize the handler
	hdl := handler.NewNotificationHandler(svc)

	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register the handler with the server
	proto.RegisterNotificationServiceServer(grpcServer, hdl)

	// Start the server
	log.Printf("starting server on port %s", serverPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
