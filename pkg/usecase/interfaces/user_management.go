package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
)

type UserManagementUsecase interface {
	GetProviders(context.Context) ([]domain.Provider, error)
}
