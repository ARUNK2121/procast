package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type ServiceUsecase interface {
	AddServicesToACategory(context.Context, models.AddServicesToACategory) error
	GetServicesInACategory(context.Context, int) ([]domain.Profession, error)
	DeleteService(context.Context, int) error
	ReActivateService(ctx context.Context, id int) error
	ListCommittedWorks(context.Context) ([]models.WorkDetails, error)
}
