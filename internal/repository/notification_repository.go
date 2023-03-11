package repository

import (
	"notification-service/internal/database"
	"notification-service/internal/models"
)

// NotificationRepository is the interface for the notification repository
type NotificationRepository interface {
	SaveNotification(message string, recipient uint, sender uint) error
}

// NotificationRepositoryImpl is the implementation of the NotificationRepository interface
type NotificationRepositoryImpl struct {
	// implementation details for the repository, such as the database client
}

// NewNotificationRepository returns a new instance of NotificationRepositoryImpl
func NewNotificationRepository() *NotificationRepositoryImpl {
	return &NotificationRepositoryImpl{}
}

// SaveNotification saves the notification to the repository
func (r *NotificationRepositoryImpl) SaveNotification(message string, recipient uint, sender uint) error {
	// implementation details for saving the notification, such as interacting with a database

	db := database.GetDB()

	notification := &models.Notification{
		Content:  message,
		SenderId: sender,
		UserId:   recipient,
	}

	return db.Create(notification).Error
}
