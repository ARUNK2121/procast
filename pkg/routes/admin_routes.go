package routes

import (
	"github.com/ARUNK2121/procast/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup,
	adminHandler *handler.AdminHandler,
	categoryHandler *handler.CategoryHandler) {

	engine.GET("/login", adminHandler.AdminLogin)
	// engine.DELETE("/logout", adminHandler.AdminLogout)

	panel := engine.Group("/panel")
	{
		panel.POST("", adminHandler.CreateNewAdmin)
		panel.DELETE("", adminHandler.DeleteAdmin)
	}

	category := engine.Group("/category")
	{
		category.POST("", categoryHandler.CreateCategory)
		category.GET("", categoryHandler.ListCategories)
		category.DELETE("", categoryHandler.DeleteCategory)
		category.PATCH("", categoryHandler.ReActivateCategory)
	}

	services := engine.Group("/services")
	{
		services.POST("", adminHandler.AddServicesToACategory)
		// services.GET("", adminHandler.GetServicesInACategory)
		// services.DELETE("", adminHandler.DeleteServices)
	}

	// region := engine.Group("/region")
	// {
	// 	state := region.Group("/state")
	// 	{
	// 		state.POST("", adminHandler.AddNewState)
	// 		state.GET("", adminHandler.GetStates)
	// 		state.DELETE("", adminHandler.DeleteState)
	// 	}

	// 	district := region.Group("district")
	// 	{
	// 		district.POST("", AddDistricts)
	// 		district.GET("", GetDistricts)
	// 		district.DELETE("", DeleteDistrict)
	// 	}
	// }

	// verification := engine.Group("/verify")
	// {
	// 	verification.GET("", adminHandler.GetAllPendingVerifications)

	// 	request := verification.Group("request")
	// 	{
	// 		request.GET("", adminHandler.ViewVerificationRequest)
	// 		request.PUT("", adminHandler.MakeProviderVerified)
	// 	}
	// }

	// works := engine.Group("/works")
	// {
	// 	works.GET("", adminHandler.ListScheduledworks)
	// }

	// providers := engine.Group("/provider")
	// {
	// 	providers.GET("/top", adminHandler.GetTopProviders)
	// 	providers.GET("", adminHandler.GetProviders)
	// 	providers.DELETE("", adminHandler.RevokeVerification)
	// }

	// users := engine.Group("/user")
	// {
	// 	users.GET("", adminHandler.GetUsers)
	// 	users.DELETE("", adminHandler.BlockUser)
	// }

}
