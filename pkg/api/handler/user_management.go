package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/ARUNK2121/procast/pkg/usecase/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserManagementHandler struct {
	usecase interfaces.UserManagementUsecase
}

func NewUserManagementHandler(usecase interfaces.UserManagementUsecase) *UserManagementHandler {
	return &UserManagementHandler{
		usecase: usecase,
	}
}

func (u *UserManagementHandler) GetProviders(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	//call usecase get array
	providers, err := u.usecase.GetProviders(ctx)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: providers, Error: nil}
	c.JSON(http.StatusOK, successRes)
}
