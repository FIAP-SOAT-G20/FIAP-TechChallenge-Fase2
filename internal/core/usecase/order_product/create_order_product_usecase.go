package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createOrderProductUseCase struct {
	gateway   port.OrderProductGateway
	presenter port.OrderProductPresenter
}

// NewCreateOrderProductUseCase creates a new CreateOrderProductUseCase
func NewCreateOrderProductUseCase(gateway port.OrderProductGateway, presenter port.OrderProductPresenter) port.CreateOrderProductUseCase {
	return &createOrderProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute creates a new OrderProduct
func (uc *createOrderProductUseCase) Execute(ctx context.Context, input dto.CreateOrderProductInput) error {
	orderProduct := entity.NewOrderProduct(input.OrderID, input.ProductID, input.Quantity)

	if err := uc.gateway.Create(ctx, orderProduct); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Result: orderProduct,
	})
	return nil
}
