package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderController struct {
	useCase port.OrderUseCase
}

func NewOrderController(useCase port.OrderUseCase) port.OrderController {
	return &OrderController{useCase}
}

func (c *OrderController) List(ctx context.Context, presenter port.Presenter, input dto.ListOrdersInput) error {
	orders, total, err := c.useCase.List(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: orders,
	})

	return nil
}

func (c *OrderController) Create(ctx context.Context, presenter port.Presenter, input dto.CreateOrderInput) error {
	order, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) Get(ctx context.Context, presenter port.Presenter, input dto.GetOrderInput) error {
	order, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) Update(ctx context.Context, presenter port.Presenter, input dto.UpdateOrderInput) error {
	order, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) Delete(ctx context.Context, presenter port.Presenter, input dto.DeleteOrderInput) error {
	order, err := c.useCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Result: order,
	})

	return nil
}
