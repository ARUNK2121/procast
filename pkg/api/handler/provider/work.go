package providerhandler

import (
	"fmt"
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

// @Summary		Get All Leads For The Provider
// @Description	 get details of all possible leads that provider can bid on them
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			pro-id	query		string	true	"pro-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/work/leads [get]
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

// @Summary		Get Details of A lead of provider
// @Description	 a call to this endpoint along with the work_id as parametr will get details of the lead
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/work/leads/:id [get]
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

// @Summary		Place Bid
// @Description	 A provider can place bid on an opening
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			pro-id	query		string	true	"pro-id"
// @Param			bid	body		models.PlaceBid	true	"Bid Details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/work/leads/:id/bid [post]
func (w *WorkHandler) PlaceBid(c *gin.Context) {
	fmt.Println("reaches")
	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println("workid", work_id)

	Pro_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("pro_id")

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

// @Summary		Replace Bid with New Bid
// @Description	 A provider can replace an existing bid on an opening
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			pro-id	query		string	true	"pro-id"
// @Param			bid	body		models.PlaceBid	true	"Bid Details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/work/leads/:id/bid [put]
func (w *WorkHandler) ReplaceBidWithNewBid(c *gin.Context) {
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

	err = w.usecase.ReplaceBidWithNewBid(model)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: "old bid has been Replaced successfully with new bid", Error: nil}
	c.JSON(http.StatusOK, res)
}

// @Summary		Compare All Other Bids On the Work
// @Description	 A provider can list all the bids placed on the work by various providers and then he can adjust his bid accordingly
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/work/leads/:id/compare [get]
func (w *WorkHandler) GetAllOtherBidsOnTheLeads(c *gin.Context) {

	work_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	bids, err := w.usecase.GetAllOtherBidsOnTheLeads(work_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: bids, Error: nil}
	c.JSON(http.StatusOK, res)

}

// @Summary		Get Works Of A Provider
// @Description	 An endpoint to display all the works of a provider
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			pro_id	query		string	true	"pro_id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/user/provider/:pro_id/works [get]
func (w *WorkHandler) GetWorksOfAProvider(c *gin.Context) {
	provider_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	works, err := w.usecase.GetWorksOfAProvider(provider_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: works, Error: nil}
	c.JSON(http.StatusOK, res)
}

// @Summary		Get All Ongoing Works Of A Provider
// @Description	 An endpoint to display all the ongoing works of a provider
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			pro_id	query		string	true	"pro-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/works/my-works/on-going [get]
func (w *WorkHandler) GetAllOnGoingWorks(c *gin.Context) {
	provider_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	works, err := w.usecase.GetAllOnGoingWorks(provider_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: works, Error: nil}
	c.JSON(http.StatusOK, res)
}

// @Summary		Get All Completed Works Of A Provider
// @Description	 An endpoint to display all the completed works of a provider
// @Tags			Provider
// @Accept			json
// @Produce		json
// @Param			pro_id	query		string	true	"pro-id"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/provider/works/my-works/completed [get]
func (w *WorkHandler) GetCompletedWorks(c *gin.Context) {
	provider_id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	works, err := w.usecase.GetCompletedWorks(provider_id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.Response{Data: works, Error: nil}
	c.JSON(http.StatusOK, res)
}
