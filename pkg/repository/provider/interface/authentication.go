package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type AuthenticationRepository interface {
	CheckIfPhoneNumberAlreadyExists(string) (bool, error)
	Register(models.ProviderRegister) (int, error)
	UploadDocument(int, string) error
}
