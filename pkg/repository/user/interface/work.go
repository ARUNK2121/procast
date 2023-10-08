package interfaces

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type WorkRepository interface {
	ListNewOpening(domain.Work) error
	FindUsername(id int) (string, error)
	GetAllWorksOfAUser(id int) ([]int, error)
	GetAllCompletedWorksOfAUser(id int) ([]int, error)
	GetAllOngoingWorksOfAUser(id int) ([]int, error)
	GetDetailsOfAWork(int) (models.MinWorkDetails, error)
	GetImagesOfAWork(int) ([]string, error)
	FindProviderIdFromWork(int) (int, error)
	FindProviderName(pro_id int) (string, error)
}
