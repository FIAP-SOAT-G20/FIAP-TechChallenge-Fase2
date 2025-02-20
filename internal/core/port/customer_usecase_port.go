package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type CustomerUseCase interface {
	List(ctx context.Context, input dto.ListCustomersInput) ([]*entity.Customer, int64, error)
	Create(ctx context.Context, input dto.CreateCustomerInput) (*entity.Customer, error)
	Get(ctx context.Context, input dto.GetCustomerInput) (*entity.Customer, error)
	Update(ctx context.Context, input dto.UpdateCustomerInput) (*entity.Customer, error)
	Delete(ctx context.Context, input dto.DeleteCustomerInput) (*entity.Customer, error)
}
