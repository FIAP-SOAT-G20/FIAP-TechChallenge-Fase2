package product

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestListProductsUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := NewListProductsUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	mockProducts := []*entity.Product{
		{
			ID:          1,
			Name:        "Test Product 1",
			Description: "Description 1",
			Price:       99.99,
			CategoryID:  1,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		},
		{
			ID:          2,
			Name:        "Test Product 2",
			Description: "Description 2",
			Price:       199.99,
			CategoryID:  1,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		},
	}

	tests := []struct {
		name        string
		input       dto.ListProductsInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should list products successfully",
			input: dto.ListProductsInput{
				Writer: mockWriter,
				Page:   1,
				Limit:  10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", uint64(0), 1, 10).
					Return(mockProducts, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.ProductPresenterInput{
						Writer: mockWriter,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockProducts,
					})
			},
			expectError: false,
		},
		{
			name: "should return error when page is invalid",
			input: dto.ListProductsInput{
				Writer: mockWriter,
				Page:   0,
				Limit:  10,
			},
			setupMocks:  func() {},
			expectError: true,
			errorType:   &domain.InvalidInputError{},
		},
		{
			name: "should return error when limit is too high",
			input: dto.ListProductsInput{
				Writer: mockWriter,
				Page:   1,
				Limit:  101,
			},
			setupMocks:  func() {},
			expectError: true,
			errorType:   &domain.InvalidInputError{},
		},
		{
			name: "should return error when repository fails",
			input: dto.ListProductsInput{
				Writer: mockWriter,
				Page:   1,
				Limit:  10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", uint64(0), 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by name",
			input: dto.ListProductsInput{
				Writer: mockWriter,
				Name:   "Test",
				Page:   1,
				Limit:  10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "Test", uint64(0), 1, 10).
					Return(mockProducts, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.ProductPresenterInput{
						Writer: mockWriter,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockProducts,
					})
			},
			expectError: false,
		},
		{
			name: "should filter by category",
			input: dto.ListProductsInput{
				Writer:     mockWriter,
				CategoryID: 1,
				Page:       1,
				Limit:      10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", uint64(1), 1, 10).
					Return(mockProducts, int64(2), nil)

				mockPresenter.EXPECT().
					Present(dto.ProductPresenterInput{
						Writer: mockWriter,
						Total:  int64(2),
						Page:   1,
						Limit:  10,
						Result: mockProducts,
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
