package controller

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type CustomerController struct {
	listCustomersUseCase  port.ListCustomersUseCase
	createCustomerUseCase port.CreateCustomerUseCase
	getCustomerUseCase    port.GetCustomerUseCase
	updateCustomerUseCase port.UpdateCustomerUseCase
	deleteCustomerUseCase port.DeleteCustomerUseCase
}

func NewCustomerController(
	listUC port.ListCustomersUseCase,
	createUC port.CreateCustomerUseCase,
	getUC port.GetCustomerUseCase,
	updateUC port.UpdateCustomerUseCase,
	deleteUC port.DeleteCustomerUseCase,
) *CustomerController {
	return &CustomerController{
		listCustomersUseCase:  listUC,
		createCustomerUseCase: createUC,
		getCustomerUseCase:    getUC,
		updateCustomerUseCase: updateUC,
		deleteCustomerUseCase: deleteUC,
	}
}

func (c *CustomerController) ListCustomers(ctx context.Context, input dto.ListCustomersInput) error {
	err := c.listCustomersUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *CustomerController) CreateCustomer(ctx context.Context, input dto.CreateCustomerInput) error {
	err := c.createCustomerUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *CustomerController) GetCustomer(ctx context.Context, input dto.GetCustomerInput) error {
	err := c.getCustomerUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *CustomerController) UpdateCustomer(ctx context.Context, input dto.UpdateCustomerInput) error {
	err := c.updateCustomerUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c *CustomerController) DeleteCustomer(ctx context.Context, input dto.DeleteCustomerInput) error {
	return c.deleteCustomerUseCase.Execute(ctx, input)
}
