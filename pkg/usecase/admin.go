package usecase

import (
	"context"

	helper "github.com/ARUNK2121/procast/pkg/helper/interfaces"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type adminUsecase struct {
	repository interfaces.AdminRepository
	helper     helper.Helper
}

func NewAdminUsecase(repo interfaces.AdminRepository, helper helper.Helper) services.AdminUsecase {
	return &adminUsecase{
		repository: repo,
		helper:     helper,
	}
}

func (ad *adminUsecase) AdminLogin(ctx context.Context, model models.AdminLogin) (models.Tokens, error) {
	// getting details of the admin based on the email provided
	adminCompareDetails, err := ad.repository.GetAdminDetailsByEmail(model.Email)
	if err != nil {
		return models.Tokens{}, err
	}

	// compare password from database and that provided from admins
	err = ad.helper.CompareHashAndPassword(adminCompareDetails.Password, model.Password)
	if err != nil {
		return models.Tokens{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

	adminDetailsResponse.ID = int(adminCompareDetails.ID)
	adminDetailsResponse.Email = adminCompareDetails.Email
	adminDetailsResponse.Name = adminCompareDetails.Name
	adminDetailsResponse.Previlege = adminCompareDetails.Previlege

	access, refresh, err := ad.helper.GenerateTokenAdmin(adminDetailsResponse)
	if err != nil {
		return models.Tokens{}, err
	}

	return models.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil

}