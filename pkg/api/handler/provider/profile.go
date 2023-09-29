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
