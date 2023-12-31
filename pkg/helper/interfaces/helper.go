package interfaces

import (
	"mime/multipart"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type Helper interface {
	CompareHashAndPassword(a string, b string) error
	Copy(a interface{}, b interface{}) (interface{}, error)
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	CreateHashPassword(string) (string, error)
	UploadToS3(file *multipart.FileHeader) (string, error)
	GenerateTokenProvider(details domain.Provider) (string, error)
	GenerateTokenUser(details domain.User) (string, error)
}
