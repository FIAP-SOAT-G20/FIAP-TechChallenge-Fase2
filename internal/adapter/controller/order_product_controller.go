package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderProductController struct {
	orderProductUseCase port.OrderProductUseCase
	Presenter           port.OrderProductPresenter
}

func NewOrderProductController(orderProductUseCase port.OrderProductUseCase) *OrderProductController {
	return &OrderProductController{orderProductUseCase, nil}
}

func (c *OrderProductController) ListOrderProducts(ctx context.Context, input dto.ListOrderProductsInput) error {
	orderProducts, total, err := c.orderProductUseCase.List(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: orderProducts,
	})

	return nil
}

func (c *OrderProductController) CreateOrderProduct(ctx context.Context, input dto.CreateOrderProductInput) error {
	orderProduct, err := c.orderProductUseCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) GetOrderProduct(ctx context.Context, input dto.GetOrderProductInput) error {
	orderProduct, err := c.orderProductUseCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) UpdateOrderProduct(ctx context.Context, input dto.UpdateOrderProductInput) error {
	orderProduct, err := c.orderProductUseCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) DeleteOrderProduct(ctx context.Context, input dto.DeleteOrderProductInput) error {
	order, err := c.orderProductUseCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Result: order,
	})

	return nil
}
