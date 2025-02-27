package port

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
)

type CategoryDatasourcePort interface {
	Insert(category *entity.Category) error
	GetByID(id uint64) (*entity.Category, error)
	GetAll(name string, page, limit int) ([]entity.Category, int64, error)
	Update(category *entity.Category) error
	Delete(id uint64) error
}
