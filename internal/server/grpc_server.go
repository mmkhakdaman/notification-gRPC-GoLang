package server

import (
	"google.golang.org/grpc"
	"net"

	"notification-service/internal/config"
	"notification-service/internal/handler"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"notification-service/proto"
)

// GRPCServer represents a gRPC server for the NotificationService
type GRPCServer struct {
	server *grpc.Server
}

// NewGRPCServer returns a new instance of GRPCServer
func NewGRPCServer() *GRPCServer {
	return &GRPCServer{server: grpc.NewServer()}
}

// Start starts the gRPC server and listens for incoming requests
func (s *GRPCServer) Start() error {
	serverPort := config.GetEnv("SERVER_PORT")
	// Create a listener on the specified port
	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		return err
	}

	// Initialize the repository
	repo := repository.NewNotificationRepository()

	// Initialize the service
	svc := service.NewNotificationService(repo)

	// Initialize the handler
	hdl := handler.NewNotificationHandler(svc)

	// Register the handler with the server
	proto.RegisterNotificationServiceServer(s.server, hdl)

	// Start the server
	return s.server.Serve(lis)
}

// Stop stops the gRPC server and closes all connections
func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}
