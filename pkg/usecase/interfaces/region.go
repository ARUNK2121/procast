package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
)

type RegionUsecase interface {
	AddNewState(context.Context, string) error
	GetStates(context.Context) ([]domain.State, error)
	DeleteState(context.Context, int) error
}
