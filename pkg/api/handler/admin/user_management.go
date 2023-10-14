package adminhandler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ARUNK2121/procast/pkg/usecase/admin/interfaces"
	"github.com/ARUNK2121/procast/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserManagementHandler struct {
	usecase interfaces.UserManagementUsecase
}

func NewUserManagementHandler(usecase interfaces.UserManagementUsecase) *UserManagementHandler {
	return &UserManagementHandler{
		usecase: usecase,
	}
}

// @Summary		Get Providers
// @Description	A call to this path  will result in the listing of all providers that have registered in procast
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/provider [get]
func (u *UserManagementHandler) GetProviders(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	//call usecase get array
	providers, err := u.usecase.GetProviders(ctx)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: providers, Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Verify Provider
// @Description	A call to this path along with the provider  id as parameter will result in the verification of the provider
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"provider-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/provider [patch]
func (u *UserManagementHandler) MakeProviderVerified(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = u.usecase.MakeProviderVerified(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully Verified provider", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Revoke Verification From Provider
// @Description	A call to this path along with the provider  id as parameter will result in the revoking verification of the provider
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"provider-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/provider [delete]
func (u *UserManagementHandler) RevokeVerification(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = u.usecase.RevokeVerification(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "revoked verification of provider", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Users
// @Description	A call to this path  will result in the listing of all users of application
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/user [get]
func (u *UserManagementHandler) GetUsers(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	//call usecase get array
	users, err := u.usecase.GetUsers(ctx)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	fmt.Println("users", users)

	//give array
	successRes := response.Response{Data: users, Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Block User
// @Description	A call to this path along with the user  id as parameter will result in the blocking of the user
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"user-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/user [delete]
func (u *UserManagementHandler) BlockUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = u.usecase.BlockUser(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "blocked user", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		UnBlock User
// @Description	A call to this path along with the user  id as parameter will result in the unblocking of user
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"provider-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/user [patch]
func (u *UserManagementHandler) UnBlockUser(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = u.usecase.UnBlockUser(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "unblocked user", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get All Pending Verifications
// @Description	A call to this path  will result in the listing of all providers request for verification which has not been accepted yet
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/verify [get]
func (u *UserManagementHandler) GetAllPendingVerifications(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	//call usecase get array
	verifications, err := u.usecase.GetAllPendingVerifications(ctx)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: verifications, Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Verification Request
// @Description	A call to this path along with the provider id will result in the displaying of the verification request of a provider
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			pro_id	query		string	true	"pro-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/verify/request [get]
func (u *UserManagementHandler) ViewVerificationRequest(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("pro_id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	//call usecase get array
	request, err := u.usecase.ViewVerificationRequest(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: request, Error: nil}
	c.JSON(http.StatusOK, successRes)
}
