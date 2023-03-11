package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"notification-service/internal/service"
	"notification-service/proto"
)

// NotificationHandler handles incoming gRPC requests for the NotificationService
type NotificationHandler struct {
	svc service.NotificationService
}

// NewNotificationHandler returns a new instance of NotificationHandler
func NewNotificationHandler(svc service.NotificationService) *NotificationHandler {
	return &NotificationHandler{svc: svc}
}

// SendNotification sends a notification to the specified recipient
func (h *NotificationHandler) SendNotification(ctx context.Context, req *proto.SendNotificationRequest) (*proto.SendNotificationResponse, error) {
	err := h.svc.SendNotification(req.Message, uint(req.Recipient), uint(req.Sender))
	if err != nil {
		// handle error
		return nil, status.Errorf(codes.Internal, "failed to send notification: %v", err)
	}

	return &proto.SendNotificationResponse{}, nil
}
