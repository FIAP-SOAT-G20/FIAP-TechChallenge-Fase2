package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type ListCustomersUseCase interface {
	Execute(ctx context.Context, input dto.ListCustomersInput) ([]*entity.Customer, int64, error)
}

type CreateCustomerUseCase interface {
	Execute(ctx context.Context, input dto.CreateCustomerInput) (*entity.Customer, error)
}

type GetCustomerUseCase interface {
	Execute(ctx context.Context, input dto.GetCustomerInput) (*entity.Customer, error)
}

type UpdateCustomerUseCase interface {
	Execute(ctx context.Context, input dto.UpdateCustomerInput) (*entity.Customer, error)
}

type DeleteCustomerUseCase interface {
	Execute(ctx context.Context, input dto.DeleteCustomerInput) (*entity.Customer, error)
}
