package repository

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
	"gorm.io/gorm"
)

type serviceRepository struct {
	DB *gorm.DB
}

func NewServiceRepository(db *gorm.DB) interfaces.ServiceRepository {
	return &serviceRepository{
		DB: db,
	}
}

func (s *serviceRepository) CheckIfServiceAlreadyExists(ctx context.Context, service string) (bool, error) {
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	err := s.DB.Raw("SELECT COUNT(*) FROM professions WHERE profession = $1", service).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count != 0 {
		return false, nil
	}

	return true, nil
}

func (s *serviceRepository) AddServicesToACategory(ctx context.Context, service models.AddServicesToACategory) error {
	tx := s.DB.Begin()
	err := tx.Exec("INSERT INTO professions(profession,category_id) VALUES($1,$2)", service.ServiceName, service.CategoryID).Error
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

func (s *serviceRepository) GetServicesInACategory(ctx context.Context, id int) ([]domain.Profession, error) {
	if ctx.Err() != nil {
		return []domain.Profession{}, errors.New("timeout")
	}
	var services []domain.Profession
	err := s.DB.Raw("SELECT * FROM professions WHERE category_id = $1", id).Scan(&services).Error
	if err != nil {
		return []domain.Profession{}, err
	}

	return services, nil
}

func (s *serviceRepository) DeleteService(ctx context.Context, id int) error {
	tx := s.DB.Begin()
	err := tx.Exec("UPDATE professions SET is_deleted = TRUE WHERE id = $1", id).Error
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

func (s *serviceRepository) ReActivateService(ctx context.Context, id int) error {
	tx := s.DB.Begin()
	err := tx.Exec("UPDATE professions SET is_deleted = False WHERE id = $1", id).Error
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
