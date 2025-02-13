package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
)

type ListStaffsUseCase interface {
	Execute(ctx context.Context, input dto.ListStaffsInput) error
}

type CreateStaffUseCase interface {
	Execute(ctx context.Context, input dto.CreateStaffInput) error
}

type GetStaffUseCase interface {
	Execute(ctx context.Context, input dto.GetStaffInput) error
}

type UpdateStaffUseCase interface {
	Execute(ctx context.Context, input dto.UpdateStaffInput) error
}

type DeleteStaffUseCase interface {
	Execute(ctx context.Context, input dto.DeleteStaffInput) error
}
