package staff

import (
	"context"
	"errors"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type staffUseCase struct {
	gateway port.StaffGateway
}

// NewStaffUseCase creates a new StaffUseCase
func NewStaffUseCase(gateway port.StaffGateway) port.StaffUseCase {
	return &staffUseCase{
		gateway: gateway,
	}
}

// List returns a list of staffs
func (uc staffUseCase) List(ctx context.Context, input dto.ListStaffsInput) ([]*entity.Staff, int64, error) {
	role := valueobject.ToStaffRole(input.Role)
	staffs, total, err := uc.gateway.FindAll(ctx, input.Name, role, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return staffs, total, nil
}

// Create creates a new staff
func (uc staffUseCase) Create(ctx context.Context, input dto.CreateStaffInput) (*entity.Staff, error) {
	role := valueobject.ToStaffRole(input.Role)
	if role == valueobject.UNDEFINED_ROLE {
		return nil, domain.NewValidationError(errors.New("Invalid role"))
	}

	staff := entity.NewStaff(input.Name, role)

	if err := uc.gateway.Create(ctx, staff); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

// Get returns a staff by ID
func (uc staffUseCase) Get(ctx context.Context, input dto.GetStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return staff, nil
}

// Update updates a staff
func (uc staffUseCase) Update(ctx context.Context, input dto.UpdateStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	role := valueobject.ToStaffRole(input.Role)
	if role == valueobject.UNDEFINED_ROLE {
		return nil, domain.NewValidationError(errors.New("Invalid role"))
	}

	staff.Update(input.Name, role)

	if err := uc.gateway.Update(ctx, staff); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

// Delete deletes a staff
func (uc staffUseCase) Delete(ctx context.Context, input dto.DeleteStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, input.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}
