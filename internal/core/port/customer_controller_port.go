package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type CustomerController interface {
	List(ctx context.Context, presenter Presenter, input dto.ListCustomersInput) error
	Create(ctx context.Context, presenter Presenter, input dto.CreateCustomerInput) error
	Get(ctx context.Context, presenter Presenter, input dto.GetCustomerInput) error
	Update(ctx context.Context, presenter Presenter, input dto.UpdateCustomerInput) error
	Delete(ctx context.Context, presenter Presenter, input dto.DeleteCustomerInput) error
}
