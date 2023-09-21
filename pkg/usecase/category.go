package usecase

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
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
	exist, err := a.repository.CheckIfCategoryAlreadyExists(ctx, Category)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("category already exists")
	}
	//create new category
	err = a.repository.CreateCategory(ctx, Category)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (a *categoryUsecase) ListCategories(ctx context.Context) ([]domain.Category, error) {

	categories, err := a.repository.ListCategories(ctx)
	if err != nil {
		return []domain.Category{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []domain.Category{}, errors.New("request timeout")
	}

	return categories, nil
}
