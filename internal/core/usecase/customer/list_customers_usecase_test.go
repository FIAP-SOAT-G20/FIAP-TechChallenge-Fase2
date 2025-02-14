package customer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestListCustomersUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockCustomerGateway(ctrl)
	mockPresenter := mockport.NewMockCustomerPresenter(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := NewListCustomersUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	mockCustomers := []*entity.Customer{
		{
			ID:        1,
			Name:      "Test Customer 1",
			Email:     "test.customer.1@email.com",
			CPF:       "12345678901",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			ID:        2,
			Name:      "Test Customer 2",
			Email:     "test.customer.2@email.com",
			CPF:       "12345678902",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}

	tests := []struct {
		name        string
		input       dto.ListCustomersInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should list products successfully",
			input: dto.ListCustomersInput{
				Writer: mockWriter,
				Page:   1,
				Limit:  10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", 1, 10).
					Return(mockCustomers, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.CustomerPresenterInput{
						Writer: mockWriter,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockCustomers,
					})
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListCustomersInput{
				Writer: mockWriter,
				Page:   1,
				Limit:  10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by name",
			input: dto.ListCustomersInput{
				Writer: mockWriter,
				Name:   "Test",
				Page:   1,
				Limit:  10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "Test", 1, 10).
					Return(mockCustomers, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.CustomerPresenterInput{
						Writer: mockWriter,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockCustomers,
					})
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := useCase.Execute(ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
