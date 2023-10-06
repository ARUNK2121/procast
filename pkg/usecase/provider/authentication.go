package providerusecase

import (
	"errors"
	"fmt"

	helper "github.com/ARUNK2121/procast/pkg/helper/interfaces"
	interfaces "github.com/ARUNK2121/procast/pkg/repository/provider/interface"
	services "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
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

func (a *authenticationUsecase) Register(model models.ProviderRegister) error {
	// check if phone number already exists
	exists, err := a.repository.CheckIfPhoneNumberAlreadyExists(model.Phone)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("phone number already exists")
	}

	//check if password matches
	if model.Password != model.RePassword {
		return errors.New("password mismatch")
	}

	hashed, err := a.helper.CreateHashPassword(model.Password)
	if err != nil {
		return err
	}

	model.Password = hashed

	// pass to repository
	id, err := a.repository.Register(model)
	if err != nil {
		return err
	}

	filename, err := a.helper.UploadToS3(model.Document)
	if err != nil {
		return err
	}

	if err := a.repository.UploadDocument(id, filename); err != nil {
		return err
	}

	return nil
}

func (a *authenticationUsecase) Login(model models.Login) (string, error) {

	exists, err := a.repository.CheckIfProviderExistsByUsername(model.Username)
	if err != nil {
		return "", err
	}

	fmt.Println("1")

	if !exists {
		return "", errors.New("check username again")
	}

	details, err := a.repository.GetProviderDetailsByUsername(model.Username)
	if err != nil {
		return "", err
	}

	fmt.Println("2")

	err = a.helper.CompareHashAndPassword(details.Password, model.Password)
	if err != nil {
		return "", errors.New("pasword mismatch")
	}
	fmt.Println("3")

	if !details.IsVerified {
		return "", errors.New("your request is under validation")
	}

	token, err := a.helper.GenerateTokenProvider(details)
	if err != nil {
		return token, err
	}

	fmt.Println("4")

	//if matches create token

	return token, nil

}
