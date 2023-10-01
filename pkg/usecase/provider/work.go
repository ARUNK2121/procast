package providerusecase

import (
	"fmt"

	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type workUseCase struct {
	repository        interfaces.WorkRepository
	profileRepository interfaces.ProfileRepository
}

func NewWorkUseCase(repo interfaces.WorkRepository, profile interfaces.ProfileRepository) services.WorkUseCase {
	return &workUseCase{
		repository:        repo,
		profileRepository: profile,
	}
}

func (w *workUseCase) GetAllLeads(pro_id int, page int) ([]models.WorkDetails, error) {

	//get providers preferred services
	services, err := w.profileRepository.GetAllServiceIdsOfAProvider(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	fmt.Println("services", services)

	//get providers preffered locations
	locations, err := w.profileRepository.GetAllPreferredLocations(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	fmt.Println("locations", locations)

	//get work ids by these preferences
	var works []int
	for _, service := range services {
		for _, location := range locations {
			lead, err := w.repository.GetLeadByServiceAndLocation(service, location)
			if err != nil {
				return []models.WorkDetails{}, err
			}
			fmt.Println("lead", lead)
			works = append(works, lead...)

		}
	}

	var model []models.WorkDetails
	for _, v := range works {
		//find details
		details, err := w.repository.GetDetailsOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		fmt.Println("details:", details)
		//find images
		images, err := w.repository.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		exists, err := w.repository.CheckIfAlreadyBidded(v, pro_id)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		//append
		var result models.WorkDetails
		result.ID = v
		result.Street = details.Street
		result.District = details.District
		result.State = details.State
		result.Profession = details.Profession
		result.User = details.User
		result.Images = images
		result.Participation = exists
		result.WorkStatus = details.WorkStatus
		model = append(model, result)
	}
	//pack into model and return the model

	fmt.Println("model", model)
	return model, nil
}

func (w *workUseCase) ViewLeads(work_id int) (models.WorkDetails, error) {
	details, err := w.repository.GetDetailsOfAWork(work_id)
	if err != nil {
		return models.WorkDetails{}, err
	}
	//find images
	images, err := w.repository.GetImagesOfAWork(work_id)
	if err != nil {
		return models.WorkDetails{}, err
	}
	//append
	var result models.WorkDetails
	result.ID = work_id
	result.Street = details.Street
	result.District = details.District
	result.State = details.State
	result.Profession = details.Profession
	result.User = details.User
	result.Images = images
	result.WorkStatus = details.WorkStatus

	return result, nil
}

func (w *workUseCase) PlaceBid(model models.PlaceBid) error {
	err := w.repository.PlaceBid(model)
	if err != nil {
		return err
	}

	return nil
}

func (w *workUseCase) ReplaceBidWithNewBid(model models.PlaceBid) error {
	err := w.repository.ReplaceBidWithNewBid(model)
	if err != nil {
		return err
	}

	return nil
}

func (w *workUseCase) GetAllOtherBidsOnTheLeads(work_id int) ([]models.BidDetails, error) {
	bids, err := w.repository.GetAllOtherBidsOnTheLeads(work_id)
	if err != nil {
		return []models.BidDetails{}, err
	}

	return bids, nil
}

func (w *workUseCase) GetMyWorks(pro_id int) ([]models.WorkDetails, error) {

	provider, err := w.repository.FindProviderName(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	works, err := w.repository.GetAllWorksOfAProvider(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	var model []models.WorkDetails

	for _, v := range works {
		details, err := w.repository.GetDetailsOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find images
		images, err := w.repository.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//append
		var result models.WorkDetails
		result.ID = v
		result.Street = details.Street
		result.District = details.District
		result.State = details.State
		result.Profession = details.Profession
		result.User = details.User
		result.Provider = provider
		result.Images = images
		result.WorkStatus = details.WorkStatus

		model = append(model, result)

	}

	return model, nil
}
