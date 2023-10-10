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

type ServiceHandler struct {
	usecase interfaces.ServiceUsecase
}

func NewServiceHandler(use interfaces.ServiceUsecase) *ServiceHandler {
	return &ServiceHandler{
		usecase: use,
	}
}

// @Summary		Add New Service
// @Description	This handler will create a new service under an existing category
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AddServicesToACategory	true	"Add service"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/services [post]
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

// @Summary		Get Services
// @Description	A call to this path along with the category id as parameter will result in the listing of all active services under that category
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			category_id	query		string	true	"category-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/services [get]
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

// @Summary		Delete Services
// @Description	A call to this path along with the service id as parameter will result in the deletion of that service
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			service_id	query		string	true	"service-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/services [delete]
func (s *ServiceHandler) DeleteService(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = s.usecase.DeleteService(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully deleted category", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Reactivate service
// @Description	A call to this path along with the service id as parameter will result in the re activation  of that service
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			service_id	query		string	true	"service-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/services [patch]
func (s *ServiceHandler) ReActivateService(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = s.usecase.ReActivateService(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully Activated service", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Committed Works
// @Description	A call to this path  will result in the listing of all committed works that havent been completed
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/work/committed [get]
func (s *ServiceHandler) ListCommittedWorks(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	//call usecase get array
	works, err := s.usecase.ListCommittedWorks(ctx)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: works, Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Completed Works
// @Description	A call to this path  will result in the listing of all completed works
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/work/completed [get]
func (s *ServiceHandler) ListCompletedWorks(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	//call usecase get array
	works, err := s.usecase.ListCompletedWorks(ctx)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: works, Error: nil}
	c.JSON(http.StatusOK, successRes)
}
