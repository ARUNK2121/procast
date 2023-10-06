package httpserver

import (
	adminhandler "github.com/ARUNK2121/procast/pkg/api/handler/admin"
	providerhandler "github.com/ARUNK2121/procast/pkg/api/handler/provider"
	userhandler "github.com/ARUNK2121/procast/pkg/api/handler/user"
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
	userManagementHandler *adminhandler.UserManagementHandler,
	ProviderAuthenticationHandler *providerhandler.AuthenticationHandler,
	profileHandler *providerhandler.ProfileHandler,
	workHandler *providerhandler.WorkHandler,
	notificationhandler *providerhandler.NotificationHandler,
	userAutheticationHandler *userhandler.AuthenticationHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryhandler, serviceHandler, regionHandler, userManagementHandler)
	routes.ProviderRoutes(engine.Group("/provider"), ProviderAuthenticationHandler, profileHandler, workHandler, notificationhandler)
	routes.UserRoutes(engine.Group("/user"), userAutheticationHandler)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(":3000")
}
