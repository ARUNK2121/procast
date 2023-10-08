package userhandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ARUNK2121/procast/pkg/domain"
	interfaces "github.com/ARUNK2121/procast/pkg/usecase/user/interface"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type WorkHandler struct {
	usecase interfaces.WorkUsecase
}

func NewWorkHandler(use interfaces.WorkUsecase) *WorkHandler {
	return &WorkHandler{
		usecase: use,
	}
}

func (p *WorkHandler) ListNewOpening(c *gin.Context) {
	var model domain.Work
	if err := c.BindJSON(&model); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("model", model)

	err := p.usecase.ListNewOpening(model)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.Response{Data: "successfully listed opening", Error: nil}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) GetAllListedWorks(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	works, err := p.usecase.GetAllListedWorks(id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.Response{Data: works, Error: nil}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) ListAllCompletedWorks(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	works, err := p.usecase.ListAllCompletedWorks(id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.Response{Data: works, Error: nil}
	c.JSON(http.StatusCreated, res)

}

func (p *WorkHandler) ListAllOngoingWorks(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	works, err := p.usecase.ListAllOngoingWorks(id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.Response{Data: works, Error: nil}
	c.JSON(http.StatusCreated, res)

}
