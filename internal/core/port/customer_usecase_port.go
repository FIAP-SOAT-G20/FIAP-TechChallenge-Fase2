package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
)

type ListCustomersUseCase interface {
	Execute(ctx context.Context, input dto.ListCustomersInput) error
}

type CreateCustomerUseCase interface {
	Execute(ctx context.Context, input dto.CreateCustomerInput) error
}

type GetCustomerUseCase interface {
	Execute(ctx context.Context, input dto.GetCustomerInput) error
}

type UpdateCustomerUseCase interface {
	Execute(ctx context.Context, input dto.UpdateCustomerInput) error
}

type DeleteCustomerUseCase interface {
	Execute(ctx context.Context, input dto.DeleteCustomerInput) error
}
