package routes

import (
	adminhandler "github.com/ARUNK2121/procast/pkg/api/handler/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(
	engine *gin.RouterGroup,
	adminHandler *adminhandler.AdminHandler,
	categoryHandler *adminhandler.CategoryHandler,
	servicehandler *adminhandler.ServiceHandler,
	regionHandler *adminhandler.RegionHandler,
	userManagementHandler *adminhandler.UserManagementHandler) {

	engine.POST("/login", adminHandler.AdminLogin)
	// engine.DELETE("/logout", adminHandler.AdminLogout)

	// engine.Use(middleware.AdminAuthMiddleware)
	{
		panel := engine.Group("/panel")
		// panel.Use(middleware.SuperAdminAuthMiddleware)
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

			district := region.Group("/district")
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
			works.GET("/completed", servicehandler.ListCompletedWorks)
		}

		verification := engine.Group("/verify")
		{
			verification.GET("", userManagementHandler.GetAllPendingVerifications)

			request := verification.Group("request")
			{
				request.GET("", userManagementHandler.ViewVerificationRequest)
			}

		}

	}

}
