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

func (c *OrderController) List(ctx context.Context, p port.Presenter, i dto.ListOrdersInput) error {
	orders, total, err := c.useCase.List(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{
		Total:  total,
		Page:   i.Page,
		Limit:  i.Limit,
		Result: orders,
	})

	return nil
}

func (c *OrderController) Create(ctx context.Context, p port.Presenter, i dto.CreateOrderInput) error {
	order, err := c.useCase.Create(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: order})

	return nil
}

func (c *OrderController) Get(ctx context.Context, p port.Presenter, i dto.GetOrderInput) error {
	order, err := c.useCase.Get(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: order})

	return nil
}

func (c *OrderController) Update(ctx context.Context, p port.Presenter, i dto.UpdateOrderInput) error {
	order, err := c.useCase.Update(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: order})

	return nil
}

func (c *OrderController) Delete(ctx context.Context, p port.Presenter, i dto.DeleteOrderInput) error {
	order, err := c.useCase.Delete(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: order})

	return nil
}
