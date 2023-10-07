package routes

import (
	userhandler "github.com/ARUNK2121/procast/pkg/api/handler/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(
	engine *gin.RouterGroup,
	authenticationHandler *userhandler.AuthenticationHandler,
	workHandler *userhandler.WorkHandler) {

	engine.POST("signup", authenticationHandler.UserSignup) //completed
	engine.GET("login", authenticationHandler.Login)        //completed

	// profile := engine.Group("/profile")
	// {
	// 	profile.GET("", profileHandler.GetProfileDetails)
	// 	profile.PUT("", profileHandler.EditProfilePicture)
	// }

	works := engine.Group("/works")
	{
		works.POST("", workHandler.ListNewOpening)
		works.GET("", workHandler.GetAllListedWorks)
		// works.GET("/completed",workHandler.ListAllCompletedWorks)
		// works.GET("/ongoing",workHandler.ListAllOngoingWorks)

		// workManagement := works.Group("/:id")
		// {
		// 	workManagement.GET("", workHandler.WorkDetails)
		// 	workManagement.GET("/bids", workHandler.CompareBids)
		// 	workManagement.PUT("/assign", workHandler.AssignWorkToProvider)
		// 	workManagement.PUT("/complete", workHandler.MakeWorkAsCompleted)
		// }

	}

	// provider := engine.Group("/provider")
	// {
	// 	provider.GET("/:pro-id", workHandler.GetDetailsOfProviders)
	// }

	// notification := engine.Group("notification")
	// {
	// 	notification.GET("", notificationHandler.GetAllNotifications)  //completed
	// 	notification.GET("/:id", notificationHandler.ViewNotification) //completed
	// }

}
