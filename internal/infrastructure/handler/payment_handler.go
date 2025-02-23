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
	router.POST("/:order_id/checkout", h.CreatePayment)
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req request.CreatePaymentUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.CreatePaymentInput{
		OrderID: req.OrderID,
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
