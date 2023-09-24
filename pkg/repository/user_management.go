package repository

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
	"gorm.io/gorm"
)

type userManagementRepository struct {
	DB *gorm.DB
}

func NewUserManagementRepository(db *gorm.DB) interfaces.UserManagementRepository {
	return &userManagementRepository{
		DB: db,
	}
}

func (r *userManagementRepository) GetProviders(ctx context.Context) ([]domain.Provider, error) {
	if ctx.Err() != nil {
		return []domain.Provider{}, errors.New("timeout")
	}
	var providers []domain.Provider
	err := r.DB.Raw("SELECT * FROM providers").Scan(&providers).Error
	if err != nil {
		return []domain.Provider{}, err
	}

	return providers, nil
}

func (u *userManagementRepository) MakeProviderVerified(ctx context.Context, id int) error {
	tx := u.DB.Begin()
	err := tx.Exec("UPDATE providers SET verified = true WHERE id = $1", id).Error
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

func (u *userManagementRepository) RevokeVerification(ctx context.Context, id int) error {
	tx := u.DB.Begin()
	err := tx.Exec("UPDATE providers SET verified = false WHERE id = $1", id).Error
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

func (r *userManagementRepository) GetUsers(ctx context.Context) ([]models.UserDetails, error) {
	if ctx.Err() != nil {
		return []models.UserDetails{}, errors.New("timeout")
	}
	var user []models.UserDetails
	err := r.DB.Raw("SELECT * FROM users").Scan(&user).Error
	if err != nil {
		return []models.UserDetails{}, err
	}

	return user, nil
}
