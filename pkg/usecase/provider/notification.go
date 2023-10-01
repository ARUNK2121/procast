package providerusecase

import (
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
