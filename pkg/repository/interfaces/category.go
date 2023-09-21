package interfaces

import "context"

type CategoryRepository interface {
	CreateCategory(context.Context, string) error
}
