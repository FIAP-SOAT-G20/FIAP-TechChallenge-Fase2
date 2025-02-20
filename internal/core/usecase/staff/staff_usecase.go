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

func (s staffUseCase) List(ctx context.Context, input dto.ListStaffsInput) ([]*entity.Staff, int64, error) {
	role := valueobject.ToStaffRole(input.Role)
	staffs, total, err := s.gateway.FindAll(ctx, input.Name, role, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return staffs, total, nil
}

func (s staffUseCase) Create(ctx context.Context, input dto.CreateStaffInput) (*entity.Staff, error) {
	role := valueobject.ToStaffRole(input.Role)
	if role == valueobject.UNDEFINED_ROLE {
		return nil, domain.NewValidationError(errors.New("Invalid role"))
	}

	staff := entity.NewStaff(input.Name, role)

	if err := s.gateway.Create(ctx, staff); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

func (s staffUseCase) Get(ctx context.Context, input dto.GetStaffInput) (*entity.Staff, error) {
	staff, err := s.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}

	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	return staff, nil
}

func (s staffUseCase) Update(ctx context.Context, input dto.UpdateStaffInput) (*entity.Staff, error) {
	staff, err := s.gateway.FindByID(ctx, input.ID)
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

	if err := s.gateway.Update(ctx, staff); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

func (s staffUseCase) Delete(ctx context.Context, input dto.DeleteStaffInput) (*entity.Staff, error) {
	staff, err := s.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domain.NewInternalError(err)
	}
	if staff == nil {
		return nil, domain.NewNotFoundError(domain.ErrNotFound)
	}

	if err := s.gateway.Delete(ctx, input.ID); err != nil {
		return nil, domain.NewInternalError(err)
	}

	return staff, nil
}

// NewStaffUseCase creates a new StaffUseCase
func NewStaffUseCase(gateway port.StaffGateway) port.StaffUseCase {
	return &staffUseCase{
		gateway: gateway,
	}
}
