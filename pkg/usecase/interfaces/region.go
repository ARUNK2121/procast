package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type RegionUsecase interface {
	AddNewState(context.Context, string) error
	GetStates(context.Context) ([]domain.State, error)
	DeleteState(context.Context, int) error
	ReActivateState(ctx context.Context, id int) error
	AddNewDistrict(context.Context, models.AddNewDistrict) error
	GetDistrictsFromState(context.Context, int) ([]domain.District, error)
	DeleteDistrictFromState(context.Context, int) error
}
