package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/request"
)

type OrderProductHandler struct {
	controller port.OrderProductController
}

func NewOrderProductHandler(controller port.OrderProductController) *OrderProductHandler {
	return &OrderProductHandler{controller: controller}
}

func (h *OrderProductHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.List)
	router.POST("/:order_id/:product_id", h.Create)
	router.GET("/:order_id/:product_id", h.Get)
	router.PUT("/:order_id/:product_id", h.Update)
	router.DELETE("/:order_id/:product_id", h.Delete)
}

// List godoc
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
func (h *OrderProductHandler) List(c *gin.Context) {
	var query request.ListOrderProductsQueryRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.ListOrderProductsInput{
		OrderID:   query.OrderID,
		ProductID: query.ProductID,
		Page:      query.Page,
		Limit:     query.Limit,
	}

	err := h.controller.List(
		c.Request.Context(),
		presenter.NewOrderProductJsonPresenter(c),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

//	@Router	/orders [post]

// Create godoc
//
//	@Summary		Create an order product
//	@Description	Create an order product
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int										true	"Order ID"
//	@Param			product_id	path		int										true	"Product ID"
//	@Param			order		body		request.CreateOrderProductBodyRequest	true	"OrderProduct data"
//	@Success		201			{object}	presenter.OrderProductJsonResponse		"Created"
//	@Failure		400			{object}	middleware.ErrorJsonResponse			"Bad Request"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [post]
func (h *OrderProductHandler) Create(c *gin.Context) {
	var uri request.CreateOrderProductUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var body request.CreateOrderProductBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.CreateOrderProductInput{
		OrderID:   uri.OrderID,
		ProductID: uri.ProductID,
		Quantity:  body.Quantity,
	}

	err := h.controller.Create(
		c.Request.Context(),
		presenter.NewOrderProductJsonPresenter(c),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Get godoc
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
func (h *OrderProductHandler) Get(c *gin.Context) {
	var uri request.GetOrderProductUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetOrderProductInput{
		OrderID:   uri.OrderID,
		ProductID: uri.ProductID,
	}

	err := h.controller.Get(
		c.Request.Context(),
		presenter.NewOrderProductJsonPresenter(c),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Update godoc
//
//	@Summary		Update order
//	@Description	Update an existing order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id	path		int										true	"Order ID"
//	@Param			product_id	path		int										true	"Product ID"
//	@Param			order		body		request.UpdateOrderProductBodyRequest	true	"OrderProduct data"
//	@Success		200			{object}	presenter.OrderProductJsonResponse		"OK"
//	@Failure		400			{object}	middleware.ErrorJsonResponse			"Bad Request"
//	@Failure		404			{object}	middleware.ErrorJsonResponse			"Not Found"
//	@Failure		500			{object}	middleware.ErrorJsonResponse			"Internal Server Error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [put]
func (h *OrderProductHandler) Update(c *gin.Context) {
	var uri request.UpdateOrderProductUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}
	var body request.UpdateOrderProductBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdateOrderProductInput{
		OrderID:   uri.OrderID,
		ProductID: uri.ProductID,
		Quantity:  body.Quantity,
	}

	err := h.controller.Update(
		c.Request.Context(),
		presenter.NewOrderProductJsonPresenter(c),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Delete godoc
//
//	@Summary		Delete orderProduct
//	@Description	Deletes a orderProduct by Order ID and Product ID
//	@Tags			orders
//	@Produce		json
//	@Param			order_id	path		int									true	"Order ID"
//	@Param			product_id	path		int									true	"Product ID"
//	@Success		200			{object}	presenter.OrderProductJsonResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		404			{object}	middleware.ErrorJsonResponse		"Not Found"
//	@Failure		500			{object}	middleware.ErrorJsonResponse		"Internal Server Error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [delete]
func (h *OrderProductHandler) Delete(c *gin.Context) {
	var uri request.DeleteOrderProductUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteOrderProductInput{
		OrderID:   uri.OrderID,
		ProductID: uri.ProductID,
	}

	err := h.controller.Delete(
		c.Request.Context(),
		presenter.NewOrderProductJsonPresenter(c),
		input,
	)

	if err != nil {
		_ = c.Error(err)
		return
	}
}
