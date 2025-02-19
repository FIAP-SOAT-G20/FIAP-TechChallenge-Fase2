package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateOrderProductUseCase struct {
	gateway   port.OrderProductGateway
	presenter port.OrderProductPresenter
}

// NewUpdateOrderProductUseCase creates a new UpdateOrderProductUseCase
func NewUpdateOrderProductUseCase(gateway port.OrderProductGateway, presenter port.OrderProductPresenter) port.UpdateOrderProductUseCase {
	return &updateOrderProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute updates a orderProduct
func (uc *updateOrderProductUseCase) Execute(ctx context.Context, input dto.UpdateOrderProductInput) error {
	orderProduct, err := uc.gateway.FindByID(ctx, input.OrderID, input.ProductID)
	if err != nil {
		return domain.NewInternalError(err)
	}

	if orderProduct == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	orderProduct.Update(input.Quantity)

	if err := uc.gateway.Update(ctx, orderProduct); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Result: orderProduct,
	})
	return nil
}
