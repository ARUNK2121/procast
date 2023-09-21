package handler

import (
	"context"
	"net/http"
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
