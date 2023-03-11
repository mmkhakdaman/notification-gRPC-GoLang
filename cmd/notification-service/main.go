package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"notification-service/internal/config"
	"notification-service/internal/handler"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"notification-service/proto"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Create a listener on the specified port
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

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
	log.Printf("starting server on port %s", cfg.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
