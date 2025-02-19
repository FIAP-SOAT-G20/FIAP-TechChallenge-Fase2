package staff

import (
	"context"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateStaffUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockStaffGateway(ctrl)
	useCase := NewStaffUseCase(mockGateway)
	ctx := context.Background()

	tests := []struct {
		name        string
		input       dto.CreateStaffInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create staff successfully",
			input: dto.CreateStaffInput{
				Name: "John Smith",
				Role: "COOK",
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreateStaffInput{
				Name: "John Smith",
				Role: "COOK",
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					Create(ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, err := useCase.Create(ctx, tt.input)

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

func TestUpdateStaffUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockStaffGateway(ctrl)

	useCase := NewStaffUseCase(mockGateway)
	ctx := context.Background()

	currentTime := time.Now()
	existingStaff := &entity.Staff{
		ID:        1,
		Name:      "John Smith",
		Role:      "COOK",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	tests := []struct {
		name        string
		input       dto.UpdateStaffInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should update staff successfully",
			input: dto.UpdateStaffInput{
				ID:   1,
				Name: "New Name",
				Role: "ATTENDANT",
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingStaff, nil)

				mockGateway.EXPECT().
					Update(ctx, gomock.Any()).
					DoAndReturn(func(_ context.Context, p *entity.Staff) error {
						assert.Equal(t, "New Name", p.Name)
						assert.Equal(t, "ATTENDANT", string(p.Role))
						return nil
					})
			},
			expectError: false,
		},
		{
			name: "should return error when staff not found",
			input: dto.UpdateStaffInput{
				ID:   1,
				Name: "New Name",
				Role: "ATTENDANT",
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
			input: dto.UpdateStaffInput{
				ID:   1,
				Name: "New Name",
				Role: "ATTENDANT",
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(existingStaff, nil)

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

			_, err := useCase.Update(ctx, tt.input)

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

func TestDeleteStaffUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockStaffGateway(ctrl)
	useCase := NewStaffUseCase(mockGateway)
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
					Return(&entity.Staff{}, nil)

				mockGateway.EXPECT().
					Delete(ctx, uint64(1)).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when staff doesn't exist",
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
					Return(&entity.Staff{}, nil)

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

			_, err := useCase.Delete(ctx, dto.DeleteStaffInput{ID: tt.id})

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

func TestGetStaffUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockStaffGateway(ctrl)
	useCase := NewStaffUseCase(mockGateway)
	ctx := context.Background()

	currentTime := time.Now()
	mockStaff := &entity.Staff{
		ID:        1,
		Name:      "Test Staff",
		Role:      "COOK",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
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
					Return(mockStaff, nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when staff doesn't exist",
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
			name: "should return internal error when gateway fails",
			id:   1,
			setupMocks: func() {
				mockGateway.EXPECT().
					FindByID(ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, err := useCase.Get(ctx, dto.GetStaffInput{
				ID: tt.id,
			})

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

func TestListStaffsUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGateway := mockport.NewMockStaffGateway(ctrl)
	useCase := NewStaffUseCase(mockGateway)
	ctx := context.Background()

	currentTime := time.Now()
	mockStaffs := []*entity.Staff{
		{
			ID:        1,
			Name:      "Test Staff 1",
			Role:      "COOK",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			ID:        2,
			Name:      "Test Staff 2",
			Role:      "COOK",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}

	tests := []struct {
		name        string
		input       dto.ListStaffsInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should list staffs successfully",
			input: dto.ListStaffsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				var role entity.Role
				mockGateway.EXPECT().
					FindAll(ctx, "", role, 1, 10).
					Return(mockStaffs, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListStaffsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				var role entity.Role
				mockGateway.EXPECT().
					FindAll(ctx, "", role, 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by name",
			input: dto.ListStaffsInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				var role entity.Role
				mockGateway.EXPECT().
					FindAll(ctx, "Test", role, 1, 10).
					Return(mockStaffs, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should filter by Role",
			input: dto.ListStaffsInput{
				Role:  "COOK",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				mockGateway.EXPECT().
					FindAll(ctx, "", entity.COOK, 1, 10).
					Return(mockStaffs, int64(2), nil)

			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			_, _, err := useCase.List(ctx, tt.input)

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
