package interfaces

import "context"

type CategoryUsecase interface {
	CreateCategory(context.Context, string) error
}
