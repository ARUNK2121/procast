package providerusecase

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
)

type notificationUsecase struct {
	repository interfaces.NotificationRepository
}

func NewNotificationUsecase(repo interfaces.NotificationRepository) services.NotificationUsecase {
	return &notificationUsecase{
		repository: repo,
	}

}

func (u *notificationUsecase) GetAllNotifications(provider_id int) ([]domain.ProviderNotification, error) {
	notifications, err := u.repository.GetAllNotifications(provider_id)
	if err != nil {
		return []domain.ProviderNotification{}, err
	}

	return notifications, nil
}
