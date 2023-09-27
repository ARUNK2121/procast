package adminusecase

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	helper "github.com/ARUNK2121/procast/pkg/helper/interfaces"
	"github.com/ARUNK2121/procast/pkg/repository/admin/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/admin/interfaces"
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

	if ctx.Err() != nil {
		return models.Tokens{}, errors.New("request time out")
	}
	// getting details of the admin based on the email provided
	adminCompareDetails, err := ad.repository.GetAdminDetailsByEmail(ctx, model.Email)
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

func (ad *adminUsecase) CreateNewAdmin(ctx context.Context, model domain.Admin) error {

	if err := ctx.Err(); err != nil {
		return errors.New("request time out")
	}

	//check if already there is an admin with existing mail
	count, err := ad.repository.CountOfAdminByEmail(ctx, model.Email)
	if err != nil {
		return err
	}

	if count != 0 {
		return errors.New("email already exists")
	}

	//hash password and store it in model

	hash, err := ad.helper.CreateHashPassword(model.Password)
	if err != nil {
		return err
	}

	//make previlege as normal admin
	model.Password = hash
	model.Previlege = "admin"

	//if no admin create new admin
	if err := ad.repository.CreateNewAdmin(ctx, model); err != nil {
		return err
	}

	return nil
}

func (a *adminUsecase) DeleteAdmin(ctx context.Context, id int) error {

	err := a.repository.DeleteAdmin(ctx, id)
	if err != nil || ctx.Err() != nil {
		return nil
	}

	return nil
}
