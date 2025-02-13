package product

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type listStaffsUseCase struct {
	gateway   port.StaffGateway
	presenter port.StaffPresenter
}

// NewListStaffsUseCase creates a new ListStaffsUseCase
func NewListStaffsUseCase(gateway port.StaffGateway, presenter port.StaffPresenter) port.ListStaffsUseCase {
	return &listStaffsUseCase{
		gateway:   gateway,
		presenter: presenter,
	}
}

// Execute lists all staffs
func (uc *listStaffsUseCase) Execute(ctx context.Context, input dto.ListStaffsInput) error {
	products, total, err := uc.gateway.FindAll(ctx, input.Name, entity.Role(input.Role), input.Page, input.Limit)
	if err != nil {
		return domain.NewInternalError(err)
	}

	uc.presenter.Present(dto.StaffPresenterInput{
		Writer: input.Writer,
		Total:  total,
		Page:   input.Page,
		Limit:  input.Limit,
		Result: products,
	})
	return nil
}
