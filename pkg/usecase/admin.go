package usecase

import (
	"fmt"

	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/interfaces"
)

type AdminUsecase struct {
	Repository interfaces.AdminRepository
}

func NewAdminUsecase(repo interfaces.AdminRepository) services.AdminUsecase {
	return &AdminUsecase{
		Repository: repo,
	}
}

func (a *AdminUsecase) A() {
	fmt.Println("hyy")
}
