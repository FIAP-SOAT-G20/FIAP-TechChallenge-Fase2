package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type OrderProductHandler struct {
	controller *controller.OrderProductController
}

type ListOrderProductsQueryRequest struct {
	OrderID   uint64 `form:"order_id,default=0" example:"1"`
	ProductID uint64 `form:"product_id,default=0" example:"1"`
	Page      int    `form:"page,default=1" example:"1"`
	Limit     int    `form:"limit,default=10" example:"10"`
}

type CreateOrderProductUriRequest struct {
	OrderID   uint64 `uri:"order_id" binding:"required"`
	ProductID uint64 `uri:"product_id" binding:"required"`
}

type CreateOrderProductBodyRequest struct {
	Quantity uint32 `json:"quantity" binding:"required" example:"1"`
}

type GetOrderProductUriRequest struct {
	OrderID   uint64 `uri:"order_id" binding:"required"`
	ProductID uint64 `uri:"product_id" binding:"required"`
}

type UpdateOrderProductUriRequest struct {
	OrderID   uint64 `uri:"order_id" binding:"required"`
	ProductID uint64 `uri:"product_id" binding:"required"`
}

type UpdateOrderProductBodyRequest struct {
	Quantity uint32 `json:"quantity" binding:"required" example:"1"`
}

type DeleteOrderProductUriRequest struct {
	OrderID   uint64 `uri:"order_id" binding:"required"`
	ProductID uint64 `uri:"product_id" binding:"required"`
}

func NewOrderProductHandler(controller *controller.OrderProductController) *OrderProductHandler {
	return &OrderProductHandler{controller: controller}
}

func (h *OrderProductHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListOrderProducts)
	router.POST("/:order_id/:product_id", h.CreateOrderProduct)
	router.GET("/:order_id/:product_id", h.GetOrderProduct)
	router.PUT("/:order_id/:product_id", h.UpdateOrderProduct)
	router.DELETE("/:order_id/:product_id", h.DeleteOrderProduct)
}

// ListOrderProducts godoc
//
//	@Summary		List orders
//	@Description	List all orders
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id	query		string										false	"Filter by order ID"
//	@Param			page		query		int											false	"Page number"		default(1)
//	@Param			limit		query		int											false	"Items per page"	default(10)
//	@Success		200			{object}	presenter.OrderProductJsonPaginatedResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorJsonResponse				"Bad Request"
//	@Failure		500			{object}	middleware.ErrorJsonResponse				"Internal Server Error"
//	@Router			/api/v1/orders/products [get]
func (h *OrderProductHandler) ListOrderProducts(c *gin.Context) {
	var req ListOrderProductsQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.ListOrderProductsInput{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
		Page:      req.Page,
		Limit:     req.Limit,
	}
	h.controller.Presenter = presenter.NewOrderProductJsonPresenter(c)
	err := h.controller.List(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

//	@Router	/orders [post]

// CreateOrderProduct godoc
//
//	@Summary		Create an order product
//	@Description	Create an order product
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int									true	"Order ID"
//	@Param			product_id	path		int									true	"Product ID"
//	@Param			order		body		CreateOrderProductBodyRequest		true	"OrderProduct data"
//	@Success		201			{object}	presenter.OrderProductJsonResponse	"Created"
//	@Failure		400			{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [post]
func (h *OrderProductHandler) CreateOrderProduct(c *gin.Context) {
	var reqUri CreateOrderProductUriRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var req CreateOrderProductBodyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.CreateOrderProductInput{
		OrderID:   reqUri.OrderID,
		ProductID: reqUri.ProductID,
		Quantity:  req.Quantity,
	}
	h.controller.Presenter = presenter.NewOrderProductJsonPresenter(c)
	err := h.controller.Create(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// GetOrderProduct godoc
//
//	@Summary		Get an order product
//	@Description	Get an order product
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int									true	"Order ID"
//	@Param			product_id	path		int									true	"Product ID"
//	@Success		200			{object}	presenter.OrderProductJsonResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		404			{object}	middleware.ErrorJsonResponse		"Not Found"
//	@Failure		500			{object}	middleware.ErrorJsonResponse		"Internal Server Error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [get]
func (h *OrderProductHandler) GetOrderProduct(c *gin.Context) {
	var req GetOrderProductUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetOrderProductInput{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
	}
	h.controller.Presenter = presenter.NewOrderProductJsonPresenter(c)
	err := h.controller.Get(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// UpdateOrderProduct godoc
//
//	@Summary		Update order
//	@Description	Update an existing order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int									true	"Order ID"
//	@Param			product_id	path		int									true	"Product ID"
//	@Param			order		body		UpdateOrderProductBodyRequest		true	"OrderProduct data"
//	@Success		200			{object}	presenter.OrderProductJsonResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		404			{object}	middleware.ErrorJsonResponse		"Not Found"
//	@Failure		500			{object}	middleware.ErrorJsonResponse		"Internal Server Error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [put]
func (h *OrderProductHandler) UpdateOrderProduct(c *gin.Context) {
	var reqUri UpdateOrderProductUriRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}
	var req UpdateOrderProductBodyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdateOrderProductInput{
		OrderID:   reqUri.OrderID,
		ProductID: reqUri.ProductID,
		Quantity:  req.Quantity,
	}
	h.controller.Presenter = presenter.NewOrderProductJsonPresenter(c)
	err := h.controller.Update(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// DeleteOrderProduct godoc
//
//	@Summary		Delete orderProduct
//	@Description	Deletes a orderProduct by Order ID and Product ID
//	@Tags			orders
//	@Produce		json
//	@Param			order_id	path		int	true	"Order ID"
//	@Param			product_id	path		int	true	"Product ID"
//	@Success		204			{object}	nil
//	@Failure		400			{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404			{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500			{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [delete]
func (h *OrderProductHandler) DeleteOrderProduct(c *gin.Context) {
	var req DeleteOrderProductUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteOrderProductInput{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
	}
	h.controller.Presenter = presenter.NewOrderProductJsonPresenter(c)
	if err := h.controller.Delete(c.Request.Context(), input); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
