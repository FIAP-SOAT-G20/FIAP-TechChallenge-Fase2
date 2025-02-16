package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderProductController struct {
	listOrderProductsUseCase  port.ListOrderProductsUseCase
	createOrderProductUseCase port.CreateOrderProductUseCase
	getOrderProductUseCase    port.GetOrderProductUseCase
	updateOrderProductUseCase port.UpdateOrderProductUseCase
	deleteOrderProductUseCase port.DeleteOrderProductUseCase
}

func NewOrderProductController(
	listUC port.ListOrderProductsUseCase,
	createUC port.CreateOrderProductUseCase,
	getUC port.GetOrderProductUseCase,
	updateUC port.UpdateOrderProductUseCase,
	deleteUC port.DeleteOrderProductUseCase,
) *OrderProductController {
	return &OrderProductController{
		listOrderProductsUseCase:  listUC,
		createOrderProductUseCase: createUC,
		getOrderProductUseCase:    getUC,
		updateOrderProductUseCase: updateUC,
		deleteOrderProductUseCase: deleteUC,
	}
}

func (c *OrderProductController) ListOrderProducts(ctx context.Context, input dto.ListOrderProductsInput) error {
	err := c.listOrderProductsUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *OrderProductController) CreateOrderProduct(ctx context.Context, input dto.CreateOrderProductInput) error {
	err := c.createOrderProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *OrderProductController) GetOrderProduct(ctx context.Context, input dto.GetOrderProductInput) error {
	err := c.getOrderProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *OrderProductController) UpdateOrderProduct(ctx context.Context, input dto.UpdateOrderProductInput) error {
	err := c.updateOrderProductUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *OrderProductController) DeleteOrderProduct(ctx context.Context, input dto.DeleteOrderProductInput) error {
	return c.deleteOrderProductUseCase.Execute(ctx, input)
}
