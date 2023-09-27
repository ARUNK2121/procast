package routes

import (
	providerhandler "github.com/ARUNK2121/procast/pkg/api/handler/provider"
	"github.com/gin-gonic/gin"
)

func ProviderRoutes(
	engine *gin.RouterGroup, proHandler *providerhandler.AuthenticationHandler) {

	// engine.GET("/register", proHandler.Register)
	// engine.DELETE("/logout", adminHandler.AdminLogout)

	//provider sign up
	//provider login

	//selecting preferred working locations
	//provider profile editing
	//provider view leads
	//provider place bids with commit request
	//check bids of other providers
	//view currently active bids
	//view notifications
	//Promotion from users

	engine.Group("/profile")
	{
		//update profile picture
		//update details
		//post images and tag works
		//
	}
	engine.Group("/works")
	{

	}

	engine.Group("/probook")
	{
		//Register to probook
	}

}
