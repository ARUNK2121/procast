package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ARUNK2121/procast/pkg/usecase/interfaces"
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
