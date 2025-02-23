package usecase

import (
	"context"
	"errors"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type staffUseCase struct {
	gateway port.StaffGateway
}

// NewStaffUseCase creates a new StaffUseCase
func NewStaffUseCase(gateway port.StaffGateway) port.StaffUseCase {
	return &staffUseCase{gateway: gateway}
}

// List returns a list of staffs
func (uc *staffUseCase) List(ctx context.Context, i dto.ListStaffsInput) ([]*entity.Staff, int64, error) {
	staffs, total, err := uc.gateway.FindAll(ctx, i.Name, i.Role, i.Page, i.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return staffs, total, nil
}

// Create creates a new staff
func (uc *staffUseCase) Create(ctx context.Context, i dto.CreateStaffInput) (*entity.Staff, error) {
	if i.Role == valueobject.UNDEFINED {
		return nil, domain.NewValidationError(errors.New("Invalid role"))
	}

	staff := entity.NewStaff(i.Name, i.Role)

	if err := uc.gateway.Create(ctx, staff); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

// Get returns a staff by ID
func (uc *staffUseCase) Get(ctx context.Context, i dto.GetStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return staff, nil
}

// Update updates a staff
func (uc *staffUseCase) Update(ctx context.Context, i dto.UpdateStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if i.Role == valueobject.UNDEFINED {
		return nil, domain.NewValidationError(errors.New("Invalid role"))
	}

	staff.Update(i.Name, i.Role)

	if err := uc.gateway.Update(ctx, staff); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

// Delete deletes a staff
func (uc *staffUseCase) Delete(ctx context.Context, i dto.DeleteStaffInput) (*entity.Staff, error) {
	staff, err := uc.gateway.FindByID(ctx, i.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := uc.gateway.Delete(ctx, i.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}
