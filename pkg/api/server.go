package httpserver

import (
	"github.com/ARUNK2121/procast/pkg/api/handler"
	"github.com/ARUNK2121/procast/pkg/routes"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler,
	categoryhandler *handler.CategoryHandler,
	serviceHandler *handler.ServiceHandler,
	regionHandler *handler.RegionHandler,
	userManagementHandler *handler.UserManagementHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryhandler, serviceHandler, regionHandler, userManagementHandler)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(":3000")
}
