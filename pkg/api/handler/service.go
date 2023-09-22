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

type ServiceHandler struct {
	usecase interfaces.ServiceUsecase
}

func NewServiceHandler(use interfaces.ServiceUsecase) *ServiceHandler {
	return &ServiceHandler{
		usecase: use,
	}
}

func (s *ServiceHandler) AddServicesToACategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	var service models.AddServicesToACategory
	err := c.BindJSON(&service)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = s.usecase.AddServicesToACategory(ctx, service)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.Response{Data: "successfully created service", Error: nil}
	c.JSON(http.StatusCreated, successRes)
}

func (s *ServiceHandler) GetServicesInACategory(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//call usecase get array
	services, err := s.usecase.GetServicesInACategory(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: services, Error: nil}
	c.JSON(http.StatusOK, successRes)
}
