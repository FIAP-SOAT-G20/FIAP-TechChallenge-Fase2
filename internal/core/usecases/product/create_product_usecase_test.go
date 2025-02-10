package product

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/dto"
	mockport "tech-challenge-2-app-example/internal/core/port/mocks"
)

func TestCreateProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	useCase := NewCreateProductUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	expectedProduct := &entity.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
	}

	tests := []struct {
		name        string
		request     dto.ProductRequest
		setupMocks  func()
		expectError bool
	}{
		{
			name: "should create product successfully",
			request: dto.ProductRequest{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)

				mockPresenter.EXPECT().
					ToResponse(gomock.Any()).
					Return(dto.ProductResponse{
						ID:          1,
						Name:        expectedProduct.Name,
						Description: expectedProduct.Description,
						Price:       expectedProduct.Price,
						CategoryID:  expectedProduct.CategoryID,
					})
			},
			expectError: false,
		},
		{
			name: "should return error when validation fails",
			request: dto.ProductRequest{
				Name:        "",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks:  func() {},
			expectError: true,
		},
		{
			name: "should return error when gateway fails",
			request: dto.ProductRequest{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			response, err := useCase.Execute(ctx, tt.request)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tt.request.Name, response.Name)
				assert.Equal(t, tt.request.Description, response.Description)
				assert.Equal(t, tt.request.Price, response.Price)
				assert.Equal(t, tt.request.CategoryID, response.CategoryID)
			}
		})
	}
}
