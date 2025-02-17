package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteOrderProductUseCase struct {
	gateway   port.OrderProductGateway
	presenter port.OrderProductPresenter
}

// NewDeleteOrderProductUseCase creates a new DeleteOrderProductUseCase
func NewDeleteOrderProductUseCase(gateway port.OrderProductGateway, presenter port.OrderProductPresenter) port.DeleteOrderProductUseCase {
	return &deleteOrderProductUseCase{gateway, presenter}
}

// Execute deletes a order
func (uc *deleteOrderProductUseCase) Execute(ctx context.Context, input dto.DeleteOrderProductInput) error {
	order, err := uc.gateway.FindByID(ctx, input.OrderID, input.ProductID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if order == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.OrderID, input.ProductID); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Result: order,
	})

	return nil
}
