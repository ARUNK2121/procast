package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type WorkUseCase interface {
	GetAllLeads(int, int) ([]models.WorkDetails, error)
	ViewLeads(int) (models.WorkDetails, error)
	PlaceBid(models.PlaceBid) error
	ReplaceBidWithNewBid(models.PlaceBid) error
	GetAllOtherBidsOnTheLeads(work_id int) ([]models.BidDetails, error)
	GetWorksOfAProvider(int) ([]models.WorkDetails, error)
	GetAllOnGoingWorks(int) ([]models.WorkDetails, error)
	GetCompletedWorks(int) ([]models.WorkDetails, error)
}
