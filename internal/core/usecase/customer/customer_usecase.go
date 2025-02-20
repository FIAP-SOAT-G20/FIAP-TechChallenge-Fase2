package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type customerUseCase struct {
	gateway port.CustomerGateway
}

// NewCustomerUseCase creates a new CreateCustomerUseCase
func NewCustomerUseCase(gateway port.CustomerGateway) port.CustomerUseCase {
	return &customerUseCase{gateway}
}

// List returns a list of Customers
func (uc *customerUseCase) List(ctx context.Context, input dto.ListCustomersInput) ([]*entity.Customer, int64, error) {
	customers, total, err := uc.gateway.FindAll(ctx, input.Name, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return customers, total, nil
}

// Create creates a new Customer
func (uc *customerUseCase) Create(ctx context.Context, input dto.CreateCustomerInput) (*entity.Customer, error) {
	customer := entity.NewCustomer(input.Name, input.Email, input.CPF)

	if err := uc.gateway.Create(ctx, customer); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return customer, nil
}

// Get returns a Customer by ID
func (uc *customerUseCase) Get(ctx context.Context, input dto.GetCustomerInput) (*entity.Customer, error) {
	customer, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return customer, nil
}

// Update updates a Customer
func (uc *customerUseCase) Update(ctx context.Context, input dto.UpdateCustomerInput) (*entity.Customer, error) {
	customer, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	customer.Update(input.Name, input.Email)

	if err := uc.gateway.Update(ctx, customer); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return customer, nil
}

// Delete deletes a Customer
func (uc *customerUseCase) Delete(ctx context.Context, input dto.DeleteCustomerInput) (*entity.Customer, error) {
	customer, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return customer, nil
}
