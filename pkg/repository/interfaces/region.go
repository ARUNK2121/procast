package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
)

type RegionRepository interface {
	AddNewState(context.Context, string) error
	CheckIfStateAlreadyExists(context.Context, string) (bool, error)
	GetStates(context.Context) ([]domain.State, error)
}
