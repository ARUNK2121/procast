package routes

import (
	providerhandler "github.com/ARUNK2121/procast/pkg/api/handler/provider"
	userhandler "github.com/ARUNK2121/procast/pkg/api/handler/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(
	engine *gin.RouterGroup,
	authenticationHandler *userhandler.AuthenticationHandler,
	workHandler *userhandler.WorkHandler,
	providerworkhandler *providerhandler.WorkHandler) {

	engine.POST("signup", authenticationHandler.UserSignup) //completed
	engine.GET("login", authenticationHandler.Login)        //completed

	works := engine.Group("/works")
	{
		works.POST("", workHandler.ListNewOpening)   //completed
		works.GET("", workHandler.GetAllListedWorks) //completed

		works.GET("/ongoing", workHandler.ListAllOngoingWorks)     //completed
		works.GET("/completed", workHandler.ListAllCompletedWorks) //completed

		workManagement := works.Group("/:id")
		{
			workManagement.GET("", workHandler.WorkDetails)                            //completed
			workManagement.GET("/bids", providerworkhandler.GetAllOtherBidsOnTheLeads) //completed
			workManagement.PUT("/assign", workHandler.AssignWorkToProvider)            //completed
			workManagement.PUT("/complete", workHandler.MakeWorkAsCompleted)           //completed
		}

	}

	// profile := engine.Group("/profile")
	// {
	// 	profile.GET("", profileHandler.GetProfileDetails)
	// 	profile.PUT("", profileHandler.EditProfilePicture)
	// }

	// provider := engine.Group("/provider")
	// {
	// 	provider.GET("/:pro-id", workHandler.GetDetailsOfProviders)
	//  provider.GET("/:pro_id/works",workHandler.GetWorksOfAProvider)
	// }

	// notification := engine.Group("notification")
	// {
	// 	notification.GET("", notificationHandler.GetAllNotifications)  //completed
	// 	notification.GET("/:id", notificationHandler.ViewNotification) //completed
	// }

}
