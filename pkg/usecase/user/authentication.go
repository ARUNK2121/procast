package userusecase

import (
	"errors"

	helper "github.com/ARUNK2121/procast/pkg/helper/interfaces"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/user/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/user/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
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

func (a *authenticationUsecase) UserSignup(model models.UserSignup) error {
	// check if phone number already exists
	exists, err := a.repository.CheckIfPhoneNumberAlreadyExists(model.Phone)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("phone number already exists")
	}

	//check if password matches
	if model.Password != model.ConfirmPassword {
		return errors.New("password mismatch")
	}

	hashed, err := a.helper.CreateHashPassword(model.Password)
	if err != nil {
		return err
	}

	model.Password = hashed

	// pass to repository
	if err := a.repository.UserSignup(model); err != nil {
		return err
	}

	return nil
}
