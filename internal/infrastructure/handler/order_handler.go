package handler

import (
	"fmt"
	"strconv"
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

type CreateOrderRequest struct {
	CustomerID uint64 `json:"customer_id" binding:"required" example:"1"`
}

func (p *CreateOrderRequest) Validate() error {
	return GetValidator().Struct(p)
}

// type ListOrderRequest struct {
// 	Name       string `json:"name" validate:"required,min=3,max=100" example:"Produto"`
// 	CategoryID uint64 `json:"category_id" example:"1"`
// 	Page       int    `json:"page" validate:"required,gte=1" example:"1"`
// 	Limit      int    `json:"limit" validate:"required,gte=1,lte=100" example:"10"`
// }

// func (p *ListOrderRequest) Validate() error {
// 	return GetValidator().Struct(p)
// }

func NewOrderHandler(controller *controller.OrderController) *OrderHandler {
	return &OrderHandler{controller: controller}
}

func (h *OrderHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListOrders)
	router.POST("/", h.CreateOrder)
	// router.GET("/:id", h.GetOrder)
	// router.PUT("/:id", h.UpdateOrder)
	// router.DELETE("/:id", h.DeleteOrder)
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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	customerID, _ := strconv.ParseUint(c.DefaultQuery("customer_id", "0"), 10, 64)
	status := strings.ToUpper(c.DefaultQuery("status", ""))

	if status != "" && !entity.IsValidOrderStatus(status) {
		fmt.Println("status", status)
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.ListOrdersInput{
		CustomerID: customerID,
		Status:     status,
		Page:       page,
		Limit:      limit,
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
//	@Param			order	body		CreateOrderRequest					true	"Order data"
//	@Success		201		{object}	presenter.OrderJsonResponse	"Created"
//	@Failure		400		{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		500		{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	// validate request
	if err := req.Validate(); err != nil {
		_ = c.Error(domain.NewInvalidInputError(err.Error()))
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
//	@Param			id	path		int								true	"Order ID"
//	@Success		200	{object}	presenter.OrderJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse		"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/orders/{id} [get]
// func (h *OrderHandler) GetOrder(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
// 		return
// 	}

// 	input := dto.GetOrderInput{
// 		ID:     id,
// 		Writer: c,
// 	}

// 	err = h.controller.GetOrder(c.Request.Context(), input)
// 	if err != nil {
// 		_ = c.Error(err)
// 		return
// 	}
// }

// // UpdateOrder godoc
// //
// //	@Summary		Update order
// //	@Description	Update an existing order
// //	@Tags			orders
// //	@Accept			json
// //	@Produce		json
// //	@Param			id		path		int								true	"Order ID"
// //	@Param			order	body		OrderRequest					true	"Order data"
// //	@Success		200		{object}	presenter.OrderJsonResponse	"OK"
// //	@Failure		400		{object}	middleware.ErrorResponse		"Bad Request"
// //	@Failure		404		{object}	middleware.ErrorResponse		"Not Found"
// //	@Failure		500		{object}	middleware.ErrorResponse		"Internal Server Error"
// //	@Router			/orders/{id} [put]
// func (h *OrderHandler) UpdateOrder(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
// 		return
// 	}

// 	var req OrderRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
// 		return
// 	}

// 	if err := req.Validate(); err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(err.Error()))
// 		return
// 	}

// 	input := dto.UpdateOrderInput{
// 		ID:          id,
// 		Name:        req.Name,
// 		Description: req.Description,
// 		Price:       req.Price,
// 		CategoryID:  req.CategoryID,
// 		Writer:      c,
// 	}

// 	err = h.controller.UpdateOrder(c.Request.Context(), input)
// 	if err != nil {
// 		_ = c.Error(err)
// 		return
// 	}
// }

// // DeleteOrder godoc
// //
// //	@Summary		Delete order
// //	@Description	Deletes a order by ID
// //	@Tags			orders
// //	@Produce		json
// //	@Param			id	path		int	true	"Order ID"
// //	@Success		204	{object}	nil
// //	@Failure		400	{object}	middleware.ErrorResponse	"Bad Request"
// //	@Failure		404	{object}	middleware.ErrorResponse	"Not Found"
// //	@Failure		500	{object}	middleware.ErrorResponse	"Internal Server Error"
// //	@Router			/orders/{id} [delete]
// func (h *OrderHandler) DeleteOrder(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
// 		return
// 	}

// 	input := dto.DeleteOrderInput{
// 		ID:     id,
// 		Writer: c,
// 	}

// 	if err := h.controller.DeleteOrder(c.Request.Context(), input); err != nil {
// 		_ = c.Error(err)
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }
