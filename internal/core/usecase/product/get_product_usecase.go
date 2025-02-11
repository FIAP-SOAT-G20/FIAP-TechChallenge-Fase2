package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type getProductUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

func NewGetProductUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.GetProductUseCase {
	return &getProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

func (uc *getProductUseCase) Execute(ctx context.Context, input dto.GetProductInput) error {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}

	if product == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	uc.presenter.Present(port.ProductPresenterDTO{
		Writer: input.Writer,
		Result: product,
	})
	return nil
}
