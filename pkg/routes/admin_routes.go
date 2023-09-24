package routes

import (
	"github.com/ARUNK2121/procast/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(
	engine *gin.RouterGroup,
	adminHandler *handler.AdminHandler,
	categoryHandler *handler.CategoryHandler,
	servicehandler *handler.ServiceHandler,
	regionHandler *handler.RegionHandler,
	userManagementHandler *handler.UserManagementHandler) {

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
		services.POST("", servicehandler.AddServicesToACategory)
		services.GET("", servicehandler.GetServicesInACategory)
		services.DELETE("", servicehandler.DeleteService)
		services.PATCH("", servicehandler.ReActivateService)
	}

	region := engine.Group("/region")
	{
		state := region.Group("/state")
		{
			state.POST("", regionHandler.AddNewState)
			state.GET("", regionHandler.GetStates)
			state.DELETE("", regionHandler.DeleteState)
			state.PATCH("", regionHandler.ReActivateState)
		}

		district := region.Group("district")
		{
			district.POST("", regionHandler.AddNewDistrict)
			district.GET("", regionHandler.GetDistrictsFromState)
			district.DELETE("", regionHandler.DeleteDistrictFromState)
			district.PATCH("", regionHandler.ReActivateDistrict)
		}
	}

	providers := engine.Group("/provider")
	{
		providers.GET("", userManagementHandler.GetProviders)
		providers.PATCH("", userManagementHandler.MakeProviderVerified)
		providers.DELETE("", userManagementHandler.RevokeVerification)
	}

	users := engine.Group("/user")
	{
		users.GET("", userManagementHandler.GetUsers)
		users.DELETE("", userManagementHandler.BlockUser)
		users.PATCH("", userManagementHandler.UnBlockUser)
	}

	works := engine.Group("/works")
	{
		works.GET("/committed", servicehandler.ListCommittedWorks)
		// works.GET("/completed", servicehandler.ListCompletedWorks)
	}

	// verification := engine.Group("/verify")
	// {
	// 	verification.GET("", adminHandler.GetAllPendingVerifications)

	// 	request := verification.Group("request")
	// 	{
	// 		request.GET("", adminHandler.ViewVerificationRequest)
	// 	}

	// }

}
