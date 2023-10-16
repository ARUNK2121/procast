package routes

import (
	providerhandler "github.com/ARUNK2121/procast/pkg/api/handler/provider"
	"github.com/ARUNK2121/procast/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func ProviderRoutes(
	engine *gin.RouterGroup,
	proHandler *providerhandler.AuthenticationHandler,
	profileHandler *providerhandler.ProfileHandler,
	workHandler *providerhandler.WorkHandler,
	notificationHandler *providerhandler.NotificationHandler) {

	engine.POST("/register", proHandler.Register) //completed
	engine.GET("/login", proHandler.Login)        //completed

	engine.Use(middleware.ProviderAuthMiddleware)
	{

		profile := engine.Group("/profile")
		{
			service := profile.Group("/service")
			{
				//list my services
				service.GET("", profileHandler.GetSelectedServices) //completed
				service.POST("", profileHandler.AddService)         //completed
				service.DELETE("", profileHandler.DeleteService)    //completed
			}

			location := profile.Group("location")
			{
				location.GET("", profileHandler.GetAllPreferredLocations)     //completed
				location.POST("", profileHandler.AddPreferredWorkingLocation) //completed
				location.DELETE("", profileHandler.RemovePreferredLocation)   //completed
			}
		}

		work := engine.Group("/works")
		{
			leads := work.Group("/leads")
			{
				leads.GET("", workHandler.GetAllLeads)                           //completed
				leads.GET("/:id", workHandler.ViewLeads)                         //completed
				leads.POST("/:id/bid", workHandler.PlaceBid)                     //completed
				leads.PUT("/:id/bid", workHandler.ReplaceBidWithNewBid)          //completed
				leads.GET("/:id/compare", workHandler.GetAllOtherBidsOnTheLeads) //completed
			}

			myWorks := work.Group("my-works")
			{
				myWorks.GET("", workHandler.GetWorksOfAProvider)         //completed
				myWorks.GET("/on-going", workHandler.GetAllOnGoingWorks) //completed
				myWorks.GET("/completed", workHandler.GetCompletedWorks) //completed
			}

			notification := engine.Group("notification")
			{
				notification.GET("", notificationHandler.GetAllNotifications)  //completed
				notification.GET("/:id", notificationHandler.ViewNotification) //completed
			}

		}

	}

}
