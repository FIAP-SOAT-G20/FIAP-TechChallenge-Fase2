package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderController struct {
	orderUseCase port.OrderUseCase
	Presenter    port.OrderPresenter
}

func NewOrderController(
	orderUseCase port.OrderUseCase,
) *OrderController {
	return &OrderController{orderUseCase, nil}
}

func (c *OrderController) ListOrders(ctx context.Context, input dto.ListOrdersInput) error {
	orders, total, err := c.orderUseCase.List(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderPresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: orders,
	})

	return nil
}

func (c *OrderController) CreateOrder(ctx context.Context, input dto.CreateOrderInput) error {
	order, err := c.orderUseCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderPresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) GetOrder(ctx context.Context, input dto.GetOrderInput) error {
	order, err := c.orderUseCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderPresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) UpdateOrder(ctx context.Context, input dto.UpdateOrderInput) error {
	order, err := c.orderUseCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderPresenterInput{
		Result: order,
	})

	return nil
}

func (c *OrderController) DeleteOrder(ctx context.Context, input dto.DeleteOrderInput) error {
	order, err := c.orderUseCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderPresenterInput{
		Result: order,
	})

	return nil
}
