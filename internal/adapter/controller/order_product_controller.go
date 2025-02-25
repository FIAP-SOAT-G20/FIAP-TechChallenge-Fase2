package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderProductController struct {
	useCase port.OrderProductUseCase
}

func NewOrderProductController(useCase port.OrderProductUseCase) port.OrderProductController {
	return &OrderProductController{useCase}
}

func (c *OrderProductController) List(ctx context.Context, p port.Presenter, i dto.ListOrderProductsInput) error {
	orderProducts, total, err := c.useCase.List(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{
		Total:  total,
		Page:   i.Page,
		Limit:  i.Limit,
		Result: orderProducts,
	})

	return nil
}

func (c *OrderProductController) Create(ctx context.Context, p port.Presenter, i dto.CreateOrderProductInput) error {
	orderProduct, err := c.useCase.Create(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: orderProduct})

	return nil
}

func (c *OrderProductController) Get(ctx context.Context, p port.Presenter, i dto.GetOrderProductInput) error {
	orderProduct, err := c.useCase.Get(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: orderProduct})

	return nil
}

func (c *OrderProductController) Update(ctx context.Context, p port.Presenter, i dto.UpdateOrderProductInput) error {
	orderProduct, err := c.useCase.Update(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: orderProduct})

	return nil
}

func (c *OrderProductController) Delete(ctx context.Context, p port.Presenter, i dto.DeleteOrderProductInput) error {
	order, err := c.useCase.Delete(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: order})

	return nil
}
