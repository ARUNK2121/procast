package usecase

import (
	"context"

	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/interfaces"
)

type categoryUsecase struct {
	repository interfaces.CategoryRepository
}

func NewCategoryUsecase(repo interfaces.CategoryRepository) services.CategoryUsecase {
	return &categoryUsecase{
		repository: repo,
	}
}

func (a *categoryUsecase) CreateCategory(ctx context.Context, Category string) error {
	//check if already a category exists in same name

	//create new category
	err := a.repository.CreateCategory(ctx, Category)
	if err != nil {
		return err
	}

	return nil
}
