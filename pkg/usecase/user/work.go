package userusecase

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/user/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type workUsecase struct {
	repository interfaces.WorkRepository
}

func NewWorkUseCase(repo interfaces.WorkRepository) services.WorkUsecase {
	return &workUsecase{
		repository: repo,
	}
}

func (w *workUsecase) ListNewOpening(model domain.Work) error {

	//pass to repository
	err := w.repository.ListNewOpening(model)
	if err != nil {
		return err
	}

	return nil
}

func (w *workUsecase) GetAllListedWorks(id int) ([]models.WorkDetails, error) {

	works, err := w.repository.GetAllWorksOfAUser(id)
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

		var provider string

		pro_id, err := w.repository.FindProviderIdFromWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		if pro_id == 0 {
			provider = "not assigned"
		} else {
			provider, err = w.repository.FindProviderName(pro_id)
			if err != nil {
				return []models.WorkDetails{}, err
			}
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

func (w *workUsecase) ListAllCompletedWorks(id int) ([]models.WorkDetails, error) {

	works, err := w.repository.GetAllCompletedWorksOfAUser(id)
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

		var provider string

		pro_id, err := w.repository.FindProviderIdFromWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		if pro_id == 0 {
			provider = "not assigned"
		} else {
			provider, err = w.repository.FindProviderName(pro_id)
			if err != nil {
				return []models.WorkDetails{}, err
			}
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

func (w *workUsecase) ListAllOngoingWorks(id int) ([]models.WorkDetails, error) {

	works, err := w.repository.GetAllOngoingWorksOfAUser(id)
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

		var provider string

		pro_id, err := w.repository.FindProviderIdFromWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		if pro_id == 0 {
			provider = "not assigned"
		} else {
			provider, err = w.repository.FindProviderName(pro_id)
			if err != nil {
				return []models.WorkDetails{}, err
			}
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

func (w *workUsecase) WorkDetails(id int) (models.WorkDetails, error) {

	details, err := w.repository.GetDetailsOfAWork(id)
	if err != nil {
		return models.WorkDetails{}, err
	}
	//find images
	images, err := w.repository.GetImagesOfAWork(id)
	if err != nil {
		return models.WorkDetails{}, err
	}

	var provider string

	pro_id, err := w.repository.FindProviderIdFromWork(id)
	if err != nil {
		return models.WorkDetails{}, err
	}

	if pro_id == 0 {
		provider = "not assigned"
	} else {
		provider, err = w.repository.FindProviderName(pro_id)
		if err != nil {
			return models.WorkDetails{}, err
		}
	}

	//append
	var result models.WorkDetails
	result.ID = id
	result.Street = details.Street
	result.District = details.District
	result.State = details.State
	result.Profession = details.Profession
	result.User = details.User
	result.Provider = provider
	result.Images = images
	result.WorkStatus = details.WorkStatus

	return result, nil
}

func (w *workUsecase) AssignWorkToProvider(work_id, pro_id int) error {

	//pass to repository
	err := w.repository.AssignWorkToProvider(work_id, pro_id)
	if err != nil {
		return err
	}

	return nil
}

func (w *workUsecase) MakeWorkAsCompleted(id int) error {

	//pass to repository
	err := w.repository.MakeWorkAsCompleted(id)
	if err != nil {
		return err
	}

	return nil
}
