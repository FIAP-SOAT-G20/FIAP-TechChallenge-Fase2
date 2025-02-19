package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteCustomerUseCase struct {
	gateway port.CustomerGateway
}

// NewDeleteCustomerUseCase creates a new DeleteCustomerUseCase
func NewDeleteCustomerUseCase(gateway port.CustomerGateway) port.DeleteCustomerUseCase {
	return &deleteCustomerUseCase{gateway}
}

// Execute deletes a customer
func (uc *deleteCustomerUseCase) Execute(ctx context.Context, input dto.DeleteCustomerInput) (*entity.Customer, error) {
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
