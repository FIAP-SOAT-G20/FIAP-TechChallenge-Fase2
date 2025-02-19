package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderController struct {
	listOrdersUseCase  port.ListOrdersUseCase
	createOrderUseCase port.CreateOrderUseCase
	getOrderUseCase    port.GetOrderUseCase
	updateOrderUseCase port.UpdateOrderUseCase
	deleteOrderUseCase port.DeleteOrderUseCase
	presenter          port.OrderPresenter
}

func NewOrderController(
	listUC port.ListOrdersUseCase,
	createUC port.CreateOrderUseCase,
	getUC port.GetOrderUseCase,
	updateUC port.UpdateOrderUseCase,
	deleteUC port.DeleteOrderUseCase,
	orderPresenter port.OrderPresenter,
) *OrderController {
	return &OrderController{
		listOrdersUseCase:  listUC,
		createOrderUseCase: createUC,
		getOrderUseCase:    getUC,
		updateOrderUseCase: updateUC,
		deleteOrderUseCase: deleteUC,
		presenter:          orderPresenter,
	}
}

func (c *OrderController) ListOrders(ctx context.Context, input dto.ListOrdersInput) error {
	orders, total, err := c.listOrdersUseCase.Execute(ctx, input)
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
	order, err := c.createOrderUseCase.Execute(ctx, input)
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
	order, err := c.getOrderUseCase.Execute(ctx, input)
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
	order, err := c.updateOrderUseCase.Execute(ctx, input)
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
	order, err := c.deleteOrderUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Result: order,
	})

	return nil
}
