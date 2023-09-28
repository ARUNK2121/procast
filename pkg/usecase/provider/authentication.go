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

	fmt.Println("1")

	hashed, err := a.helper.CreateHashPassword(model.Password)
	if err != nil {
		return err
	}

	fmt.Println("2")

	model.Password = hashed

	fmt.Println("3")
	// pass to repository
	id, err := a.repository.Register(model)
	if err != nil {
		return err
	}

	fmt.Println("4")

	filename, err := a.helper.UploadToS3(model.Document)
	if err != nil {
		return err
	}

	fmt.Println("5")

	if err := a.repository.UploadDocument(id, filename); err != nil {
		return err
	}

	fmt.Println("6")

	return nil
}

func (a *authenticationUsecase) Login(models.ProLogin) (string, error) {
	//check if username matches email
	//if not then check phone
	//retrieve user details
	//compare details
	//if matches create token

	return "token", nil

}
