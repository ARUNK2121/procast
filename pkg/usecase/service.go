package usecase

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type serviceUsecase struct {
	repository interfaces.ServiceRepository
}

func NewServiceUsecase(repo interfaces.ServiceRepository) services.ServiceUsecase {
	return &serviceUsecase{
		repository: repo,
	}
}

func (a *serviceUsecase) AddServicesToACategory(ctx context.Context, service models.AddServicesToACategory) error {
	if err := ctx.Err(); err != nil {
		return errors.New("request timeout")
	}
	//check if already a category exists in same name
	exist, err := a.repository.CheckIfServiceAlreadyExists(ctx, service.ServiceName)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("service already exists")
	}
	//create new category
	err = a.repository.AddServicesToACategory(ctx, service)
	if err != nil {
		return err
	}

	return nil
}

func (a *serviceUsecase) GetServicesInACategory(ctx context.Context, id int) ([]domain.Profession, error) {

	service, err := a.repository.GetServicesInACategory(ctx, id)
	if err != nil {
		return []domain.Profession{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []domain.Profession{}, errors.New("request timeout")
	}

	return service, nil
}
