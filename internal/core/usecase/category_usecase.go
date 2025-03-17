package usecase

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type CategoryUsecase struct {
	categoryGateway port.CategoryGateway
}

func NewCategoryUsecase(categoryGateway port.CategoryGateway) port.CategoryUsecasePort {
	return &CategoryUsecase{
		categoryGateway: categoryGateway,
	}
}

func (cu *CategoryUsecase) Create(category *entity.Category) error {
	return cu.categoryGateway.Insert(category)
}

func (cu *CategoryUsecase) GetByID(id uint64) (*entity.Category, error) {
	return cu.categoryGateway.GetByID(id)
}

func (cu *CategoryUsecase) List(name string, page, limit int) ([]entity.Category, int64, error) {
	return cu.categoryGateway.GetAll(name, page, limit)
}

func (cu *CategoryUsecase) Update(category *entity.Category) error {
	_, err := cu.categoryGateway.GetByID(category.ID)
	if err != nil {
		return domain.ErrNotFound
	}

	return cu.categoryGateway.Update(category)
}

func (cu *CategoryUsecase) Delete(id uint64) error {
	_, err := cu.categoryGateway.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return cu.categoryGateway.Delete(id)
}
