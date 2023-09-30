package routes

import (
	providerhandler "github.com/ARUNK2121/procast/pkg/api/handler/provider"
	"github.com/gin-gonic/gin"
)

func ProviderRoutes(
	engine *gin.RouterGroup,
	proHandler *providerhandler.AuthenticationHandler,
	profileHandler *providerhandler.ProfileHandler,
	workHandler *providerhandler.WorkHandler) {

	engine.POST("/register", proHandler.Register) //completed
	engine.GET("/login", proHandler.Login)        //completed

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

		leads := work.Group("leads")
		{
			leads.GET("", workHandler.GetAllLeads)                           //not tested
			leads.GET("/:id", workHandler.ViewLeads)                         //not tested
			leads.POST("/:id/place-bid", workHandler.PlaceBid)               //not tested
			leads.POST("/:id/edit-bid", workHandler.ReplaceBidWithNewBid)    //not tested
			leads.GET("/:id/compare", workHandler.GetAllOtherBidsOnTheLeads) //not tested
		}

		// my_works:=work.Group("my-works")
		// {
		// 	my_works.GET("",workHandler.GetMyWorks)
		// 	my_works.GET("/completed",workHanlder.GetAllCompletedWorks)
		// 	my_works.GET("/on-going",workHandler.GetAllOnGoingWorks)
		// }

		// my_bids:=work.Group("my-bids")
		// {
		// 	my_bids:=work.Group("",workHandler.GetCurrentlyParticipatingBids)
		// }

	}

	// notification := engine.Group("notification")
	// {
	// 	notification.GET("", notificationHandler.GetAllNotifications)
	// 	notification.GET("/:id", notificationHandler.ViewNotification)
	// }

}
