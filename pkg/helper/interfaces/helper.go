package interfaces

import "github.com/ARUNK2121/procast/pkg/utils/models"

type Helper interface {
	CompareHashAndPassword(a string, b string) error
	Copy(a interface{}, b interface{}) (interface{}, error)
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error)
	CreateHashPassword(string) (string, error)
}
