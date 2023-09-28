package providerhandler

import (
	"net/http"

	interfaces "github.com/ARUNK2121/procast/pkg/usecase/provider/interface"
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

func (a *AuthenticationHandler) Register(c *gin.Context) {
	//take input into model
	name := c.Request.FormValue("name")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")
	phone := c.Request.FormValue("phone")
	image, err := c.FormFile("document")
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var model models.ProviderRegister

	model.Name = name
	model.Email = email
	model.Password = password
	model.RePassword = repassword
	model.Phone = phone
	model.Document = image
	//call usecase

	if err := a.Usecase.Register(model); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.Response{Data: "your request will be under inspection of admins", Error: nil}
	c.JSON(http.StatusCreated, res)
}

func (a *AuthenticationHandler) Login(c *gin.Context) {
	//take input into model
	var model models.ProLogin
	if err := c.BindJSON(&model); err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//call usecase
	token, err := a.Usecase.Login(model)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//return result
	res := response.Response{Data: token, Error: nil}
	c.JSON(http.StatusOK, res)

}
