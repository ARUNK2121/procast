package userusecase

import (
	helper "github.com/ARUNK2121/procast/pkg/helper/interfaces"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/user/interface"
)

type authenticationUsecase struct {
	repository interfaces.AuthenticationRepository
	helper     helper.Helper
}

func NewAuthenticationUsecase(repo interfaces.AuthenticationRepository, help helper.Helper) services.AuthenticationUsecase {
	return &authenticationUsecase{
		repository: repo,
		helper:     help,
	}
}
