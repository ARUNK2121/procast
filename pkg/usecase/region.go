package usecase

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/interfaces"
)

type regionUsecase struct {
	repository interfaces.RegionRepository
}

func NewRegionUsecase(repo interfaces.RegionRepository) services.RegionUsecase {
	return &regionUsecase{
		repository: repo,
	}
}

func (a *regionUsecase) AddNewState(ctx context.Context, Category string) error {
	//check if already a category exists in same name
	exist, err := a.repository.CheckIfStateAlreadyExists(ctx, Category)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("category already exists")
	}
	//create new category
	err = a.repository.AddNewState(ctx, Category)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (a *regionUsecase) GetStates(ctx context.Context) ([]domain.State, error) {

	states, err := a.repository.GetStates(ctx)
	if err != nil {
		return []domain.State{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []domain.State{}, errors.New("request timeout")
	}

	return states, nil
}

func (r *regionUsecase) DeleteState(ctx context.Context, id int) error {

	err := r.repository.DeleteState(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (a *regionUsecase) ReActivateState(ctx context.Context, id int) error {

	err := a.repository.ReActivateState(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}
