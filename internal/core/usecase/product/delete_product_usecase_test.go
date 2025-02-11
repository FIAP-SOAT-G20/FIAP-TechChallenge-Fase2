package product

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapters/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
)

func TestDeleteProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockProductGateway(ctrl)
	mockPresenter := mockport.NewMockProductPresenter(ctrl)
	useCase := NewDeleteProductUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should delete product successfully",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(&entity.Product{}, nil)

				mockGateway.EXPECT().
					Delete(ctx, uint64(1)).
					Return(nil)

				mockPresenter.EXPECT().
					Present(gomock.Any())
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
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when gateway fails on find",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when gateway fails on delete",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(&entity.Product{}, nil)

				mockGateway.EXPECT().
					Delete(ctx, uint64(1)).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := useCase.Execute(ctx, dto.DeleteProductInput{ID: tt.id})

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
