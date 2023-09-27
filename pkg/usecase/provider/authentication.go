package providerusecase

import (
	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
)

type authenticationUsecase struct {
	repository interfaces.AuthenticationRepository
}

func NewAuthenticationUsecase(repo interfaces.AuthenticationRepository) services.AuthenticationUsecase {
	return &authenticationUsecase{
		repository: repo,
	}
}
