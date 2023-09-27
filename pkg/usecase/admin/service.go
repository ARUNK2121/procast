package adminusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/admin/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/admin/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type serviceUsecase struct {
	repository               interfaces.ServiceRepository
	UserManagementRepository interfaces.UserManagementRepository
	RegionRepository         interfaces.RegionRepository
}

func NewServiceUsecase(repo interfaces.ServiceRepository, umr interfaces.UserManagementRepository, rr interfaces.RegionRepository) services.ServiceUsecase {
	return &serviceUsecase{
		repository:               repo,
		UserManagementRepository: umr,
		RegionRepository:         rr,
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

func (a *serviceUsecase) DeleteService(ctx context.Context, id int) error {

	err := a.repository.DeleteService(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (a *serviceUsecase) ReActivateService(ctx context.Context, id int) error {

	err := a.repository.ReActivateService(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (s *serviceUsecase) ListCommittedWorks(ctx context.Context) ([]models.WorkDetails, error) {
	var result []models.WorkDetails
	works, err := s.repository.GetCommittedWorks(ctx)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	for _, v := range works {
		//find district
		fmt.Println("district", v.DistrictID)
		DistrictName, err := s.RegionRepository.FindDistrictFromId(v.DistrictID)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find state
		StateName, err := s.RegionRepository.FindStateFromId(v.StateID)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find service
		Service, err := s.repository.FindServiceFromId(v.TargetProfessionID)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find user
		user, err := s.UserManagementRepository.FindUserFromId(v.UserID)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find provider
		provider, err := s.UserManagementRepository.FindProviderFromId(v.ProID)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		//get images

		result = append(result, models.WorkDetails{
			ID:         v.ID,
			Street:     v.Street,
			District:   DistrictName,
			State:      StateName,
			Profession: Service,
			User:       user,
			Provider:   provider,
			Images:     []string{},
			WorkStatus: v.WorkStatus,
		})
	}

	return result, nil
}

func (s *serviceUsecase) ListCompletedWorks(ctx context.Context) ([]models.WorkDetails, error) {
	var result []models.WorkDetails
	works, err := s.repository.GetCompletedWorks(ctx)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	for _, v := range works {
		//find district
		fmt.Println("district", v.DistrictID)
		DistrictName, err := s.RegionRepository.FindDistrictFromId(v.DistrictID)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find state
		StateName, err := s.RegionRepository.FindStateFromId(v.StateID)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find service
		Service, err := s.repository.FindServiceFromId(v.TargetProfessionID)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find user
		user, err := s.UserManagementRepository.FindUserFromId(v.UserID)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find provider
		provider, err := s.UserManagementRepository.FindProviderFromId(v.ProID)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		//get images

		result = append(result, models.WorkDetails{
			ID:         v.ID,
			Street:     v.Street,
			District:   DistrictName,
			State:      StateName,
			Profession: Service,
			User:       user,
			Provider:   provider,
			Images:     []string{},
			WorkStatus: v.WorkStatus,
		})
	}

	return result, nil
}
