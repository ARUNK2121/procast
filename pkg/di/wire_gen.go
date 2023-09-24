// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/ARUNK2121/procast/pkg/api"
	"github.com/ARUNK2121/procast/pkg/api/handler"
	"github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/db"
	"github.com/ARUNK2121/procast/pkg/helper"
	"github.com/ARUNK2121/procast/pkg/repository"
	"github.com/ARUNK2121/procast/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*httpserver.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	adminRepository := repository.NewAdminRepository(gormDB)
	interfacesHelper := helper.NewHelper(cfg)
	adminUsecase := usecase.NewAdminUsecase(adminRepository, interfacesHelper)
	adminHandler := handler.NewAdminHandler(adminUsecase)
	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)
	serviceRepository := repository.NewServiceRepository(gormDB)
	userManagementRepository := repository.NewUserManagementRepository(gormDB)
	regionRepository := repository.NewRegionRepository(gormDB)
	serviceUsecase := usecase.NewServiceUsecase(serviceRepository, userManagementRepository, regionRepository)
	serviceHandler := handler.NewServiceHandler(serviceUsecase)
	regionUsecase := usecase.NewRegionUsecase(regionRepository)
	regionHandler := handler.NewRegionHandler(regionUsecase)
	userManagementUsecase := usecase.NewUserManagementUsecase(userManagementRepository)
	userManagementHandler := handler.NewUserManagementHandler(userManagementUsecase)
	serverHTTP := httpserver.NewServerHTTP(adminHandler, categoryHandler, serviceHandler, regionHandler, userManagementHandler)
	return serverHTTP, nil
}
