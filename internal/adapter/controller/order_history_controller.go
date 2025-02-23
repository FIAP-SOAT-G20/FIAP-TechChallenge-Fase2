package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type OrderHistoryController struct {
	useCase port.OrderHistoryUseCase
}

func NewOrderHistoryController(useCase port.OrderHistoryUseCase) *OrderHistoryController {
	return &OrderHistoryController{useCase}
}

func (c *OrderHistoryController) List(ctx context.Context, p port.Presenter, i dto.ListOrderHistoriesInput) error {
	orderHistories, total, err := c.useCase.List(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{
		Total:  total,
		Page:   i.Page,
		Limit:  i.Limit,
		Result: orderHistories,
	})

	return nil
}

func (c *OrderHistoryController) Create(ctx context.Context, p port.Presenter, i dto.CreateOrderHistoryInput) error {
	orderHistory, err := c.useCase.Create(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: orderHistory})

	return nil
}

func (c *OrderHistoryController) Get(ctx context.Context, p port.Presenter, i dto.GetOrderHistoryInput) error {
	orderHistory, err := c.useCase.Get(ctx, i)
	if err != nil {
		return err
	}

	p.Present(dto.PresenterInput{Result: orderHistory})

	return nil
}

func (c *OrderHistoryController) Delete(ctx context.Context, p port.Presenter, i dto.DeleteOrderHistoryInput) error {
	orderHistory, err := c.useCase.Delete(ctx, i)

	if err != nil {
		return err
	}
	p.Present(dto.PresenterInput{Result: orderHistory})

	return nil
}
