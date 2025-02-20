package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type CustomerController struct {
	customerUC port.CustomerUseCase
	Presenter  port.CustomerPresenter
}

func NewCustomerController(customerUC port.CustomerUseCase) *CustomerController {
	return &CustomerController{customerUC, nil}
}

func (c *CustomerController) ListCustomers(ctx context.Context, input dto.ListCustomersInput) error {
	customers, total, err := c.customerUC.List(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.CustomerPresenterInput{
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: customers,
	})

	return nil
}

func (c *CustomerController) CreateCustomer(ctx context.Context, input dto.CreateCustomerInput) error {
	customer, err := c.customerUC.Create(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.CustomerPresenterInput{
		Result: customer,
	})

	return nil
}

func (c *CustomerController) GetCustomer(ctx context.Context, input dto.GetCustomerInput) error {
	customer, err := c.customerUC.Get(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.CustomerPresenterInput{
		Result: customer,
	})

	return nil
}

func (c *CustomerController) UpdateCustomer(ctx context.Context, input dto.UpdateCustomerInput) error {
	customer, err := c.customerUC.Update(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.CustomerPresenterInput{
		Result: customer,
	})

	return nil
}

func (c *CustomerController) DeleteCustomer(ctx context.Context, input dto.DeleteCustomerInput) error {
	customer, err := c.customerUC.Delete(ctx, input)
	if err != nil {
		return err
	}

	c.Presenter.Present(dto.CustomerPresenterInput{
		Result: customer,
	})

	return nil
}
