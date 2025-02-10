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

func TestUpdateProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)

	useCase := NewUpdateProductUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	existingProduct := &entity.Product{
		ID:          1,
		Name:        "Old Name",
		Description: "Old Description",
		Price:       10.0,
		CategoryID:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name        string
		id          uint64
		input       dto.ProductRequest
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should update product successfully",
			id:   1,
			input: dto.ProductRequest{
				Name:        "New Name",
				Description: "New Description",
				Price:       20.0,
				CategoryID:  2,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingProduct, nil)

				mockGateway.EXPECT().
					Update(ctx, gomock.Any()).
					DoAndReturn(func(_ context.Context, p *entity.Product) error {
						assert.Equal(t, "New Name", p.Name)
						assert.Equal(t, "New Description", p.Description)
						assert.Equal(t, 20.0, p.Price)
						assert.Equal(t, uint64(2), p.CategoryID)
						return nil
					})

				mockPresenter.EXPECT().
					ToResponse(gomock.Any()).
					Return(dto.ProductResponse{
						ID:          1,
						Name:        "New Name",
						Description: "New Description",
						Price:       20.0,
						CategoryID:  2,
					})
			},
			expectError: false,
		},
		{
			name: "should return error when product not found",
			id:   1,
			input: dto.ProductRequest{
				Name:        "New Name",
				Description: "New Description",
				Price:       20.0,
				CategoryID:  2,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   errors.NewNotFoundError("produto não encontrado"),
		},
		{
			name: "should return error when validation fails",
			id:   1,
			input: dto.ProductRequest{
				Name:        "", // invalid empty name
				Description: "New Description",
				Price:       20.0,
				CategoryID:  2,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingProduct, nil)
			},
			expectError: true,
			errorType:   &errors.ValidationError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			response, err := useCase.Execute(ctx, tt.id, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tt.input.Name, response.Name)
				assert.Equal(t, tt.input.Description, response.Description)
				assert.Equal(t, tt.input.Price, response.Price)
				assert.Equal(t, tt.input.CategoryID, response.CategoryID)
			}
		})
	}
}
