package usecase

import (
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSignInUsecase_GetByCPF(t *testing.T) {
	mockCustomerDatasource := new(mocks.CustomerDatasourceMock)
	usecase := NewSignInUsecase(mockCustomerDatasource)

	t.Run("success", func(t *testing.T) {
		expectedCustomer := &entity.Customer{
			ID:   1,
			Name: "Test Customer",
			CPF:  "12345678900",
		}

		mockCustomerDatasource.On("GetByCPF", "12345678900").Return(expectedCustomer, nil)

		customer, err := usecase.GetByCPF("12345678900")

		assert.NoError(t, err)
		assert.Equal(t, expectedCustomer, customer)
		mockCustomerDatasource.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockCustomerDatasource.On("GetByCPF", "99999999999").Return(nil, entity.ErrNotFound)

		customer, err := usecase.GetByCPF("99999999999")

		assert.Error(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, customer)
		mockCustomerDatasource.AssertExpectations(t)
	})

	t.Run("invalid cpf", func(t *testing.T) {
		mockCustomerDatasource.On("GetByCPF", "invalid").Return(nil, entity.ErrInvalidInput)

		customer, err := usecase.GetByCPF("invalid")

		assert.Error(t, err)
		assert.Equal(t, entity.ErrInvalidInput, err)
		assert.Nil(t, customer)
		mockCustomerDatasource.AssertExpectations(t)
	})
} 