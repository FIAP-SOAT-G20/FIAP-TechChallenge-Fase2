package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type StaffController interface {
	List(ctx context.Context, presenter Presenter, input dto.ListStaffsInput) ([]byte, error)
	Create(ctx context.Context, presenter Presenter, input dto.CreateStaffInput) ([]byte, error)
	Get(ctx context.Context, presenter Presenter, input dto.GetStaffInput) ([]byte, error)
	Update(ctx context.Context, presenter Presenter, input dto.UpdateStaffInput) ([]byte, error)
	Delete(ctx context.Context, presenter Presenter, input dto.DeleteStaffInput) ([]byte, error)
}
