package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateProductUseCase struct {
	gateway   port.ProductGateway
	presenter port.ProductPresenter
}

func NewUpdateProductUseCase(gateway port.ProductGateway, presenter port.ProductPresenter) port.UpdateProductUseCase {
	return &updateProductUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

func (uc *updateProductUseCase) Execute(ctx context.Context, input dto.UpdateProductInput) error {
	// Busca o produto existente
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if product == nil {
		return domain.NewNotFoundError("produto não encontrado")
	}

	// Atualiza o produto
	if err := product.Update(input.Name, input.Description, input.Price, input.CategoryID); err != nil {
		return domain.NewValidationError(err)
	}

	// Persiste as alterações
	if err := uc.gateway.Update(ctx, product); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(port.ProductPresenterDTO{
		Writer: input.Writer,
		Result: product,
	})
	return nil
}
