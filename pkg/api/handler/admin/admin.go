package adminhandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/usecase/admin/interfaces"
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

// @Summary		Admin Login
// @Description	This is the Login handler for procast admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AdminLogin	true	"Admin login details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/login [post]
func (ad *AdminHandler) AdminLogin(c *gin.Context) {

	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	Token, err := ad.usecase.AdminLogin(ctx, adminDetails)
	if err != nil {
		errRes := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.Response{Data: Token, Error: nil}
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Create New Admin
// @Description	Using this end point the super admin can create new admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		domain.Admin	true	"Admin Details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/panel [post]
func (ad *AdminHandler) CreateNewAdmin(c *gin.Context) {

	var adminDetails domain.Admin
	if err := c.BindJSON(&adminDetails); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err := ad.usecase.CreateNewAdmin(ctx, adminDetails)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.Response{Data: "successfully created new admin", Error: nil}
	c.JSON(http.StatusCreated, successRes)
}

// @Summary		Delete Admin
// @Description	using this handler super admins can delete other admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"admin-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/panel [delete]
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
