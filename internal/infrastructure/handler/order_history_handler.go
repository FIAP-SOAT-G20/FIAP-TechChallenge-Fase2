package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/request"
)

type OrderHistoryHandler struct {
	controller port.OrderHistoryController
}

func NewOrderHistoryHandler(controller port.OrderHistoryController) *OrderHistoryHandler {
	return &OrderHistoryHandler{controller: controller}
}

func (h *OrderHistoryHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.List)
	router.GET("/:id", h.Get)
	router.DELETE("/:id", h.Delete)
}

// List godoc
//
//	@Summary		List orderHistories
//	@Description	List all orderHistories
//	@Tags			orderHistories
//	@Accept			json
//	@Produce		json
//	@Param			order_id	query		string										false	"Filter by order_id"
//	@Param			status		query		string										false	"Filter by status. Available options: OPEN, CANCELLED, PENDING, RECEIVED, PREPARING, READY, COMPLETED"
//	@Param			page		query		int											false	"Page number"		default(1)
//	@Param			limit		query		int											false	"Items per page"	default(10)
//	@Success		200			{object}	presenter.OrderHistoryJsonPaginatedResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorJsonResponse				"Bad Request"
//	@Failure		500			{object}	middleware.ErrorJsonResponse				"Internal Server Error"
//	@Router			/orders/histories [get]
func (h *OrderHistoryHandler) List(c *gin.Context) {
	var query request.ListOrderHistoriesQueryRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.ListOrderHistoriesInput{
		OrderID: query.OrderID,
		Status:  query.Status,
		Page:    query.Page,
		Limit:   query.Limit,
	}

	output, err := h.controller.List(
		c.Request.Context(),
		presenter.NewOrderHistoryJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}

// Get godoc
//
//	@Summary		Get orderHistory
//	@Description	Search for a orderHistory by ID
//	@Tags			orderHistories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int									true	"OrderHistory ID"
//	@Success		200	{object}	presenter.OrderHistoryJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse		"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse		"Internal Server Error"
//	@Router			/orders/histories/{id} [get]
func (h *OrderHistoryHandler) Get(c *gin.Context) {
	var uri request.GetOrderHistoryUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetOrderHistoryInput{
		ID: uri.ID,
	}

	output, err := h.controller.Get(
		c.Request.Context(),
		presenter.NewOrderHistoryJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}

// Delete godoc
//
//	@Summary		Delete orderHistory
//	@Description	Deletes a orderHistory by ID
//	@Tags			orderHistories
//	@Produce		json
//	@Param			id	path		int									true	"OrderHistory ID"
//	@Success		200	{object}	presenter.OrderHistoryJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse		"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse		"Internal Server Error"
//	@Router			/orders/histories/{id} [delete]
func (h *OrderHistoryHandler) Delete(c *gin.Context) {
	var uri request.DeleteOrderHistoryUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteOrderHistoryInput{
		ID: uri.ID,
	}
	output, err := h.controller.Delete(
		c.Request.Context(),
		presenter.NewOrderHistoryJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}
