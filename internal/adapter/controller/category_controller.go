package controller

import (
	"net/http"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/response"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryUsecase port.CategoryUsecasePort
}

func NewCategoryController(categoryUsecase port.CategoryUsecasePort) *CategoryController {
	return &CategoryController{
		categoryUsecase: categoryUsecase,
	}
}

func (cc *CategoryController) Register(router *gin.RouterGroup) {
	router.POST("/", cc.CreateCategory)
	router.GET("/", cc.ListCategories)
	router.GET("/:id", cc.GetCategory)
	router.PUT("/:id", cc.UpdateCategory)
	router.DELETE("/:id", cc.DeleteCategory)
}

func (cc *CategoryController) GroupRouterPattern() string {
	return "/api/v1/categories"
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var req request.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	category := req.ToEntity()

	err := cc.categoryUsecase.Create(category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response.NewCategoryResponse(category))
}

func (cc *CategoryController) GetCategory(ctx *gin.Context) {
	var req request.GetCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	category, err := cc.categoryUsecase.GetByID(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.NewCategoryResponse(category))
}

func (cc *CategoryController) ListCategories(ctx *gin.Context) {
	name := ctx.Query("name")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: "invalid page parameter"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: "invalid limit parameter"})
		return
	}

	categories, total, err := cc.categoryUsecase.List(name, pageInt, limitInt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.NewCategoriesPaginatedResponse(categories, total, pageInt, limitInt))
}

func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	var req request.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	category := req.ToEntity(id)

	err = cc.categoryUsecase.Update(category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.NewCategoryResponse(category))
}

func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	var req request.DeleteCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err := cc.categoryUsecase.Delete(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
