package handler

import (
	"fmt"
	"net/http"

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
	router.POST("/callback", h.Update)
}

// Create godoc
//
//	@Summary		Create a payment
//	@Description	Creates a new payment
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			payment body									body		request.CreatePaymentRequest	true	"Payment data"
//	@Success		201										{object}	presenter.PaymentJsonResponse	"Created"
//	@Failure		400										{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		500										{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/api/v1/payments/{order_id}/checkout	[post]
func (h *PaymentHandler) Create(c *gin.Context) {
	var uri request.CreatePaymentUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.CreatePaymentInput{
		OrderID: uri.OrderID,
	}

	output, err := h.controller.Create(
		c.Request.Context(),
		presenter.NewPaymentJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}

// Create godoc
//
//	@Summary		Update a payment
//	@Description	Update a new payment
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			payment	body					request.UpdatePaymentRequest	true	"Payment data"
//	@Success		201										{object}	presenter.PaymentJsonResponse	"Created"
//	@Failure		400										{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		500										{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/api/v1/payments/{order_id}/checkout	[post]
func (h *PaymentHandler) Update(c *gin.Context) {

	var body request.UpdatePaymentRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdatePaymentInput{
		Resource: body.Resource,
		Topic:    body.Topic,
	}

	output, err := h.controller.Update(
		c.Request.Context(),
		presenter.NewPaymentJsonPresenter(),
		input,
	)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", output)
}
