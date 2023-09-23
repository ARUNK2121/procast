package repository

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
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
	err := r.DB.Exec("INSERT INTO states(state) VALUES($1)", state).Error
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
