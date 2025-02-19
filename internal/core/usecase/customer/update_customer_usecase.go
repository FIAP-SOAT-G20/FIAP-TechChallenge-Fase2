package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateCustomerUseCase struct {
	gateway port.CustomerGateway
}

// NewUpdateCustomerUseCase creates a new UpdateCustomerUseCase
func NewUpdateCustomerUseCase(gateway port.CustomerGateway) port.UpdateCustomerUseCase {
	return &updateCustomerUseCase{gateway}
}

// Execute updates a customer
func (uc *updateCustomerUseCase) Execute(ctx context.Context, input dto.UpdateCustomerInput) (*entity.Customer, error) {
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
