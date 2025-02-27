package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
)

type CategoryDatasourceMock struct {
	mock.Mock
}

func (m *CategoryDatasourceMock) Insert(category *entity.Category) error {
	args := m.Called(category)
	return args.Get(0).(error)
}

func (m *CategoryDatasourceMock) GetByID(id uint64) (*entity.Category, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Category), args.Get(1).(error)
}

func (m *CategoryDatasourceMock) GetAll(name string, page, limit int) ([]entity.Category, int64, error) {
	args := m.Called(name, page, limit)
	return args.Get(0).([]entity.Category), args.Get(1).(int64), args.Get(2).(error)
}

func (m *CategoryDatasourceMock) Update(category *entity.Category) error {
	args := m.Called(category)
	return args.Get(0).(error)
}

func (m *CategoryDatasourceMock) Delete(id uint64) error {
	args := m.Called(id)
	return args.Get(0).(error)
} 