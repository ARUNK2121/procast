package interfaces

import "github.com/ARUNK2121/procast/pkg/domain"

type NotificationUsecase interface {
	GetAllNotifications(provider_id int) ([]domain.ProviderNotification, error)
}
