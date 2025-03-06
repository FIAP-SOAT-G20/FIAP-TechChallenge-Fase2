package usecase_test

import (
	"context"
	"testing"
	"time"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

func (s *ProductUsecaseSuiteTest) TestProductsUseCase_List() {
	currentTime := time.Now()
	mockProducts := []*entity.Product{
		{
			ID:          1,
			Name:        "Test Product 1",
			Description: "Description 1",
			Price:       99.99,
			CategoryID:  1,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		},
		{
			ID:          2,
			Name:        "Test Product 2",
			Description: "Description 2",
			Price:       199.99,
			CategoryID:  1,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		},
	}

	tests := []struct {
		name        string
		input       dto.ListProductsInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should list products successfully",
			input: dto.ListProductsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", uint64(0), 1, 10).
					Return(mockProducts, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListProductsInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", uint64(0), 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by name",
			input: dto.ListProductsInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "Test", uint64(0), 1, 10).
					Return(mockProducts, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should filter by category",
			input: dto.ListProductsInput{
				CategoryID: 1,
				Page:       1,
				Limit:      10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", uint64(1), 1, 10).
					Return(mockProducts, int64(2), nil)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			products, total, err := s.useCase.List(s.ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, products)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, products)
				assert.Equal(t, len(mockProducts), len(products))
				assert.Equal(t, int64(2), total)
			}
		})
	}
}

func (s *ProductUsecaseSuiteTest) TestProductUseCase_Create() {
	tests := []struct {
		name        string
		input       dto.CreateProductInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create product successfully",
			input: dto.CreateProductInput{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					Create(s.ctx, gomock.Any()).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return error when gateway fails",
			input: dto.CreateProductInput{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				CategoryID:  1,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					Create(s.ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			product, err := s.useCase.Create(s.ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, product)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, tt.input.Name, product.Name)
				assert.Equal(t, tt.input.Description, product.Description)
				assert.Equal(t, tt.input.Price, product.Price)
				assert.Equal(t, tt.input.CategoryID, product.CategoryID)
			}
		})
	}
}

func (s *ProductUsecaseSuiteTest) TestProductUseCase_Get() {
	currentTime := time.Now()
	mockProduct := &entity.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		CategoryID:  1,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
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
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(mockProduct, nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when product doesn't exist",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return internal error when gateway fails",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			product, err := s.useCase.Get(s.ctx, dto.GetProductInput{
				ID: tt.id,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, product)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, mockProduct, product)
			}
		})
	}
}

func (s *ProductUsecaseSuiteTest) TestProductUseCase_Update() {
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
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(existingProduct, nil)

				s.mockGateway.EXPECT().
					Update(s.ctx, gomock.Any()).
					DoAndReturn(func(_ context.Context, p *entity.Product) error {
						assert.Equal(s.T(), "New Name", p.Name)
						assert.Equal(s.T(), "New Description", p.Description)
						assert.Equal(s.T(), 20.0, p.Price)
						assert.Equal(s.T(), uint64(2), p.CategoryID)
						return nil
					})
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
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
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
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(existingProduct, nil)

				s.mockGateway.EXPECT().
					Update(s.ctx, gomock.Any()).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			product, err := s.useCase.Update(s.ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, product)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, tt.input.Name, product.Name)
				assert.Equal(t, tt.input.Description, product.Description)
				assert.Equal(t, tt.input.Price, product.Price)
			}
		})
	}
}

func (s *ProductUsecaseSuiteTest) TestProductUseCase_Delete() {
	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should delete _test successfully",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(&entity.Product{}, nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when _test doesn't exist",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, nil)
			},
			expectError: true,
			errorType:   &domain.NotFoundError{},
		},
		{
			name: "should return error when gateway fails on find",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(nil, assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should return error when gateway fails on delete",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(&entity.Product{}, nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			product, err := s.useCase.Delete(s.ctx, dto.DeleteProductInput{ID: tt.id})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, product)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
			}
		})
	}
}
