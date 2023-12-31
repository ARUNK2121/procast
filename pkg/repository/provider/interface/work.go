package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type WorkRepository interface {
	GetLeadByServiceAndLocation(service, location int) ([]int, error)
	GetDetailsOfAWork(int) (models.MinWorkDetails, error)
	GetImagesOfAWork(int) ([]string, error)
	PlaceBid(model models.PlaceBid) error
	ReplaceBidWithNewBid(model models.PlaceBid) error
	GetAllOtherBidsOnTheLeads(work_id int) ([]models.BidDetails, error)
	CheckIfAlreadyBidded(work_id, pro_id int) (bool, error)
	GetAllWorksOfAProvider(pro_id int) ([]int, error)
	FindProviderName(pro_id int) (string, error)
	GetCommittedWorksOfAProvider(pro_id int) ([]int, error)
	GetCompletedWorksOfAProvider(pro_id int) ([]int, error)
}
