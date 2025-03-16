package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
)

func TestAuthUseCase_Authenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		// Arrange
		mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
		mockJWTService := mockport.NewMockJWTService(ctrl)
		useCase := usecase.NewAuthUseCase(mockCustomerUseCase, mockJWTService)

		ctx := context.Background()
		input := dto.AuthenticateInput{
			CPF: "12345678901",
		}

		mockCustomer := &entity.Customer{
			ID:    1,
			Name:  "Test Customer",
			Email: "test@example.com",
			CPF:   "12345678901",
		}
		mockToken := "test-jwt-token"
		mockExpiresIn := int64(86400)

		// Set up expectations
		mockCustomerUseCase.EXPECT().
			FindByCPF(ctx, dto.FindCustomerByCPFInput(input)).
			Return(mockCustomer, nil)

		mockJWTService.EXPECT().
			GenerateToken(mockCustomer.ID, mockCustomer.CPF, mockCustomer.Name).
			Return(mockToken, mockExpiresIn, nil)

		// Act
		customer, token, expiresIn, err := useCase.Authenticate(ctx, input)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, customer)
		assert.Equal(t, mockCustomer, customer)
		assert.Equal(t, mockToken, token)
		assert.Equal(t, mockExpiresIn, expiresIn)
	})

	t.Run("customer_not_found", func(t *testing.T) {
		// Arrange
		mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
		mockJWTService := mockport.NewMockJWTService(ctrl)
		useCase := usecase.NewAuthUseCase(mockCustomerUseCase, mockJWTService)

		ctx := context.Background()
		input := dto.AuthenticateInput{
			CPF: "12345678901",
		}

		expectedErr := errors.New("customer not found")

		// Set up expectations
		mockCustomerUseCase.EXPECT().
			FindByCPF(ctx, dto.FindCustomerByCPFInput(input)).
			Return(nil, expectedErr)

		// Act
		customer, token, expiresIn, err := useCase.Authenticate(ctx, input)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, customer)
		assert.Empty(t, token)
		assert.Equal(t, int64(0), expiresIn)
	})

	t.Run("token_generation_error", func(t *testing.T) {
		// Arrange
		mockCustomerUseCase := mockport.NewMockCustomerUseCase(ctrl)
		mockJWTService := mockport.NewMockJWTService(ctrl)
		useCase := usecase.NewAuthUseCase(mockCustomerUseCase, mockJWTService)

		ctx := context.Background()
		input := dto.AuthenticateInput{
			CPF: "12345678901",
		}

		mockCustomer := &entity.Customer{
			ID:    1,
			Name:  "Test Customer",
			Email: "test@example.com",
			CPF:   "12345678901",
		}
		expectedErr := errors.New("token generation error")

		// Set up expectations
		mockCustomerUseCase.EXPECT().
			FindByCPF(ctx, dto.FindCustomerByCPFInput(input)).
			Return(mockCustomer, nil)

		mockJWTService.EXPECT().
			GenerateToken(mockCustomer.ID, mockCustomer.CPF, mockCustomer.Name).
			Return("", int64(0), expectedErr)

		// Act
		customer, token, expiresIn, err := useCase.Authenticate(ctx, input)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, customer)
		assert.Empty(t, token)
		assert.Equal(t, int64(0), expiresIn)
	})
}
