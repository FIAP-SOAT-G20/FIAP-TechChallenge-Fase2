package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type updateStaffUseCase struct {
	gateway   port.StaffGateway
	presenter port.StaffPresenter
}

// NewUpdateStaffUseCase creates a new UpdateStaffUseCase
func NewUpdateStaffUseCase(gateway port.StaffGateway, presenter port.StaffPresenter) port.UpdateStaffUseCase {
	return &updateStaffUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute updates a staff
func (uc *updateStaffUseCase) Execute(ctx context.Context, input dto.UpdateStaffInput) error {
	product, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if product == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	product.Update(input.Name, entity.Role(input.Role))

	if err := uc.gateway.Update(ctx, product); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.StaffPresenterInput{
		Writer: input.Writer,
		Result: product,
	})
	return nil
}
