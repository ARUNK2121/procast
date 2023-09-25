package usecase

import (
	"context"
	"errors"

	"github.com/ARUNK2121/procast/pkg/repository/interfaces"
	services "github.com/ARUNK2121/procast/pkg/usecase/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
)

type userManagementUsecase struct {
	repository        interfaces.UserManagementRepository
	serviceRepository interfaces.ServiceRepository
}

func NewUserManagementUsecase(repo interfaces.UserManagementRepository, serviceRepo interfaces.ServiceRepository) services.UserManagementUsecase {
	return &userManagementUsecase{
		repository:        repo,
		serviceRepository: serviceRepo,
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

func (a *userManagementUsecase) BlockUser(ctx context.Context, id int) error {

	err := a.repository.BlockUser(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (a *userManagementUsecase) UnBlockUser(ctx context.Context, id int) error {

	err := a.repository.UnBlockUser(ctx, id)
	if err != nil {
		return err
	}
	err = ctx.Err()
	if err != nil {
		return errors.New("request timeout")
	}

	return nil
}

func (u *userManagementUsecase) GetAllPendingVerifications(ctx context.Context) ([]models.Verification, error) {

	verification, err := u.repository.GetAllPendingVerifications(ctx)
	if err != nil {
		return []models.Verification{}, err
	}
	err = ctx.Err()
	if err != nil {
		return []models.Verification{}, errors.New("request timeout")
	}

	return verification, nil
}

func (u *userManagementUsecase) ViewVerificationRequest(ctx context.Context, id int) (models.VerificationDetails, error) {

	var verification models.VerificationDetails
	//id
	verification.ID = id
	//name
	name, err := u.repository.FindProviderFromId(id)
	if err != nil {
		return models.VerificationDetails{}, err
	}

	verification.Name = name
	//document image
	image, err := u.repository.FindDocumentsOfProviderFromID(id)
	if err != nil {
		return models.VerificationDetails{}, err
	}

	verification.DocumentImage = image
	//all services as an array
	services, err := u.serviceRepository.FindIdOfServicesOfAProvider(id)
	if err != nil {
		return models.VerificationDetails{}, err
	}

	var serviceNames []string

	for _, v := range services {
		s, err := u.serviceRepository.FindServiceFromId(v)
		if err != nil {
			return models.VerificationDetails{}, err
		}

		serviceNames = append(serviceNames, s)
	}

	verification.Services = serviceNames

	return verification, nil
}
