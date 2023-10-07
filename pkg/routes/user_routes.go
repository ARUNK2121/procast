package routes

import (
	userhandler "github.com/ARUNK2121/procast/pkg/api/handler/user"
	"github.com/gin-gonic/gin"
)

func UserRoutes(
	engine *gin.RouterGroup,
	authenticationHandler *userhandler.AuthenticationHandler) {

	engine.POST("signup", authenticationHandler.UserSignup) //completed
	engine.GET("login", authenticationHandler.Login)        //completed

	// profile := engine.Group("/profile")
	// {
	// 	profile.GET("", profileHandler.GetProfileDetails)
	// 	profile.PUT("", profileHandler.EditProfilePicture)
	// }

	works := engine.Group("/woks")
	{
		works.POST("", workHandler.ListNewOpening)
		// works.GET("", workHandler.GetMyWorks)

		// workManagement := works.Group("/:id")
		// {
		// 	workManagement.GET("", workHandler.WorkDetails)
		// 	workManagement.GET("/bids", workHandler.CompareBids)
		// 	workManagement.PUT("/assign", workHandler.AssignWorkToProvider)
		// 	workManagement.PUT("/complete", workHandler.MakeWorkAsCompleted)
		// }

	}
	// notification := engine.Group("notification")
	// {
	// 	notification.GET("", notificationHandler.GetAllNotifications)  //completed
	// 	notification.GET("/:id", notificationHandler.ViewNotification) //completed
	// }

}
