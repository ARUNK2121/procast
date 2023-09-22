package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

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

// try timeout by time.sleep
func (ad *AdminHandler) AdminLogin(c *gin.Context) { // login handler for the admin

	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	time.Sleep(20 * time.Second)

	err := ad.usecase.CreateNewAdmin(ctx, adminDetails)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.Response{Data: "successfully created new admin", Error: nil}
	c.JSON(http.StatusCreated, successRes)
}

// delete admins
func (ad *AdminHandler) DeleteAdmin(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err = ad.usecase.DeleteAdmin(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.Response{Data: "successfully deleted admin", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// admin add category
//admin delete category
//admin add profession
//admin add state
//admin add district

//admin verify provider
//admin notifications
//take charge of verifications
//pending verifications
//admin block provider(making verified false)
//admin change profile image
//list providers by category
//sort providers by criteria
//list scheduled works
