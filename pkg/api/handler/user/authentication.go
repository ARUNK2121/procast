package userhandler

import (
	"net/http"

	interfaces "github.com/ARUNK2121/procast/pkg/usecase/user/interface"
	"github.com/ARUNK2121/procast/pkg/utils/models"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	Usecase interfaces.AuthenticationUsecase
}

func NewAuthenticationHandler(use interfaces.AuthenticationUsecase) *AuthenticationHandler {
	return &AuthenticationHandler{
		Usecase: use,
	}
}

func (a *AuthenticationHandler) UserSignup(c *gin.Context) {

	var model models.UserSignup
	if err := c.BindJSON(&model); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := a.Usecase.UserSignup(model); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.Response{Data: "signup completed successfully", Error: nil}
	c.JSON(http.StatusCreated, res)
}

func (a *AuthenticationHandler) Login(c *gin.Context) {

	var model models.Login
	if err := c.BindJSON(&model); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := a.Usecase.Login(model)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.Response{Data: token, Error: nil}
	c.JSON(http.StatusCreated, res)
}
