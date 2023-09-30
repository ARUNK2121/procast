package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type WorkUseCase interface {
	GetAllLeads(int, int) ([]models.WorkDetails, error)
	ViewLeads(int) (models.WorkDetails, error)
	PlaceBid(models.PlaceBid) error
}
