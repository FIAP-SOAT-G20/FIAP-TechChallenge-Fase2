package product

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/domain/errors"
	"tech-challenge-2-app-example/internal/core/dto"
	mockport "tech-challenge-2-app-example/internal/core/port/mocks"
)

func TestGetProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	useCase := NewGetProductUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	mockProduct := &entity.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should get product successfully",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(mockProduct, nil)

				mockPresenter.EXPECT().
					ToResponse(mockProduct).
					Return(dto.ProductResponse{
						ID:          mockProduct.ID,
						Name:        mockProduct.Name,
						Description: mockProduct.Description,
						Price:       mockProduct.Price,
						CategoryID:  mockProduct.CategoryID,
					})
			},
			expectError: false,
		},
		{
			name: "should return not found error when product doesn't exist",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &errors.NotFoundError{},
		},
		{
			name: "should return internal error when gateway fails",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &errors.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			response, err := useCase.Execute(ctx, tt.id)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, mockProduct.ID, response.ID)
				assert.Equal(t, mockProduct.Name, response.Name)
				assert.Equal(t, mockProduct.Description, response.Description)
				assert.Equal(t, mockProduct.Price, response.Price)
				assert.Equal(t, mockProduct.CategoryID, response.CategoryID)
			}
		})
	}
}
