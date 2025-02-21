package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type CustomerController struct {
	useCase   port.CustomerUseCase
	Presenter port.Presenter
}

func NewCustomerController(useCase port.CustomerUseCase) *CustomerController {
	return &CustomerController{useCase, nil}
}

func (c *CustomerController) List(ctx context.Context, input dto.ListCustomersInput) error {
	customers, total, err := c.useCase.List(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: customers,
	})

	return nil
}

func (c *CustomerController) Create(ctx context.Context, input dto.CreateCustomerInput) error {
	customer, err := c.useCase.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: customer,
	})

	return nil
}

func (c *CustomerController) Get(ctx context.Context, input dto.GetCustomerInput) error {
	customer, err := c.useCase.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: customer,
	})

	return nil
}

func (c *CustomerController) Update(ctx context.Context, input dto.UpdateCustomerInput) error {
	customer, err := c.useCase.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: customer,
	})

	return nil
}

func (c *CustomerController) Delete(ctx context.Context, input dto.DeleteCustomerInput) error {
	customer, err := c.useCase.Delete(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.PresenterInput{
		Result: customer,
	})

	return nil
}
