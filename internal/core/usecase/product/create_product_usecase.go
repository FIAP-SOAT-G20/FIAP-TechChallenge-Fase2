package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createProductUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

func NewCreateProductUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.CreateProductUseCase {
	return &createProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

func (uc *createProductUseCase) Execute(ctx context.Context, input dto.CreateProductInput) error {
	product := entity.NewProduct(input.Name, input.Description, input.Price, input.CategoryID)

	if err := uc.gateway.Create(ctx, product); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.ProductPresenterInput{
		Writer: input.Writer,
		Result: product,
	})
	return nil
}
