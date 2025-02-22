package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
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

func (c *PaymentController) CreatePayment(ctx context.Context, p port.Presenter, OrderID uint64) error {
	payment, err := c.createPaymentUseCase.Execute(ctx, OrderID)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{
		Result: payment,
	})

	return nil
}
