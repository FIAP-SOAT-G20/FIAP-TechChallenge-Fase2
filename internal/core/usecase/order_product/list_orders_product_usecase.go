package orderproduct

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listOrderProductsUseCase struct {
	gateway   port.OrderProductGateway
	presenter port.OrderProductPresenter
}

// NewListOrderProductsUseCase creates a new ListOrderProductsUseCase
func NewListOrderProductsUseCase(gateway port.OrderProductGateway, presenter port.OrderProductPresenter) port.ListOrderProductsUseCase {
	return &listOrderProductsUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute lists all orderProducts
func (uc *listOrderProductsUseCase) Execute(ctx context.Context, input dto.ListOrderProductsInput) error {
	orderProducts, total, err := uc.gateway.FindAll(ctx, input.OrderID, input.ProductID, input.Page, input.Limit)
	if err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.OrderProductPresenterInput{
		Writer: input.Writer,
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: orderProducts,
	})
	return nil
}
