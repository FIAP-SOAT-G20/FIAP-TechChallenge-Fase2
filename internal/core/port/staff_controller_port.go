package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

type StaffController interface {
	List(ctx context.Context, presenter Presenter, input dto.ListStaffsInput) error
	Create(ctx context.Context, presenter Presenter, input dto.CreateStaffInput) error
	Get(ctx context.Context, presenter Presenter, input dto.GetStaffInput) error
	Update(ctx context.Context, presenter Presenter, input dto.UpdateStaffInput) error
	Delete(ctx context.Context, presenter Presenter, input dto.DeleteStaffInput) error
}
