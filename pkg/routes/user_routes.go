package routes

import (
	userhandler "github.com/ARUNK2121/procast/pkg/api/handler/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(
	engine *gin.RouterGroup,
	authenticationHandler *userhandler.AuthenticationHandler) {

	engine.POST("signup", authenticationHandler.UserSignup)
	// engine.GET("login",authenticationHandler.Login)

	// profile:=engine.Group("/profile")
	// {

	// }

	// works:=engine.Group("/woks")
	// {

	// }

	// notification:=engine.Group("/notification")
	// {

	// }
}

//user profile editing
//open a work
//provider profile inspection
//Up-voting or down-voting provider
//check notifications
//list works
//view bids
//hand-over work
//declare the work completed and rate work
