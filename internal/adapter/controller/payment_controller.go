package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type PaymentController struct {
	createPaymentUseCase port.CreatePaymentUseCase
}

func NewPaymentController(
	createUC port.CreatePaymentUseCase,
) *PaymentController {
	return &PaymentController{
		createPaymentUseCase: createUC,
	}
}

func (c *PaymentController) CreatePayment(ctx context.Context, OrderID uint64, writer dto.ResponseWriter) error {
	err := c.createPaymentUseCase.Execute(ctx, OrderID, writer)
	if err != nil {
		return err
	}

	return nil
}
