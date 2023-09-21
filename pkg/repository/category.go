package repository

import (
	"context"
	"errors"

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
