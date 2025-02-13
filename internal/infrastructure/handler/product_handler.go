package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type ProductHandler struct {
	controller *controller.ProductController
}

func NewProductHandler(controller *controller.ProductController) *ProductHandler {
	return &ProductHandler{controller: controller}
}

func (h *ProductHandler) GroupRouterPattern() string {
	return "/products"
}

func (h *ProductHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListProducts)
	router.POST("/", h.CreateProduct)
	router.GET("/:id", h.GetProduct)
	router.PUT("/:id", h.UpdateProduct)
	router.DELETE("/:id", h.DeleteProduct)
}

// ListProducts lista os produtos
//
//	@Summary		Listar produtos
//	@Description	Retorna uma lista paginada de produtos
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int		false	"Número da página"	default(1)
//	@Param			limit		query		int		false	"Itens por página"	default(10)
//	@Param			name		query		string	false	"Filtrar por nome"
//	@Param			category_id	query		int		false	"Filtrar por categoria"
//	@Success		200			{object}	dto.PaginatedResponse
//	@Failure		400			{object}	dto.ErrorResponse
//	@Failure		500			{object}	dto.ErrorResponse
//	@Router			/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	var req dto.ProductListRequest

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	categoryID, _ := strconv.ParseUint(c.DefaultQuery("category_id", "0"), 10, 64)

	req.Name = c.Query("name")
	req.CategoryID = categoryID
	req.Page = page
	req.Limit = limit

	response, err := h.controller.ListProducts(c.Request.Context(), req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateProduct cria um novo produto
//
//	@Summary		Criar produto
//	@Description	Cria um novo produto
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			product	body		dto.ProductRequest	true	"Dados do produto"
//	@Success		201		{object}	dto.ProductResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	response, err := h.controller.CreateProduct(c.Request.Context(), req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetProduct busca um produto pelo ID
//
//	@Summary		Buscar produto
//	@Description	Busca um produto pelo ID
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID do produto"
//	@Success		200	{object}	dto.ProductResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	response, err := h.controller.GetProduct(c.Request.Context(), id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProduct atualiza um produto
//
//	@Summary		Atualizar produto
//	@Description	Atualiza um produto existente
//	@Tags			produtos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"ID do produto"
//	@Param			product	body		dto.ProductRequest	true	"Dados do produto"
//	@Success		200		{object}	dto.ProductResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Failure		500		{object}	dto.ErrorResponse
//	@Router			/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var req dto.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	response, err := h.controller.UpdateProduct(c.Request.Context(), id, req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteProduct deleta um produto
//
//	@Summary		Deletar produto
//	@Description	Remove um produto existente
//	@Tags			produtos
//	@Produce		json
//	@Param			id	path		int	true	"ID do produto"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	if err := h.controller.DeleteProduct(c.Request.Context(), id); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
