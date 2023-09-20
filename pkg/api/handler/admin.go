package handler

import (
	"context"
	"net/http"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/usecase/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase interfaces.AdminUsecase
}

func NewAdminHandler(usecase interfaces.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		usecase: usecase,
	}
}

func (ad *AdminHandler) AdminLogin(c *gin.Context) { // login handler for the admin

	// var adminDetails models.AdminLogin
	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	Tokens, err := ad.usecase.AdminLogin(ctx, adminDetails)
	if err != nil {
		errRes := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	c.Set("Access", Tokens.AccessToken)
	c.Set("Refresh", Tokens.RefreshToken)

	successRes := response.Response{Data: Tokens, Error: nil}
	c.JSON(http.StatusOK, successRes)

}

func (ad *AdminHandler) CreateNewAdmin(c *gin.Context) {

	var adminDetails domain.Admin
	if err := c.BindJSON(&adminDetails); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	err := ad.usecase.CreateNewAdmin(ctx, adminDetails)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//give response
	successRes := response.Response{Data: "successfully created new admin", Error: nil}
	c.JSON(http.StatusCreated, successRes)
}
