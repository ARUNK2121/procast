package providerhandler

import interfaces "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"

type notificationHandler struct {
	usecase interfaces.NotificationUsecase
}

func NewNotificationUsecase(use interfaces.NotificationUsecase) *notificationHandler {
	return &notificationHandler{
		usecase: use,
	}
}
