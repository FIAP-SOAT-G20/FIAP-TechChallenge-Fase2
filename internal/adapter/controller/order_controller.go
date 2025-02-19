package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderController struct {
	listUC    port.ListOrdersUseCase
	createUC  port.CreateOrderUseCase
	getUC     port.GetOrderUseCase
	updateUC  port.UpdateOrderUseCase
	deleteUC  port.DeleteOrderUseCase
	presenter port.OrderPresenter
}

func NewOrderController(
	listUC port.ListOrdersUseCase,
	createUC port.CreateOrderUseCase,
	getUC port.GetOrderUseCase,
	updateUC port.UpdateOrderUseCase,
	deleteUC port.DeleteOrderUseCase,
	presenter port.OrderPresenter,
) *OrderController {
	return &OrderController{
		listUC,
		createUC,
		getUC,
		updateUC,
		deleteUC,
		presenter,
	}
}

func (c *OrderController) ListOrders(ctx context.Context, input dto.ListOrdersInput) error {
	orders, total, err := c.listUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: orders,
	})

	return nil
}

func (c *OrderController) CreateOrder(ctx context.Context, input dto.CreateOrderInput) error {
	order, err := c.createUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Result: order,
	})

	return nil
}

func (c *OrderController) GetOrder(ctx context.Context, input dto.GetOrderInput) error {
	order, err := c.getUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Result: order,
	})

	return nil
}

func (c *OrderController) UpdateOrder(ctx context.Context, input dto.UpdateOrderInput) error {
	order, err := c.updateUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Result: order,
	})

	return nil
}

func (c *OrderController) DeleteOrder(ctx context.Context, input dto.DeleteOrderInput) error {
	order, err := c.deleteUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Result: order,
	})

	return nil
}
