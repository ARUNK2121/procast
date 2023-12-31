package adminrepository

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/repository/admin/interfaces"
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

func (r *userManagementRepository) GetProviders(ctx context.Context) ([]models.ProviderDetails, error) {
	if ctx.Err() != nil {
		return []models.ProviderDetails{}, errors.New("timeout")
	}
	var providers []models.ProviderDetails
	err := r.DB.Raw("SELECT * FROM providers").Scan(&providers).Error
	if err != nil {
		return []models.ProviderDetails{}, err
	}

	return providers, nil
}

func (u *userManagementRepository) MakeProviderVerified(ctx context.Context, id int) error {
	tx := u.DB.Begin()
	err := tx.Exec("UPDATE providers SET is_verified = true WHERE id = $1", id).Error
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
	err := tx.Exec("UPDATE providers SET is_verified = false WHERE id = $1", id).Error
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

func (u *userManagementRepository) BlockUser(ctx context.Context, id int) error {
	tx := u.DB.Begin()
	err := tx.Exec("UPDATE users SET is_blocked = true WHERE id = $1", id).Error
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

func (u *userManagementRepository) UnBlockUser(ctx context.Context, id int) error {
	tx := u.DB.Begin()
	err := tx.Exec("UPDATE users SET is_blocked = false WHERE id = $1", id).Error
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

func (s *userManagementRepository) FindUserFromId(id int) (string, error) {
	var name string
	err := s.DB.Raw("SELECT name FROM users WHERE id = $1", id).Scan(&name).Error
	if err != nil {
		return "", err
	}

	return name, nil
}

func (s *userManagementRepository) FindProviderFromId(id int) (string, error) {
	var name string
	err := s.DB.Raw("SELECT name FROM providers WHERE id = $1", id).Scan(&name).Error
	if err != nil {
		return "", err
	}

	return name, nil
}

func (r *userManagementRepository) GetAllPendingVerifications(ctx context.Context) ([]models.Verification, error) {
	if ctx.Err() != nil {
		return []models.Verification{}, errors.New("timeout")
	}
	var verifications []models.Verification
	err := r.DB.Raw("SELECT id,name FROM providers WHERE is_verified = false").Scan(&verifications).Error
	if err != nil {
		return []models.Verification{}, err
	}

	return verifications, nil
}

func (r *userManagementRepository) FindDocumentsOfProviderFromID(id int) (string, error) {

	var document string
	err := r.DB.Raw("SELECT id_proof FROM id_proofs WHERE pro_id = $1", id).Scan(&document).Error
	if err != nil {
		return "", err
	}

	return document, nil
}
