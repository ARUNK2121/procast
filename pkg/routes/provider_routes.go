package routes

import (
	providerhandler "github.com/ARUNK2121/procast/pkg/api/handler/provider"
	"github.com/gin-gonic/gin"
)

func ProviderRoutes(
	engine *gin.RouterGroup, proHandler *providerhandler.AuthenticationHandler) {

	engine.POST("/register", proHandler.Register)
	engine.GET("/login", proHandler.Login)

	engine.Group("/profile")
	{
		//update profile picture
		//update details
		//post images and tag works
		//change preferences
		//location
	}

	engine.Group("/works")
	{
		//view leads
		//view currently participating bids
		//bid on a work
		//edit and re submit bid
		//view work details
		//check bids of other providers
	}

	engine.Group("/probook")
	{
		//Register to probook
		//add more services
		//remove a service
	}

	engine.Group("notification")
	{
		//first page list
		//open notification
	}

}
