package staff

import (
	"context"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type staffUseCase struct {
	gateway port.StaffGateway
}

func (s staffUseCase) List(ctx context.Context, input dto.ListStaffsInput) ([]*entity.Staff, int64, error) {
	var role entity.Role
	if input.Role != "" {
		role = entity.Role(input.Role)
	}
	staffs, total, err := s.gateway.FindAll(ctx, input.Name, role, input.Page, input.Limit)
	if err != nil {
		return nil, 0, domain.NewInternalError(err)
	}

	return staffs, total, nil
}

func (s staffUseCase) Create(ctx context.Context, input dto.CreateStaffInput) (*entity.Staff, error) {
	staff := entity.NewStaff(input.Name, entity.Role(input.Role))

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

	staff.Update(input.Name, entity.Role(input.Role))

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
