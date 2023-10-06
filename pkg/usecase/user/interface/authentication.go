package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type AuthenticationUsecase interface {
	UserSignup(models.UserSignup) error
}
