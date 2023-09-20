package routes

import (
	"github.com/ARUNK2121/procast/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler) {
	engine.GET("/login", adminHandler.AdminLogin)
}
