package order

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteOrderUseCase struct {
	gateway   port.OrderGateway
	presenter port.OrderPresenter
}

// NewDeleteOrderUseCase creates a new DeleteOrderUseCase
func NewDeleteOrderUseCase(gateway port.OrderGateway, presenter port.OrderPresenter) port.DeleteOrderUseCase {
	return &deleteOrderUseCase{gateway, presenter}
}

// Execute deletes a order
func (uc *deleteOrderUseCase) Execute(ctx context.Context, input dto.DeleteOrderInput) error {
	order, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if order == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.OrderPresenterInput{
		Writer: input.Writer,
		Result: order,
	})

	return nil
}
