package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderProductController struct {
	listUC    port.ListOrderProductsUseCase
	createUC  port.CreateOrderProductUseCase
	getUC     port.GetOrderProductUseCase
	updateUC  port.UpdateOrderProductUseCase
	deleteUC  port.DeleteOrderProductUseCase
	presenter port.OrderProductPresenter
}

func NewOrderProductController(
	listUC port.ListOrderProductsUseCase,
	createUC port.CreateOrderProductUseCase,
	getUC port.GetOrderProductUseCase,
	updateUC port.UpdateOrderProductUseCase,
	deleteUC port.DeleteOrderProductUseCase,
	presenter port.OrderProductPresenter,
) *OrderProductController {
	return &OrderProductController{
		listUC,
		createUC,
		getUC,
		updateUC,
		deleteUC,
		presenter,
	}
}

func (c *OrderProductController) ListOrderProducts(ctx context.Context, input dto.ListOrderProductsInput) error {
	orderProducts, total, err := c.listUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: orderProducts,
	})

	return nil
}

func (c *OrderProductController) CreateOrderProduct(ctx context.Context, input dto.CreateOrderProductInput) error {
	orderProduct, err := c.createUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) GetOrderProduct(ctx context.Context, input dto.GetOrderProductInput) error {
	orderProduct, err := c.getUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) UpdateOrderProduct(ctx context.Context, input dto.UpdateOrderProductInput) error {
	orderProduct, err := c.updateUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) DeleteOrderProduct(ctx context.Context, input dto.DeleteOrderProductInput) error {
	order, err := c.deleteUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Result: order,
	})

	return nil
}
