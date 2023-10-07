package userusecase

import (
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/user/interface"
)

type profileUsecase struct {
	repository interfaces.ProfileRepository
}

func NewProfileUsecase(repo interfaces.ProfileRepository) services.ProfileUsecase {
	return &profileUsecase{
		repository: repo,
	}
}
