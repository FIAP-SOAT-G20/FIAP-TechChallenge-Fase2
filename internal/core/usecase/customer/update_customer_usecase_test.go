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

func TestUpdateCustomerUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockCustomerGateway(ctrl)
	mockPresenter := mockport.NewMockCustomerPresenter(ctrl)
	mockWriter := mockdto.NewMockResponseWriter(ctrl)

	useCase := NewUpdateCustomerUseCase(mockGateway, mockPresenter)
	ctx := context.Background()

	currentTime := time.Now()
	existingCustomer := &entity.Customer{
		ID:        1,
		Name:      "Old Name",
		Email:     "old.email@email.com",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	tests := []struct {
		name        string
		input       dto.UpdateCustomerInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should update customer successfully",
			input: dto.UpdateCustomerInput{
				ID:     1,
				Name:   "New Name",
				Email:  "new.name@email.com",
				Writer: mockWriter,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingCustomer, nil)

				mockGateway.EXPECT().
					Update(ctx, gomock.Any()).
					DoAndReturn(func(_ context.Context, p *entity.Customer) error {
						assert.Equal(t, "New Name", p.Name)
						assert.Equal(t, "new.name@email.com", p.Email)
						return nil
					})

				mockPresenter.EXPECT().
					Present(gomock.Any())
			},
			expectError: false,
		},
		{
			name: "should return error when customer not found",
			input: dto.UpdateCustomerInput{
				ID:     1,
				Name:   "New Name",
				Email:  "new.name@email.com",
				Writer: mockWriter,
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
			input: dto.UpdateCustomerInput{
				ID:     1,
				Name:   "New Name",
				Email:  "new.name@email.com",
				Writer: mockWriter,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingCustomer, nil)

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
