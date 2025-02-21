package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/request"
)

type OrderHandler struct {
	controller *controller.OrderController
}

func NewOrderHandler(controller *controller.OrderController) *OrderHandler {
	return &OrderHandler{controller: controller}
}

func (h *OrderHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.List)
	router.POST("/", h.Create)
	router.GET("/:id", h.Get)
	router.PUT("/:id", h.Update)
	router.PATCH("/:id", h.UpdatePartial)
	router.DELETE("/:id", h.Delete)
}

// List godoc
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
//	@Failure		400			{object}	middleware.ErrorJsonResponse			"Bad Request"
//	@Failure		500			{object}	middleware.ErrorJsonResponse			"Internal Server Error"
//	@Router			/orders [get]
func (h *OrderHandler) List(c *gin.Context) {
	var query request.ListOrdersQueryRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.ListOrdersInput{
		CustomerID: query.CustomerID,
		Status:     query.Status,
		Page:       query.Page,
		Limit:      query.Limit,
	}
	h.controller.Presenter = presenter.NewOrderJsonPresenter(c)
	err := h.controller.List(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Create godoc
//
//	@Summary		Create order
//	@Description	Creates a new order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		request.CreateOrderBodyRequest			true	"Order data"
//	@Success		201		{object}	presenter.OrderJsonResponse		"Created"
//	@Failure		400		{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		500		{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/orders [post]
func (h *OrderHandler) Create(c *gin.Context) {
	var body request.CreateOrderBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.CreateOrderInput{
		CustomerID: body.CustomerID,
	}
	h.controller.Presenter = presenter.NewOrderJsonPresenter(c)
	err := h.controller.Create(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Get godoc
//
//	@Summary		Get order
//	@Description	Search for a order by ID
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"Order ID"
//	@Success		200	{object}	presenter.OrderJsonResponse		"OK"
//	@Failure		400	{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/orders/{id} [get]
func (h *OrderHandler) Get(c *gin.Context) {
	var uri request.GetOrderUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetOrderInput{
		ID: uri.ID,
	}
	h.controller.Presenter = presenter.NewOrderJsonPresenter(c)
	err := h.controller.Get(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Update godoc
//
//	@Summary		Update order
//	@Description	Update an existing order
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
//	@Param			id		path		int								true	"Order ID"
//	@Param			order	body		request.UpdateOrderBodyRequest			true	"Order data"
//	@Success		200		{object}	presenter.OrderJsonResponse		"OK"
//	@Failure		400		{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404		{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500		{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/orders/{id} [put]
func (h *OrderHandler) Update(c *gin.Context) {
	var uri request.UpdateOrderUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var body request.UpdateOrderBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}
	h.controller.Presenter = presenter.NewOrderJsonPresenter(c)
	input := dto.UpdateOrderInput{
		ID:         uri.ID,
		CustomerID: body.CustomerID,
		Status:     body.Status,
		StaffID:    body.StaffID,
	}

	err := h.controller.Update(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// UpdatePartial godoc
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
//	@Param			id		path		int								true	"Order ID"
//	@Param			order	body		request.UpdateOrderPartilRequest		true	"Order data"
//	@Success		200		{object}	presenter.OrderJsonResponse		"OK"
//	@Failure		400		{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404		{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500		{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/orders/{id} [patch]
func (h *OrderHandler) UpdatePartial(c *gin.Context) {
	var uri request.UpdateOrderUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var body request.UpdateOrderPartilBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}
	h.controller.Presenter = presenter.NewOrderJsonPresenter(c)
	input := dto.UpdateOrderInput{
		ID:         uri.ID,
		CustomerID: body.CustomerID,
		Status:     body.Status,
		StaffID:    body.StaffID,
	}

	err := h.controller.Update(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Delete godoc
//
//	@Summary		Delete order
//	@Description	Deletes a order by ID
//	@Tags			orders
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/orders/{id} [delete]
func (h *OrderHandler) Delete(c *gin.Context) {
	var uri request.DeleteOrderUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteOrderInput{
		ID: uri.ID,
	}
	h.controller.Presenter = presenter.NewOrderJsonPresenter(c)
	if err := h.controller.Delete(c.Request.Context(), input); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
