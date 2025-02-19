package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type ProductHandler struct {
	controller *controller.ProductController
}

type ListProductQueryRequest struct {
	Name       string `form:"name" example:"Product A"`
	CategoryID uint64 `form:"category_id" example:"1"`
	Page       int    `form:"page,default=1" example:"1"`
	Limit      int    `form:"limit,default=10" example:"10"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100" example:"Product A"`
	Description string  `json:"description" validate:"max=500" example:"Product A description"`
	Price       float64 `json:"price" validate:"required,gt=0" example:"99.99"`
	CategoryID  uint64  `json:"category_id" validate:"required,gt=0" example:"1"`
}

// func (p *CreateProductRequest) Validate() error {
// 	return GetValidator().Struct(p)
// }

type GetProductUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateProductUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100" example:"Product A"`
	Description string  `json:"description" validate:"max=500" example:"Product A description"`
	Price       float64 `json:"price" validate:"required,gt=0" example:"99.99"`
	CategoryID  uint64  `json:"category_id" validate:"required,gt=0" example:"1"`
}

type DeleteProductUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

func NewProductHandler(controller *controller.ProductController) *ProductHandler {
	return &ProductHandler{controller: controller}
}

func (h *ProductHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListProducts)
	router.POST("/", h.CreateProduct)
	router.GET("/:id", h.GetProduct)
	router.PUT("/:id", h.UpdateProduct)
	router.DELETE("/:id", h.DeleteProduct)
}

// ListProducts godoc
//
//	@Summary		List products
//	@Description	List all products
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int										false	"Page number"		default(1)
//	@Param			limit		query		int										false	"Items per page"	default(10)
//	@Param			name		query		string									false	"Filter by name"
//	@Param			category_id	query		int										false	"Filter by category ID"
//	@Success		200			{object}	presenter.ProductJsonPaginatedResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorResponse				"Bad Request"
//	@Failure		500			{object}	middleware.ErrorResponse				"Internal Server Error"
//	@Router			/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	var req ListProductQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidQueryParams))
		return
	}

	input := dto.ListProductsInput{
		Name:       c.Query("name"),
		CategoryID: req.CategoryID,
		Page:       req.Page,
		Limit:      req.Limit,
	}

	if c.GetHeader("Accept") == "application/xml" {
		h.controller.Presenter = presenter.NewProductXmlPresenter(c)
	} else {
		h.controller.Presenter = presenter.NewProductJsonPresenter(c)
	}

	err := h.controller.ListProducts(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// CreateProduct godoc
//
//	@Summary		Create product
//	@Description	Creates a new product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		CreateProductRequest			true	"Product data"
//	@Success		201		{object}	presenter.ProductJsonResponse	"Created"
//	@Failure		400		{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		500		{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.CreateProductInput{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	if c.GetHeader("Accept") == "application/xml" {
		h.controller.Presenter = presenter.NewProductXmlPresenter(c)
	} else {
		h.controller.Presenter = presenter.NewProductJsonPresenter(c)
	}

	err := h.controller.CreateProduct(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// GetProduct godoc
//
//	@Summary		Get product
//	@Description	Search for a product by ID
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"Product ID"
//	@Success		200	{object}	presenter.ProductJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse		"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	var req GetProductUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetProductInput{
		ID: req.ID,
	}

	if c.GetHeader("Accept") == "application/xml" {
		h.controller.Presenter = presenter.NewProductXmlPresenter(c)
	} else {
		h.controller.Presenter = presenter.NewProductJsonPresenter(c)
	}

	err := h.controller.GetProduct(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// UpdateProduct godoc
//
//	@Summary		Update product
//	@Description	Update an existing product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"Product ID"
//	@Param			product	body		UpdateProductRequest			true	"Product data"
//	@Success		200		{object}	presenter.ProductJsonResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		404		{object}	middleware.ErrorResponse		"Not Found"
//	@Failure		500		{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var reqUri UpdateProductUriRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdateProductInput{
		ID:          reqUri.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	if c.GetHeader("Accept") == "application/xml" {
		h.controller.Presenter = presenter.NewProductXmlPresenter(c)
	} else {
		h.controller.Presenter = presenter.NewProductJsonPresenter(c)
	}

	err := h.controller.UpdateProduct(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// DeleteProduct godoc
//
//	@Summary		Delete product
//	@Description	Deletes a product by ID
//	@Tags			products
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	var reqUri DeleteProductUriRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteProductInput{
		ID: reqUri.ID,
	}

	if c.GetHeader("Accept") ==  "application/xml" {
		h.controller.Presenter = presenter.NewProductXmlPresenter(c)
	} else {
		h.controller.Presenter = presenter.NewProductJsonPresenter(c)
	}

	if err := h.controller.DeleteProduct(c.Request.Context(), input); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
