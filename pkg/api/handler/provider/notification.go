package providerhandler

import (
	"net/http"
	"strconv"

	interfaces "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	usecase interfaces.NotificationUsecase
}

func NewNotificationHandler(use interfaces.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{
		usecase: use,
	}
}

func (n *NotificationHandler) GetAllNotifications(c *gin.Context) {
	provider_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	notifications, err := n.usecase.GetAllNotifications(provider_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: notifications, Error: nil}
	c.JSON(http.StatusOK, res)
}
