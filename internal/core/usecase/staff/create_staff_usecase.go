package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type createStaffUseCase struct {
	gateway   port.StaffGateway
	presenter port.StaffPresenter
}

// NewCreateStaffUseCase creates a new CreateStaffUseCase
func NewCreateStaffUseCase(gateway port.StaffGateway, presenter port.StaffPresenter) port.CreateStaffUseCase {
	return &createStaffUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute creates a new staff
func (uc *createStaffUseCase) Execute(ctx context.Context, input dto.CreateStaffInput) error {
	staff := entity.NewStaff(input.Name, entity.Role(input.Role))

	if err := uc.gateway.Create(ctx, staff); err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.StaffPresenterInput{
		Writer: input.Writer,
		Result: staff,
	})
	return nil
}
