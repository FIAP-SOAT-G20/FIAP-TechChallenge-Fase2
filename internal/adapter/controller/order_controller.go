package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderController struct {
	useCase   port.OrderUseCase
	Presenter port.Presenter
}

func NewOrderController(
	useCase port.OrderUseCase,
) *OrderController {
	return &OrderController{useCase, nil}
}

func (c *OrderController) List(ctx context.Context, input dto.ListOrdersInput) error {
	orders, total, err := c.useCase.List(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: orders,
	})

	return nil
}

func (c *OrderController) Create(ctx context.Context, input dto.CreateOrderInput) error {
	order, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) Get(ctx context.Context, input dto.GetOrderInput) error {
	order, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) Update(ctx context.Context, input dto.UpdateOrderInput) error {
	order, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) Delete(ctx context.Context, input dto.DeleteOrderInput) error {
	order, err := c.useCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: order,
	})

	return nil
}
