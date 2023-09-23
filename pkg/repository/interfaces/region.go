package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type RegionRepository interface {
	AddNewState(context.Context, string) error
	CheckIfStateAlreadyExists(context.Context, string) (bool, error)
	GetStates(context.Context) ([]domain.State, error)
	DeleteState(context.Context, int) error
	ReActivateState(ctx context.Context, id int) error
	AddNewDistrict(context.Context, models.AddNewDistrict) error
	CheckIfDistrictAlreadyExists(context.Context, string) (bool, error)
	GetDistrictsFromState(context.Context, int) ([]domain.District, error)
	DeleteDistrictFromState(context.Context, int) error
}
