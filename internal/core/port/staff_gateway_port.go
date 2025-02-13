package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type StaffGateway interface {
	FindByID(ctx context.Context, id uint64) (*entity.Staff, error)
	FindAll(ctx context.Context, name string, role entity.Role, page, limit int) ([]*entity.Staff, int64, error)
	Create(ctx context.Context, staff *entity.Staff) error
	Update(ctx context.Context, staff *entity.Staff) error
	Delete(ctx context.Context, id uint64) error
}
