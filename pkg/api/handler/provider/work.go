package providerhandler

import (
	"net/http"
	"strconv"

	interfaces "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type WorkHandler struct {
	usecase interfaces.WorkUseCase
}

func NewWorkHandler(use interfaces.WorkUseCase) *WorkHandler {
	return &WorkHandler{
		usecase: use,
	}
}

func (w *WorkHandler) GetAllLeads(c *gin.Context) {
	provider_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	leads, err := w.usecase.GetAllLeads(provider_id, page)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: leads, Error: nil}
	c.JSON(http.StatusOK, res)
}

func (w *WorkHandler) ViewLeads(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	lead, err := w.usecase.ViewLeads(work_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: lead, Error: nil}
	c.JSON(http.StatusOK, res)
}

func (w *WorkHandler) PlaceBid(c *gin.Context) {
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	Pro_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var model models.PlaceBid

	if err := c.BindJSON(&model); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	model.WorkID = work_id
	model.ProID = Pro_id

	err = w.usecase.PlaceBid(model)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: "bid has been placed successfully", Error: nil}
	c.JSON(http.StatusOK, res)
}
