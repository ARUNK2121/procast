package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type ServiceRepository interface {
	CheckIfServiceAlreadyExists(context.Context, string) (bool, error)
	AddServicesToACategory(context.Context, models.AddServicesToACategory) error
	GetServicesInACategory(context.Context, int) ([]domain.Profession, error)
	DeleteService(context.Context, int) error
	ReActivateService(ctx context.Context, id int) error
	GetCommittedWorks(context.Context) ([]domain.Work, error)
	FindServiceFromId(id int) (string, error)
}
