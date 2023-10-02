package providerRepository

import (
	"github.com/ARUNK2121/procast/pkg/domain"
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

func (n *notificationRepository) GetAllNotifications(pro_id int) ([]domain.ProviderNotification, error) {
	var notification []domain.ProviderNotification
	if err := n.DB.Raw(`SELECT * FROM provider_notifications WHERE pro_id = $1 ORDER BY time DESC`, pro_id).Scan(&notification).Error; err != nil {
		return []domain.ProviderNotification{}, err
	}

	return notification, nil
}
