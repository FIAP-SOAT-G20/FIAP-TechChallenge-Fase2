package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type OrderHandler struct {
	controller *controller.OrderController
}

type ListOrdersQueryRequest struct {
	Page       int    `form:"page,default=1" example:"1"`
	Limit      int    `form:"limit,default=10" example:"10"`
	CustomerID uint64 `form:"customer_id" example:"1" default:"0"`
	Status     string `form:"status" example:"PENDING"`
}

type CreateOrderBodyRequest struct {
	CustomerID uint64 `json:"customer_id" binding:"required" example:"1"`
}

type GetOrderUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateOrderUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateOrderBodyRequest struct {
	CustomerID uint64             `json:"customer_id" binding:"required" example:"1"`
	TotalBill  float32            `json:"total_bill" binding:"required" example:"100.00"`
	Status     entity.OrderStatus `json:"status" binding:"required" example:"PENDING"`
}

type UpdateOrderPartilBodyRequest struct {
	CustomerID uint64             `json:"customer_id" example:"1"`
	TotalBill  float32            `json:"total_bill" example:"100.00"`
	Status     entity.OrderStatus `json:"status" example:"PENDING"`
}

type DeleteOrderUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

func NewOrderHandler(controller *controller.OrderController) *OrderHandler {
	return &OrderHandler{controller: controller}
}

func (h *OrderHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListOrders)
	router.POST("/", h.CreateOrder)
	router.GET("/:id", h.GetOrder)
	router.PUT("/:id", h.UpdateOrder)
	router.PATCH("/:id", h.UpdateOrderPartial)
	router.DELETE("/:id", h.DeleteOrder)
}

// ListOrders godoc
//
//	@Summary		List orders
//	@Description	List all orders
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int										false	"Page number"		default(1)
//	@Param			limit		query		int										false	"Items per page"	default(10)
//	@Param			name		query		string									false	"Filter by name"
//	@Param			category_id	query		int										false	"Filter by category ID"
//	@Success		200			{object}	presenter.OrderJsonPaginatedResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorResponse				"Bad Request"
//	@Failure		500			{object}	middleware.ErrorResponse				"Internal Server Error"
//	@Router			/orders [get]
func (h *OrderHandler) ListOrders(c *gin.Context) {
	var req ListOrdersQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.ListOrdersInput{
		CustomerID: req.CustomerID,
		Status:     entity.OrderStatus(strings.ToUpper(req.Status)), // TODO: Validate status wit a custom validator
		Page:       req.Page,
		Limit:      req.Limit,
		Writer:     c,
	}

	err := h.controller.ListOrders(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// CreateOrder godoc
//
//	@Summary		Create order
//	@Description	Creates a new order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		CreateOrderBodyRequest		true	"Order data"
//	@Success		201		{object}	presenter.OrderJsonResponse	"Created"
//	@Failure		400		{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		500		{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderBodyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.CreateOrderInput{
		CustomerID: req.CustomerID,
		Writer:     c,
	}

	err := h.controller.CreateOrder(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// GetOrder godoc
//
//	@Summary		Get order
//	@Description	Search for a order by ID
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int							true	"Order ID"
//	@Success		200	{object}	presenter.OrderJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	var req GetOrderUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetOrderInput{
		ID:     req.ID,
		Writer: c,
	}

	err := h.controller.GetOrder(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// UpdateOrder godoc
//
//	@Summary		Update order
//	@Description	Update an existing order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Order ID"
//	@Param			order	body		UpdateOrderBodyRequest		true	"Order data"
//	@Success		200		{object}	presenter.OrderJsonResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404		{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500		{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	var reqUri UpdateOrderUriRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var req UpdateOrderBodyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdateOrderInput{
		ID:         reqUri.ID,
		CustomerID: req.CustomerID,
		TotalBill:  req.TotalBill,
		Status:     req.Status,
		Writer:     c,
	}

	err := h.controller.UpdateOrder(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// UpdateOrderPartial godoc
//
//	@Summary		Partial update order
//	@Description	Partially updates an existing order
//	@Description	The status are: **OPEN**, **CANCELLED**, **PENDING**, **RECEIVED**, **PREPARING**, **READY**, **COMPLETED**
//	@Description	## Transition of status:
//	@Description	- OPEN      -> CANCELLED || PENDING
//	@Description	- CANCELLED -> {},
//	@Description	- PENDING   -> OPEN || RECEIVED
//	@Description	- RECEIVED  -> PREPARING
//	@Description	- PREPARING -> READY
//	@Description	- READY     -> COMPLETED
//	@Description	- COMPLETED -> {}
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Order ID"
//	@Param			order	body		UpdateOrderPartilBodyRequest		true	"Order data"
//	@Success		200		{object}	presenter.OrderJsonResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404		{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500		{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/orders/{id} [patch]
func (h *OrderHandler) UpdateOrderPartial(c *gin.Context) {
	var reqUri UpdateOrderUriRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var req UpdateOrderPartilBodyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdateOrderInput{
		ID:         reqUri.ID,
		CustomerID: req.CustomerID,
		TotalBill:  req.TotalBill,
		Status:     req.Status,
		Writer:     c,
	}

	err := h.controller.UpdateOrder(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// DeleteOrder godoc
//
//	@Summary		Delete order
//	@Description	Deletes a order by ID
//	@Tags			orders
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	var req DeleteOrderUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteOrderInput{
		ID:     req.ID,
		Writer: c,
	}

	if err := h.controller.DeleteOrder(c.Request.Context(), input); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
