package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderProductController struct {
	useCase   port.OrderProductUseCase
	Presenter port.OrderProductPresenter
}

func NewOrderProductController(useCase port.OrderProductUseCase) *OrderProductController {
	return &OrderProductController{useCase, nil}
}

func (c *OrderProductController) List(ctx context.Context, input dto.ListOrderProductsInput) error {
	orderProducts, total, err := c.useCase.List(ctx, input)
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

func (c *OrderProductController) Create(ctx context.Context, input dto.CreateOrderProductInput) error {
	orderProduct, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) Get(ctx context.Context, input dto.GetOrderProductInput) error {
	orderProduct, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) Update(ctx context.Context, input dto.UpdateOrderProductInput) error {
	orderProduct, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Result: orderProduct,
	})

	return nil
}

func (c *OrderProductController) Delete(ctx context.Context, input dto.DeleteOrderProductInput) error {
	order, err := c.useCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.OrderProductPresenterInput{
		Result: order,
	})

	return nil
}
