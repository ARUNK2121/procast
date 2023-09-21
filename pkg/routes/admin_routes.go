package routes

import (
	"github.com/ARUNK2121/procast/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler) {
	engine.GET("/login", adminHandler.AdminLogin)
	panel := engine.Group("/panel")
	{
		panel.POST("", adminHandler.CreateNewAdmin)
		panel.DELETE("", adminHandler.DeleteAdmin)
	}

	category := engine.Group("/category")
	{
		category.POST("", categoryHandler.CreateCategory)
		// category.GET("", adminHandler.ListCategories)
		// category.DELETE("", adminHandler.DeleteCategory)
		// category.PATCH("", adminHandler.EditCategoryName)
	}

	// services := engine.Group("/services")
	// {
	// 	services.GET("", adminHandler.GetServicesInACategory)
	// 	services.POST("", adminHandler.AddServicesToACategory)
	// 	services.DELETE("", adminHandler.DeleteServices)
	// }

	// region := engine.Group("/region")
	// {
	// 	state := region.Group("/state")
	// 	{
	// 		state.GET("", adminHandler.GetStates)
	// 		state.POST("", adminHandler.AddNewState)
	// 		state.DELETE("", adminHandler.DeleteState)
	// 	}

	// 	district := region.Group("district")
	// 	{
	// 		district.GET("", GetDistricts)
	// 		district.POST("", AddDistricts)
	// 		district.DELETE("", DeleteDistrict)
	// 	}
	// }

	// verification := engine.Group("/verify")
	// {
	// 	verification.GET("", adminHandler.GetPendingVerification)
	// 	verification.PUT("", adminHandler.MakeProviderVerified)
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
