package interfaces

import "github.com/ARUNK2121/procast/pkg/domain"

type NotificationRepository interface {
	GetAllNotifications(provider_id int) ([]domain.ProviderNotification, error)
	MakeNotificationAsRead(notification_id int) error
	ViewNotification(notification_id int) (domain.ProviderNotification, error)
}
