package handler

import (
	"net/http"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUsecase port.CategoryUsecasePort
}

func NewCategoryHandler(categoryUsecase port.CategoryUsecasePort) *CategoryHandler {
	return &CategoryHandler{
		categoryUsecase: categoryUsecase,
	}
}

func (ch *CategoryHandler) Register(router *gin.RouterGroup) {
	router.POST("/", ch.CreateCategory)
	router.GET("/", ch.ListCategories)
	router.GET("/:id", ch.GetCategory)
	router.PUT("/:id", ch.UpdateCategory)
	router.DELETE("/:id", ch.DeleteCategory)
}

func (ch *CategoryHandler) GroupRouterPattern() string {
	return "/api/v1/categories"
}

func (ch *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	category := req.ToEntity()

	err := ch.categoryUsecase.Create(category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.NewCategoryResponse(category))
}

func (ch *CategoryHandler) GetCategory(ctx *gin.Context) {
	var req dto.GetCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	category, err := ch.categoryUsecase.GetByID(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.NewCategoryResponse(category))
}

func (ch *CategoryHandler) ListCategories(ctx *gin.Context) {
	name := ctx.Query("name")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid page parameter"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: "invalid limit parameter"})
		return
	}

	categories, total, err := ch.categoryUsecase.List(name, pageInt, limitInt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.NewCategoriesPaginatedResponse(categories, total, pageInt, limitInt))
}

func (ch *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var req dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	category := req.ToEntity(id)

	err = ch.categoryUsecase.Update(category)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.NewCategoryResponse(category))
}

func (ch *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	var req dto.DeleteCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	err := ch.categoryUsecase.Delete(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
