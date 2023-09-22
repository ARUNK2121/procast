package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ARUNK2121/procast/pkg/usecase/interfaces"
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
