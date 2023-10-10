package adminhandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ARUNK2121/procast/pkg/usecase/admin/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/models"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type RegionHandler struct {
	usecase interfaces.RegionUsecase
}

func NewRegionHandler(use interfaces.RegionUsecase) *RegionHandler {
	return &RegionHandler{
		usecase: use,
	}
}

// @Summary		Add New State
// @Description	This handler creates a new state which means that the company is now expanded its size to a new state
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AddNewState	true	"Create state"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/region/state [post]
func (r *RegionHandler) AddNewState(c *gin.Context) {
	var region models.AddNewState
	err := c.BindJSON(&region)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err = r.usecase.AddNewState(ctx, region.State)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.Response{Data: "successfully added new state", Error: nil}
	c.JSON(http.StatusCreated, successRes)
}

// @Summary		List States
// @Description	A call to this path will list all the states that has been already added by admins before or in other words this handler will displays the span of the company in states
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/region/state [get]
func (r *RegionHandler) GetStates(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	//call usecase get array
	states, err := r.usecase.GetStates(ctx)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: states, Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Delete State
// @Description	A call to this path along with the district id as parameter will result in the deletion of that state
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"state-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/region/state [delete]
func (r *RegionHandler) DeleteState(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("state_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = r.usecase.DeleteState(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully made state inactive", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Reactivate State
// @Description	A call to this path along with the distroct id as parameter will result in the re activation  of that state
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"state-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/region/state  [patch]
func (r *RegionHandler) ReActivateState(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("state_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = r.usecase.ReActivateState(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully Activated state", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Add New District
// @Description	Districts are the smallest unit in which procast works,This handler creates a new district under a state which is already exists
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AddNewDistrict	true	"Add District"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/region/district [post]
func (r *RegionHandler) AddNewDistrict(c *gin.Context) {
	var region models.AddNewDistrict
	err := c.BindJSON(&region)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err = r.usecase.AddNewDistrict(ctx, region)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.Response{Data: "successfully added new District", Error: nil}
	c.JSON(http.StatusCreated, successRes)
}

// @Summary		Get Districts
// @Description	A call to this path along with the state id as parameter will result in the listing of all active districts under that state
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			state_id	query		string	true	"state-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/region/district [get]
func (r *RegionHandler) GetDistrictsFromState(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("state_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//call usecase get array
	districts, err := r.usecase.GetDistrictsFromState(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: districts, Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Delete district
// @Description	A call to this path along with the district id as parameter will result in the deletion of that district
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			district_id	query		string	true	"district-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/region/district [delete]
func (r *RegionHandler) DeleteDistrictFromState(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = r.usecase.DeleteDistrictFromState(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully deleted district", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Reactivate district
// @Description	A call to this path along with the district id as parameter will result in the re activation  of that district
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			district_id	query		string	true	"district-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/region/district  [patch]
func (r *RegionHandler) ReActivateDistrict(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = r.usecase.ReActivateDistrict(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully Activated district", Error: nil}
	c.JSON(http.StatusOK, successRes)
}
