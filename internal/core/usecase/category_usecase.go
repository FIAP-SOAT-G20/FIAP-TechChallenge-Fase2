package usecase

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
)

type CategoryUsecase struct {
	categoryDatasource port.CategoryDatasourcePort
}

func NewCategoryUsecase(categoryDatasource port.CategoryDatasourcePort) port.CategoryUsecasePort {
	return &CategoryUsecase{
		categoryDatasource: categoryDatasource,
	}
}

func (cu *CategoryUsecase) Create(category *entity.Category) error {
	return cu.categoryDatasource.Insert(category)
}

func (cu *CategoryUsecase) GetByID(id uint64) (*entity.Category, error) {
	return cu.categoryDatasource.GetByID(id)
}

func (cu *CategoryUsecase) List(name string, page, limit int) ([]entity.Category, int64, error) {
	return cu.categoryDatasource.GetAll(name, page, limit)
}

func (cu *CategoryUsecase) Update(category *entity.Category) error {
	_, err := cu.categoryDatasource.GetByID(category.ID)
	if err != nil {
		return entity.ErrNotFound
	}

	return cu.categoryDatasource.Update(category)
}

func (cu *CategoryUsecase) Delete(id uint64) error {
	_, err := cu.categoryDatasource.GetByID(id)
	if err != nil {
		return entity.ErrNotFound
	}

	return cu.categoryDatasource.Delete(id)
}
