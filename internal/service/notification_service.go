package service

import (
	"notification-service/internal/repository"
)

// NotificationService is the interface for the notification service
type NotificationService interface {
	SendNotification(message string, recipient uint, sender uint) error
}

// NotificationServiceImpl is the implementation of the NotificationService interface
type NotificationServiceImpl struct {
	repo repository.NotificationRepository
}

// NewNotificationService returns a new instance of NotificationServiceImpl
func NewNotificationService(repo repository.NotificationRepository) *NotificationServiceImpl {
	return &NotificationServiceImpl{repo: repo}
}

// SendNotification sends a notification to the specified recipient
func (s *NotificationServiceImpl) SendNotification(message string, recipient uint, sender uint) error {
	// implementation details for sending the notification
	return s.repo.SaveNotification(message, recipient, sender)
}
