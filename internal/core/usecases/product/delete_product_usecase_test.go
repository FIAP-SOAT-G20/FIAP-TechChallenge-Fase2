package product

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"tech-challenge-2-app-example/internal/core/domain/entity"
	mockport "tech-challenge-2-app-example/internal/core/port/mocks"
)

func TestDeleteProductUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockport.NewMockProductRepository(ctrl)
	useCase := NewDeleteProductUseCase(mockRepo)

	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
	}{
		{
			name: "should delete product successfully",
			id:   1,
			setupMocks: func() {
				mockRepo.EXPECT().FindByID(gomock.Any(), uint64(1)).Return(&entity.Product{}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), uint64(1)).Return(nil)
			},
			expectError: false,
		},
		{
			name: "should fail when product not found",
			id:   1,
			setupMocks: func() {
				mockRepo.EXPECT().FindByID(gomock.Any(), uint64(1)).Return(nil, nil)
			},
			expectError: true,
		},
		{
			name: "should fail when repository fails",
			id:   1,
			setupMocks: func() {
				mockRepo.EXPECT().FindByID(gomock.Any(), uint64(1)).Return(nil, errors.New("database error"))
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := useCase.Execute(context.Background(), tt.id)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
