package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type AuthenticationUsecase interface {
	Register(models.ProviderRegister) error
	Login(models.Login) (string, error)
}
