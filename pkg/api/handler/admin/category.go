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

type CategoryHandler struct {
	usecase interfaces.CategoryUsecase
}

func NewCategoryHandler(use interfaces.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		usecase: use,
	}
}

// @Summary		Create Category
// @Description	This handler creates a new category of works and severals services will be there under each categories
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.CreateCategory	true	"Create Category"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/category [post]
func (ad *CategoryHandler) CreateCategory(c *gin.Context) {

	var category models.CreateCategory
	err := c.BindJSON(&category)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err = ad.usecase.CreateCategory(ctx, category.Category)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	successRes := response.Response{Data: "successfully created category", Error: nil}
	c.JSON(http.StatusCreated, successRes)
}

// @Summary		List Categories
// @Description	A call to this path will list all the catrgories that has been already added by admins before
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/category [get]
func (ad *CategoryHandler) ListCategories(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	//call usecase get array
	categories, err := ad.usecase.ListCategories(ctx)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: categories, Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Delete Category
// @Description	A call to this path along with the category id as parameter will result in the deletion of that category
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"category-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category [delete]
func (ad *CategoryHandler) DeleteCategory(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = ad.usecase.DeleteCategory(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully deleted category", Error: nil}
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Reactivate Category
// @Description	A call to this path along with the category id as parameter will result in the re activation  of that category
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"category-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category [patch]
func (ad *CategoryHandler) ReActivateCategory(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//call usecase get array
	err = ad.usecase.ReActivateCategory(ctx, id)
	if err != nil {
		res := response.Response{Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//give array
	successRes := response.Response{Data: "successfully Activated category", Error: nil}
	c.JSON(http.StatusOK, successRes)
}
