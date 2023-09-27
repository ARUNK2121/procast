//go:build wireinject
// +build wireinject

package di

import (
	httpserver "github.com/ARUNK2121/procast/pkg/api"
	adminhandler "github.com/ARUNK2121/procast/pkg/api/handler/admin"
	"github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/db"
	"github.com/ARUNK2121/procast/pkg/helper"
	adminrepository "github.com/ARUNK2121/procast/pkg/repository/admin"
	adminusecase "github.com/ARUNK2121/procast/pkg/usecase/admin"

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
	)

	return &httpserver.ServerHTTP{}, nil
}
