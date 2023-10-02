//go:build wireinject
// +build wireinject

package di

import (
	httpserver "github.com/ARUNK2121/procast/pkg/api"
	adminhandler "github.com/ARUNK2121/procast/pkg/api/handler/admin"
	providerhandler "github.com/ARUNK2121/procast/pkg/api/handler/provider"
	"github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/db"
	"github.com/ARUNK2121/procast/pkg/helper"
	adminrepository "github.com/ARUNK2121/procast/pkg/repository/admin"
	providerRepository "github.com/ARUNK2121/procast/pkg/repository/provider"
	adminusecase "github.com/ARUNK2121/procast/pkg/usecase/admin"
	providerusecase "github.com/ARUNK2121/procast/pkg/usecase/provider"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*httpserver.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		adminrepository.NewAdminRepository,
		adminusecase.NewAdminUsecase,
		adminhandler.NewAdminHandler,
		httpserver.NewServerHTTP,
		helper.NewHelper,
		adminhandler.NewCategoryHandler,
		adminusecase.NewCategoryUsecase,
		adminrepository.NewCategoryRepository,
		adminrepository.NewServiceRepository,
		adminusecase.NewServiceUsecase,
		adminhandler.NewServiceHandler,
		adminrepository.NewRegionRepository,
		adminusecase.NewRegionUsecase,
		adminhandler.NewRegionHandler,
		adminrepository.NewUserManagementRepository,
		adminusecase.NewUserManagementUsecase,
		adminhandler.NewUserManagementHandler,
		providerusecase.NewAuthenticationUsecase,
		providerhandler.NewAuthenticationHandler,
		providerRepository.NewAuthenticationRepository,
		providerhandler.NewProfileHandler,
		providerusecase.NewProfileUsecase,
		providerRepository.NewProfileRepository,
		providerhandler.NewWorkHandler,
		providerusecase.NewWorkUseCase,
		providerRepository.NewWorkRepository,
		providerhandler.NewNotificationHandler,
		providerusecase.NewNotificationUsecase,
		providerRepository.NewNotificationRepository,
	)

	return &httpserver.ServerHTTP{}, nil
}
