package interfaces

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type WorkUsecase interface {
	ListNewOpening(domain.Work) error
	GetAllListedWorks(id int) ([]models.WorkDetails, error)
	ListAllCompletedWorks(id int) ([]models.WorkDetails, error)
	ListAllOngoingWorks(id int) ([]models.WorkDetails, error)
	WorkDetails(id int) (models.WorkDetails, error)
	AssignWorkToProvider(work_id, pro_id int) error
}
