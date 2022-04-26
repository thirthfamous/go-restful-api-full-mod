package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"thirthfamous/golang-restful-api-clean-architecture/helper"
	"thirthfamous/golang-restful-api-clean-architecture/model/domain"
	"thirthfamous/golang-restful-api-clean-architecture/model/web"
	"thirthfamous/golang-restful-api-clean-architecture/service"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(ctx *gin.Context) {
	input := web.CategoryCreateRequest{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helper.WriteToResponseBody(http.StatusBadRequest, ctx, err.Error())
		return
	}
	// Create Category
	categoryResponse, err := controller.CategoryService.Create(ctx, input)

	if err != nil {
		helper.WriteToResponseBody(http.StatusBadRequest, ctx, err.Error())
		return
	}

	helper.WriteToResponseBody(http.StatusOK, ctx, categoryResponse)
}
func (controller *CategoryControllerImpl) Update(ctx *gin.Context) {
	input := web.CategoryUpdateRequest{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		helper.WriteToResponseBody(http.StatusBadRequest, ctx, err.Error())
		return
	}
	categoryId, erro := strconv.Atoi(ctx.Param("categoryId"))
	helper.PanicIfError(erro)
	input.Id = categoryId
	// Create Category
	categoryResponse, err := controller.CategoryService.Update(ctx, input)

	if err != nil {
		helper.WriteToResponseBody(http.StatusBadRequest, ctx, err.Error())
		return
	}

	helper.WriteToResponseBody(http.StatusOK, ctx, categoryResponse)
}
func (controller *CategoryControllerImpl) Delete(ctx *gin.Context) {
	categoryId, erro := strconv.Atoi(ctx.Param("categoryId"))
	helper.PanicIfError(erro)

	err := controller.CategoryService.Delete(ctx, categoryId)

	if err != nil {
		if err.Error() == "NOT FOUND" {
			helper.WriteToResponseBody(http.StatusNotFound, ctx, err.Error())
			return
		}
		helper.WriteToResponseBody(http.StatusBadRequest, ctx, err.Error())
		return
	}

	helper.WriteToResponseBody(http.StatusOK, ctx, domain.Category{})
}
func (controller *CategoryControllerImpl) FindById(ctx *gin.Context) {
	categoryId, erro := strconv.Atoi(ctx.Param("categoryId"))
	helper.PanicIfError(erro)

	result, err := controller.CategoryService.FindById(ctx, categoryId)
	fmt.Println(err)

	if err != nil {
		if err.Error() == "NOT FOUND" {
			helper.WriteToResponseBody(http.StatusNotFound, ctx, err.Error())
			return
		}
		helper.WriteToResponseBody(http.StatusBadRequest, ctx, err.Error())
		return
	}

	helper.WriteToResponseBody(http.StatusOK, ctx, result)
}
func (controller *CategoryControllerImpl) FindAll(ctx *gin.Context) {

	// Get All Category
	categoryResponse, err := controller.CategoryService.FindAll(ctx)
	if err != nil {
		if err.Error() == "NOT FOUND" {
			helper.WriteToResponseBody(http.StatusNotFound, ctx, err.Error())
			return
		}
		helper.WriteToResponseBody(http.StatusInternalServerError, ctx, err.Error())
		return
	}

	helper.WriteToResponseBody(http.StatusOK, ctx, categoryResponse)
}
