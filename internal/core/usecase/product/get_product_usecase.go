package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
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

func (uc *getProductUseCase) Execute(ctx context.Context, id uint64) (*usecase.ProductOutput, error) {
	product, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if product == nil {
		return nil, domain.NewNotFoundError("Produto não encontrado")
	}

	output := uc.presenter.ToOutput(product)
	return output, nil
}
