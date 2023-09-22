package repository

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

func (cat *categoryRepository) CreateCategory(ctx context.Context, category string) error {
	tx := cat.DB.Begin()
	err := cat.DB.Exec("INSERT INTO categories(category) VALUES($1)", category).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ctx.Err()
	if err != nil {
		tx.Rollback()
		return errors.New("timeout")
	}
	tx.Commit()
	return nil
}

func (cat *categoryRepository) ListCategories(ctx context.Context) ([]domain.Category, error) {
	if ctx.Err() != nil {
		return []domain.Category{}, errors.New("timeout")
	}
	var categories []domain.Category
	err := cat.DB.Raw("SELECT * FROM CATEGORIES").Scan(&categories).Error
	if err != nil {
		return []domain.Category{}, err
	}

	return categories, nil
}

func (cat *categoryRepository) CheckIfCategoryAlreadyExists(ctx context.Context, category string) (bool, error) {
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	err := cat.DB.Raw("SELECT COUNT(*) FROM CATEGORIES WHERE category = $1", category).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count != 0 {
		return false, nil
	}

	return true, nil
}

func (cat *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	tx := cat.DB.Begin()
	err := cat.DB.Exec("UPDATE categories SET is_deleted = TRUE WHERE id = $1", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ctx.Err()
	if err != nil {
		tx.Rollback()
		return errors.New("timeout")
	}
	tx.Commit()
	return nil
}

func (cat *categoryRepository) ReActivateCategory(ctx context.Context, id int) error {
	tx := cat.DB.Begin()
	err := cat.DB.Exec("UPDATE categories SET is_deleted = False WHERE id = $1", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = ctx.Err()
	if err != nil {
		tx.Rollback()
		return errors.New("timeout")
	}
	tx.Commit()
	return nil
}
