package port

import (
	"context"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
)

type StaffUseCase interface {
	List(ctx context.Context, input dto.ListStaffsInput) ([]*entity.Staff, int64, error)
	Create(ctx context.Context, input dto.CreateStaffInput) (*entity.Staff, error)
	Get(ctx context.Context, input dto.GetStaffInput) (*entity.Staff, error)
	Update(ctx context.Context, input dto.UpdateStaffInput) (*entity.Staff, error)
	Delete(ctx context.Context, input dto.DeleteStaffInput) (*entity.Staff, error)
}
