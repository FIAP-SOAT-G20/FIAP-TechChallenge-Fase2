package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type deleteStaffUseCase struct {
	gateway   port.StaffGateway
	presenter port.StaffPresenter
}

// NewDeleteStaffUseCase creates a new DeleteStaffUseCase
func NewDeleteStaffUseCase(gateway port.StaffGateway, presenter port.StaffPresenter) port.DeleteStaffUseCase {
	return &deleteStaffUseCase{gateway, presenter}
}

// Execute deletes a staff
func (uc *deleteStaffUseCase) Execute(ctx context.Context, input dto.DeleteStaffInput) error {
	staff, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}
	if staff == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.StaffPresenterInput{
		Writer: input.Writer,
		Result: staff,
	})

	return nil
}
