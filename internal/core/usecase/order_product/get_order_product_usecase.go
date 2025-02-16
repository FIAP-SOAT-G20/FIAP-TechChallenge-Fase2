package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type getOrderProductUseCase struct {
	gateway   port.OrderProductGateway
	presenter port.OrderProductPresenter
}

// NewGetOrderProductUseCase creates a new GetOrderProductUseCase
func NewGetOrderProductUseCase(gateway port.OrderProductGateway, presenter port.OrderProductPresenter) port.GetOrderProductUseCase {
	return &getOrderProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute gets a orderProduct
func (uc *getOrderProductUseCase) Execute(ctx context.Context, input dto.GetOrderProductInput) error {
	orderProduct, err := uc.gateway.FindByID(ctx, input.OrderID, input.ProductID)
	if err != nil {
		return domain.NewInternalError(err)
	}

	if orderProduct == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	uc.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Result: orderProduct,
	})
	return nil
}
