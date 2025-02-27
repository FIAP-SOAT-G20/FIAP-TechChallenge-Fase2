package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
)

type CategoryUsecaseMock struct {
	mock.Mock
}

func (m *CategoryUsecaseMock) Create(category *entity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *CategoryUsecaseMock) GetByID(id uint64) (*entity.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Category), args.Error(1)
}

func (m *CategoryUsecaseMock) List(name string, page, limit int) ([]entity.Category, int64, error) {
	args := m.Called(name, page, limit)
	return args.Get(0).([]entity.Category), args.Get(1).(int64), args.Error(2)
}

func (m *CategoryUsecaseMock) Update(category *entity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *CategoryUsecaseMock) Delete(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
} 