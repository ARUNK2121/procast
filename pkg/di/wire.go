//go:build wireinject
// +build wireinject

package di

import (
	httpserver "github.com/ARUNK2121/procast/pkg/api"
	handler "github.com/ARUNK2121/procast/pkg/api/handler"
	"github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/db"
	"github.com/ARUNK2121/procast/pkg/helper"
	"github.com/ARUNK2121/procast/pkg/repository"
	"github.com/ARUNK2121/procast/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*httpserver.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewAdminRepository,
		usecase.NewAdminUsecase,
		handler.NewAdminHandler,
		httpserver.NewServerHTTP,
		helper.NewHelper,
		handler.NewCategoryHandler,
		usecase.NewCategoryUsecase,
		repository.NewCategoryRepository,
		repository.NewServiceRepository,
		usecase.NewServiceUsecase,
		handler.NewServiceHandler,
		repository.NewRegionRepository,
		usecase.NewRegionUsecase,
		handler.NewRegionHandler,
		repository.NewUserManagementRepository,
		usecase.NewUserManagementUsecase,
		handler.NewUserManagementHandler,
	)

	return &httpserver.ServerHTTP{}, nil
}
