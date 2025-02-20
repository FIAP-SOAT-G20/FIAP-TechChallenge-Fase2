package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type CreatePaymentUriRequest struct {
	OrderID uint64 `uri:"order_id" binding:"required"`
}

type PaymentHandler struct {
	controller *controller.PaymentController
}

func NewPaymentHandler(controller *controller.PaymentController) *PaymentHandler {
	return &PaymentHandler{controller: controller}
}

func (h *PaymentHandler) Register(router *gin.RouterGroup) {
	router.POST("/:order_id/checkout", h.CreatePayment)
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req CreatePaymentUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	err := h.controller.CreatePayment(c.Request.Context(), req.OrderID, c)
	if err != nil {
		_ = c.Error(err)
		return
	}
}
