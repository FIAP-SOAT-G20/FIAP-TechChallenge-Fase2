package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type customerController struct {
	useCase port.CustomerUseCase
}

func NewCustomerController(useCase port.CustomerUseCase) port.CustomerController {
	return &customerController{useCase}
}

func (c *customerController) List(
	ctx context.Context,
	presenter port.Presenter,
	input dto.ListCustomersInput,
) error {
	customers, total, err := c.useCase.List(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: customers,
	})

	return nil
}

func (c *customerController) Create(
	ctx context.Context,
	presenter port.Presenter,
	input dto.CreateCustomerInput,
) error {
	customer, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Result: customer,
	})

	return nil
}

func (c *customerController) Get(
	ctx context.Context,
	presenter port.Presenter,
	input dto.GetCustomerInput,
) error {
	customer, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Result: customer,
	})

	return nil
}

func (c *customerController) Update(
	ctx context.Context,
	presenter port.Presenter,
	input dto.UpdateCustomerInput,
) error {
	customer, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Result: customer,
	})

	return nil
}

func (c *customerController) Delete(
	ctx context.Context,
	presenter port.Presenter,
	input dto.DeleteCustomerInput,
) error {
	customer, err := c.useCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	presenter.Present(dto.PresenterInput{
		Result: customer,
	})

	return nil
}
