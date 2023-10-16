package providerhandler

import (
	"net/http"
	"strconv"

	interfaces "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	usecase interfaces.ProfileUsecase
}

func NewProfileHandler(use interfaces.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		usecase: use,
	}
}

// @Summary		Add Service
// @Description	providers can add a service to their providable services
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			user_id	query		string	true	"user-id"
// @Param			service_id	query		string	true	"service-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/profile/service [post]
func (p *ProfileHandler) AddService(c *gin.Context) {

	user_id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	service_id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.AddService(user_id, service_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: "successfully added service", Error: nil}
	c.JSON(http.StatusOK, res)

}

// @Summary		Delete Service
// @Description	providers can remove a service from their providable services
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			user_id	 query		string	true	"user-id"
// @Param			service_id	query		string	true	"service-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/profile/service [delete]
func (p *ProfileHandler) DeleteService(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	service_id, err := strconv.Atoi(c.Query("service_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.DeleteService(user_id, service_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.Response{Data: "successfully removed service", Error: nil}
	c.JSON(http.StatusOK, res)
}

// @Summary		Add Preferred Location
// @Description	providers can add a location to their preferred location services
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			user_id	query		string	true	"user-id"
// @Param			district_id	query		string	true	"district-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/profile/location [post]
func (p *ProfileHandler) AddPreferredWorkingLocation(c *gin.Context) {

	user_id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	district_id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.AddPreferredWorkingLocation(user_id, district_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: "successfully added location to preferred locations", Error: nil}
	c.JSON(http.StatusOK, res)

}

// @Summary		Remove Preferred Location
// @Description	providers can remove a location from their preferred location services
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			user_id	query		string	true	"user-id"
// @Param			district_id	query		string	true	"district-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/profile/location [delete]
func (p *ProfileHandler) RemovePreferredLocation(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	district_id, err := strconv.Atoi(c.Query("district_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err = p.usecase.RemovePreferredLocation(user_id, district_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.Response{Data: "successfully removed location from preferred", Error: nil}
	c.JSON(http.StatusOK, res)
}

// @Summary		Get Selected Services
// @Description	providers can list their selected services
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			user_id	query		string	true	"user-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/profile/service [get]
func (p *ProfileHandler) GetSelectedServices(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	services, err := p.usecase.GetMyServices(user_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.Response{Data: services, Error: nil}
	c.JSON(http.StatusOK, res)
}

// @Summary		Get All Preferred Locations
// @Description	providers can list their selected services
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			user_id	query		string	true	"user-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/profile/service [get]
func (p *ProfileHandler) GetAllPreferredLocations(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	locations, err := p.usecase.GetAllPreferredLocations(user_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.Response{Data: locations, Error: nil}
	c.JSON(http.StatusOK, res)
}

// @Summary		Get Provider Details For User
// @Description	 get details of providers when user clicks on provider profile
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			pro-id	query		string	true	"pro-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/user/provider/:pro-id [get]
func (p *ProfileHandler) GetDetailsOfProviders(c *gin.Context) {
	pro_id, err := strconv.Atoi(c.Param("pro-id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	details, err := p.usecase.GetDetailsOfProviders(pro_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.Response{Data: details, Error: nil}
	c.JSON(http.StatusOK, res)
}
