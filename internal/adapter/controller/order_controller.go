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
}

func NewOrderController(
	listUC port.ListOrdersUseCase,
	createUC port.CreateOrderUseCase,
	getUC port.GetOrderUseCase,
	updateUC port.UpdateOrderUseCase,
	deleteUC port.DeleteOrderUseCase,
) *OrderController {
	return &OrderController{
		listOrdersUseCase:  listUC,
		createOrderUseCase: createUC,
		getOrderUseCase:    getUC,
		updateOrderUseCase: updateUC,
		deleteOrderUseCase: deleteUC,
	}
}

func (c *OrderController) ListOrders(ctx context.Context, input dto.ListOrdersInput) error {
	err := c.listOrdersUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *OrderController) CreateOrder(ctx context.Context, input dto.CreateOrderInput) error {
	err := c.createOrderUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *OrderController) GetOrder(ctx context.Context, input dto.GetOrderInput) error {
	err := c.getOrderUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *OrderController) UpdateOrder(ctx context.Context, input dto.UpdateOrderInput) error {
	err := c.updateOrderUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *OrderController) DeleteOrder(ctx context.Context, input dto.DeleteOrderInput) error {
	return c.deleteOrderUseCase.Execute(ctx, input)
}
