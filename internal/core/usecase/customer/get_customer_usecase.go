package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type getCustomerUseCase struct {
	gateway port.CustomerGateway
}

// NewGetCustomerUseCase creates a new GetCustomerUseCase
func NewGetCustomerUseCase(gateway port.CustomerGateway) port.GetCustomerUseCase {
	return &getCustomerUseCase{gateway}
}

// Execute gets a customer
func (uc *getCustomerUseCase) Execute(ctx context.Context, input dto.GetCustomerInput) (*entity.Customer, error) {
	customer, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if customer == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return customer, nil
}
