package usecase

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type userManagementUsecase struct {
	repository interfaces.UserManagementRepository
}

func NewUserManagementUsecase(repo interfaces.UserManagementRepository) services.UserManagementUsecase {
	return &userManagementUsecase{
		repository: repo,
	}
}

func (u *userManagementUsecase) GetProviders(ctx context.Context) ([]models.ProviderDetails, error) {

	users, err := u.repository.GetProviders(ctx)
	if err != nil {
		return []models.ProviderDetails{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []models.ProviderDetails{}, errors.New("request timeout")
	}

	return users, nil
}

func (a *userManagementUsecase) MakeProviderVerified(ctx context.Context, id int) error {

	err := a.repository.MakeProviderVerified(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (a *userManagementUsecase) RevokeVerification(ctx context.Context, id int) error {

	err := a.repository.RevokeVerification(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (u *userManagementUsecase) GetUsers(ctx context.Context) ([]models.UserDetails, error) {

	users, err := u.repository.GetUsers(ctx)
	if err != nil {
		return []models.UserDetails{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []models.UserDetails{}, errors.New("request timeout")
	}

	return users, nil
}
