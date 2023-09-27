package interfaces

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/domain"
)

type CategoryUsecase interface {
	CreateCategory(context.Context, string) error
	ListCategories(context.Context) ([]domain.Category, error)
	DeleteCategory(context.Context, int) error
	ReActivateCategory(ctx context.Context, id int) error
}
