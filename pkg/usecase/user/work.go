package userusecase

import (
	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/user/interface"
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
