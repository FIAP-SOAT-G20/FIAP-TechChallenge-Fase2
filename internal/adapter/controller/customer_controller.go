package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type CustomerController struct {
	listUC    port.ListCustomersUseCase
	createUC  port.CreateCustomerUseCase
	getUC     port.GetCustomerUseCase
	updateUC  port.UpdateCustomerUseCase
	deleteUC  port.DeleteCustomerUseCase
	presenter port.CustomerPresenter
}

func NewCustomerController(
	listUC port.ListCustomersUseCase,
	createUC port.CreateCustomerUseCase,
	getUC port.GetCustomerUseCase,
	updateUC port.UpdateCustomerUseCase,
	deleteUC port.DeleteCustomerUseCase,
	presenter port.CustomerPresenter,
) *CustomerController {
	return &CustomerController{
		listUC,
		createUC,
		getUC,
		updateUC,
		deleteUC,
		presenter,
	}
}

func (c *CustomerController) ListCustomers(ctx context.Context, input dto.ListCustomersInput) error {
	customers, total, err := c.listUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: customers,
	})

	return nil
}

func (c *CustomerController) CreateCustomer(ctx context.Context, input dto.CreateCustomerInput) error {
	customer, err := c.createUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Result: customer,
	})

	return nil
}

func (c *CustomerController) GetCustomer(ctx context.Context, input dto.GetCustomerInput) error {
	customer, err := c.getUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Result: customer,
	})

	return nil
}

func (c *CustomerController) UpdateCustomer(ctx context.Context, input dto.UpdateCustomerInput) error {
	customer, err := c.updateUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Result: customer,
	})

	return nil
}

func (c *CustomerController) DeleteCustomer(ctx context.Context, input dto.DeleteCustomerInput) error {
	customer, err := c.deleteUC.Execute(ctx, input)
	if err != nil {
		return err
	}

	c.presenter.Present(dto.CustomerPresenterInput{
		Writer: input.Writer,
		Result: customer,
	})

	return nil
}
