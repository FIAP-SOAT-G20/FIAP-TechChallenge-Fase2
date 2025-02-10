package product

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
)

func TestUpdateProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)

	useCase := NewUpdateProductUseCase(mockGateway, mockPresenter)
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

	mockOutput := &usecase.ProductOutput{
		ID:          1,
		Name:        "New Name",
		Description: "New Description",
		Price:       20.0,
		CategoryID:  2,
		CreatedAt:   currentTime.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   currentTime.Format("2006-01-02T15:04:05Z07:00"),
	}

	tests := []struct {
		name        string
		id          uint64
		input       usecase.UpdateProductInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should update product successfully",
			id:   1,
			input: usecase.UpdateProductInput{
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
					ToOutput(gomock.Any()).
					Return(mockOutput)
			},
			expectError: false,
		},
		{
			name: "should return error when product not found",
			id:   1,
			input: usecase.UpdateProductInput{
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
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when validation fails",
			id:   1,
			input: usecase.UpdateProductInput{
				Name:        "", // Nome vazio deve falhar na validação
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
			errorType:   &domain.ValidationError{},
		},
		{
			name: "should return error when gateway update fails",
			id:   1,
			input: usecase.UpdateProductInput{
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
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			output, err := useCase.Execute(ctx, tt.id, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, mockOutput, output)
			}
		})
	}
}
