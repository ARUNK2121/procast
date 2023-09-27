package httpserver

import (
	adminhandler "github.com/ARUNK2121/procast/pkg/api/handler/admin"
	"github.com/ARUNK2121/procast/pkg/routes"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *adminhandler.AdminHandler,
	categoryhandler *adminhandler.CategoryHandler,
	serviceHandler *adminhandler.ServiceHandler,
	regionHandler *adminhandler.RegionHandler,
	userManagementHandler *adminhandler.UserManagementHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryhandler, serviceHandler, regionHandler, userManagementHandler)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(":3000")
}
