package providerRepository

import (
	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
	"gorm.io/gorm"
)

type notificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) interfaces.NotificationRepository {
	return &notificationRepository{
		DB: db,
	}
}
