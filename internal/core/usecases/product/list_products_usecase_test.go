package product

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	"tech-challenge-2-app-example/internal/core/dto"
	mockport "tech-challenge-2-app-example/internal/core/port/mocks"
)

func TestListProductsUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockport.NewMockProductRepository(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	useCase := NewListProductsUseCase(mockRepo, mockPresenter)
	ctx := context.Background()

	products := []*entity.Product{
		{
			ID:          1,
			Name:        "Product 1",
			Description: "Description 1",
			Price:       99.99,
			CategoryID:  1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	tests := []struct {
		name        string
		request     dto.ProductListRequest
		setupMocks  func()
		expectError bool
	}{
		{
			name: "should list products successfully",
			request: dto.ProductListRequest{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockRepo.EXPECT().
					FindAll(ctx, "", uint64(0), 1, 10).
					Return(products, int64(1), nil)

				mockPresenter.EXPECT().
					ToPaginatedResponse(products, int64(1), 1, 10).
					Return(dto.PaginatedResponse{
						Total:    1,
						Page:     1,
						Limit:    10,
						Products: []dto.ProductResponse{},
					})
			},
			expectError: false,
		},
		{
			name: "should return error with invalid page",
			request: dto.ProductListRequest{
				Page:  0,
				Limit: 10,
			},
			setupMocks:  func() {},
			expectError: true,
		},
		{
			name: "should return error with invalid limit",
			request: dto.ProductListRequest{
				Page:  1,
				Limit: 101,
			},
			setupMocks:  func() {},
			expectError: true,
		},
		{
			name: "should return error when repository fails",
			request: dto.ProductListRequest{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockRepo.EXPECT().
					FindAll(ctx, "", uint64(0), 1, 10).
					Return(nil, int64(0), assert.AnError)
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
				assert.Equal(t, tt.request.Page, response.Page)
				assert.Equal(t, tt.request.Limit, response.Limit)
			}
		})
	}
}
