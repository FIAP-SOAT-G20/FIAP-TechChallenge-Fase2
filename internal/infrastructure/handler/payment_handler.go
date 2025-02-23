package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/request"
)

type PaymentHandler struct {
	controller port.PaymentController
}

func NewPaymentHandler(controller port.PaymentController) *PaymentHandler {
	return &PaymentHandler{controller}
}

func (h *PaymentHandler) Register(router *gin.RouterGroup) {
	router.POST("/:order_id/checkout", h.Create)
}

// Create godoc
//
//	@Summary		Create a payment
//	@Description	Creates a new payment
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			product	body		request.CreatePaymentRequest	true		"Payment data"
//	@Success		201		{object}	presenter.PaymentJsonResponse		"Created"
//	@Failure		400		{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		500		{object}	middleware.ErrorJsonResponse		"Internal Server Error"
func (h *PaymentHandler) Create(c *gin.Context) {
	var uri request.CreatePaymentUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.CreatePaymentInput{
		OrderID: uri.OrderID,
	}

	err := h.controller.Create(
		c.Request.Context(),
		presenter.NewPaymentJsonPresenter(c),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}
}
