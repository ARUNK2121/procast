package usecase

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/interfaces"
)

type userManagementUsecase struct {
	repository interfaces.UserManagementRepository
}

func NewUserManagementUsecase(repo interfaces.UserManagementRepository) services.UserManagementUsecase {
	return &userManagementUsecase{
		repository: repo,
	}
}

func (u *userManagementUsecase) GetProviders(ctx context.Context) ([]domain.Provider, error) {

	states, err := u.repository.GetProviders(ctx)
	if err != nil {
		return []domain.Provider{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []domain.Provider{}, errors.New("request timeout")
	}

	return states, nil
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
