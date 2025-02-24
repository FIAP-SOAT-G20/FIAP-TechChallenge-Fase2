package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type PaymentController struct {
	useCase port.PaymentUseCase
}

func NewPaymentController(useCase port.PaymentUseCase) *PaymentController {
	return &PaymentController{useCase}
}

func (c *PaymentController) Create(ctx context.Context, p port.Presenter, i dto.CreatePaymentInput) error {
	payment, err := c.useCase.Create(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: payment})

	return nil
}

func (c *PaymentController) Update(ctx context.Context, p port.Presenter, i dto.UpdatePaymentInput) error {
	payment, err := c.useCase.Update(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: payment})

	return nil
}
