package customer

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listCustomersUseCase struct {
	gateway port.CustomerGateway
}

// NewListCustomersUseCase creates a new ListCustomersUseCase
func NewListCustomersUseCase(gateway port.CustomerGateway) port.ListCustomersUseCase {
	return &listCustomersUseCase{gateway}
}

// Execute lists all customers
func (uc *listCustomersUseCase) Execute(ctx context.Context, input dto.ListCustomersInput) ([]*entity.Customer, int64, error) {
	customers, total, err := uc.gateway.FindAll(ctx, input.Name, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return customers, total, nil
}
