package repository

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
	"gorm.io/gorm"
)

type regionrepository struct {
	DB *gorm.DB
}

func NewRegionRepository(db *gorm.DB) interfaces.RegionRepository {
	return &regionrepository{
		DB: db,
	}
}

func (r *regionrepository) AddNewState(ctx context.Context, state string) error {
	tx := r.DB.Begin()
	err := tx.Exec("INSERT INTO states(state) VALUES($1)", state).Error
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

func (r *regionrepository) CheckIfStateAlreadyExists(ctx context.Context, state string) (bool, error) {
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	err := r.DB.Raw("SELECT COUNT(*) FROM states WHERE state = $1", state).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count != 0 {
		return false, nil
	}

	return true, nil
}

func (r *regionrepository) GetStates(ctx context.Context) ([]domain.State, error) {
	if ctx.Err() != nil {
		return []domain.State{}, errors.New("timeout")
	}
	var states []domain.State
	err := r.DB.Raw("SELECT * FROM states").Scan(&states).Error
	if err != nil {
		return []domain.State{}, err
	}

	return states, nil
}

func (r *regionrepository) DeleteState(ctx context.Context, id int) error {
	tx := r.DB.Begin()
	err := tx.Exec("UPDATE states SET is_deleted = TRUE WHERE id = $1", id).Error
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

func (r *regionrepository) ReActivateState(ctx context.Context, id int) error {
	tx := r.DB.Begin()
	err := tx.Exec("UPDATE states SET is_deleted = False WHERE id = $1", id).Error
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

func (r *regionrepository) CheckIfDistrictAlreadyExists(ctx context.Context, district string) (bool, error) {
	if ctx.Err() != nil {
		return false, errors.New("timeout")
	}
	var count int64
	err := r.DB.Raw("SELECT COUNT(*) FROM districts WHERE district = $1", district).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count != 0 {
		return false, nil
	}

	return true, nil
}

func (r *regionrepository) AddNewDistrict(ctx context.Context, district models.AddNewDistrict) error {
	tx := r.DB.Begin()
	err := tx.Exec("INSERT INTO districts(district,state_id) VALUES($1,$2)", district.District, district.StateID).Error
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

func (r *regionrepository) GetDistrictsFromState(ctx context.Context, id int) ([]domain.District, error) {
	if ctx.Err() != nil {
		return []domain.District{}, errors.New("timeout")
	}
	var districts []domain.District
	err := r.DB.Raw("SELECT * FROM districts WHERE state_id = $1", id).Scan(&districts).Error
	if err != nil {
		return []domain.District{}, err
	}

	return districts, nil
}

func (r *regionrepository) DeleteDistrictFromState(ctx context.Context, id int) error {
	tx := r.DB.Begin()
	err := tx.Exec("UPDATE districts SET is_deleted = TRUE WHERE id = $1", id).Error
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
