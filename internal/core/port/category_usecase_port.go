package port

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
)

type CategoryUsecasePort interface {
	Create(category *entity.Category) error
	GetByID(id uint64) (*entity.Category, error)
	List(name string, page, limit int) ([]entity.Category, int64, error)
	Update(category *entity.Category) error
	Delete(id uint64) error
} 