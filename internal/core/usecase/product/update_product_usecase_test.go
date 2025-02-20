package product_test

import (
	"context"
	"go.uber.org/mock/gomock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	mockdto "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase/product"
)

func TestUpdateProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)
	useCase := product.NewUpdateProductUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	existingProduct := &entity.Product{
		ID:          1,
		Name:        "Old Name",
		Description: "Old Description",
		Price:       10.0,
		CategoryID:  1,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	tests := []struct {
		name        string
		input       dto.UpdateProductInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should update product successfully",
			input: dto.UpdateProductInput{
				ID:          1,
				Name:        "New Name",
				Description: "New Description",
				Price:       20.0,
				CategoryID:  2,
				Writer:      mockWriter,
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
					Present(gomock.Any())
			},
			expectError: false,
		},
		{
			name: "should return error when product not found",
			input: dto.UpdateProductInput{
				ID:          1,
				Name:        "New Name",
				Description: "New Description",
				Price:       20.0,
				CategoryID:  2,
				Writer:      mockWriter,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when gateway update fails",
			input: dto.UpdateProductInput{
				ID:          1,
				Name:        "New Name",
				Description: "New Description",
				Price:       20.0,
				CategoryID:  2,
				Writer:      mockWriter,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingProduct, nil)

				mockGateway.EXPECT().
					Update(ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
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
