package providerusecase

import (
	"errors"

	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type profileUsecase struct {
	repository interfaces.ProfileRepository
}

func NewProfileUsecase(repo interfaces.ProfileRepository) services.ProfileUsecase {
	return &profileUsecase{
		repository: repo,
	}
}

func (p *profileUsecase) AddService(user_id int, service_id int) error {
	//check if the service already exists if not add the service
	exists, err := p.repository.CheckIfServiceIsAlreadyRegistered(user_id, service_id)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("already exists")
	}

	if err := p.repository.AddService(user_id, service_id); err != nil {
		return err
	}

	return nil
}

func (p *profileUsecase) DeleteService(user_id int, service_id int) error {
	//check if the service already exists is yes delete the service
	exists, err := p.repository.CheckIfServiceIsAlreadyRegistered(user_id, service_id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("service not found")
	}

	if err := p.repository.DeleteService(user_id, service_id); err != nil {
		return err
	}

	return nil
}

func (p *profileUsecase) AddPreferredWorkingLocation(id int, district int) error {
	//check if the service already exists if not add the service
	exists, err := p.repository.CheckIfDistrictIsAlreadyAdded(id, district)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("already exists")
	}

	if err := p.repository.AddLocation(id, district); err != nil {
		return err
	}

	return nil
}

func (p *profileUsecase) RemovePreferredLocation(id int, district int) error {

	if err := p.repository.RemovePreferredLocation(id, district); err != nil {
		return err
	}

	return nil
}

func (p *profileUsecase) GetMyServices(id int) ([]models.GetServices, error) {
	var model []models.GetServices
	//find serviceIDs
	all_id, err := p.repository.GetAllServiceIdsOfAProvider(id)
	if err != nil {
		return []models.GetServices{}, err
	}
	//find category
	for _, v := range all_id {
		service := models.GetServices{}
		service.ID = v
		details, err := p.repository.FindServiceDetailsFromID(v)
		if err != nil {
			return []models.GetServices{}, err
		}
		service.ServiceName = details.Profession

		service.Category_id = details.CategoryID
		model = append(model, service)
	}

	//find service name
	return model, nil
}

func (p *profileUsecase) GetAllPreferredLocations(id int) ([]models.GetLocations, error) {
	var model []models.GetLocations
	//find serviceIDs
	locations, err := p.repository.GetAllPreferredLocations(id)
	if err != nil {
		return []models.GetLocations{}, err
	}
	//find category
	for _, v := range locations {
		details, err := p.repository.GetLocationDetails(v)
		if err != nil {
			return []models.GetLocations{}, err
		}
		details.ID = v
		model = append(model, details)
	}

	//find service name
	return model, nil
}
