package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type getStaffUseCase struct {
	gateway   port.StaffGateway
	presenter port.StaffPresenter
}

// NewGetStaffUseCase creates a new GetStaffUseCase
func NewGetStaffUseCase(gateway port.StaffGateway, presenter port.StaffPresenter) port.GetStaffUseCase {
	return &getStaffUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute gets a staff
func (uc *getStaffUseCase) Execute(ctx context.Context, input dto.GetStaffInput) error {
	staff, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return domain.NewInternalError(err)
	}

	if staff == nil {
		return domain.NewNotFoundError(domain.ErrNotFound)
	}

	uc.presenter.Present(dto.StaffPresenterInput{
		Writer: input.Writer,
		Result: staff,
	})
	return nil
}
