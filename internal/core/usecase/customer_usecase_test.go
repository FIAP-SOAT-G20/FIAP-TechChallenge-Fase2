package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
)

func (s *CustomerUsecaseSuiteTest) TestCustomersUseCase_List() {
	currentTime := time.Now()
	mockCustomers := []*entity.Customer{
		{
			ID:        1,
			Name:      "Test Customer 1",
			Email:     "test.customer.1@email.com",
			CPF:       "12345678901",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			ID:        2,
			Name:      "Test Customer 2",
			Email:     "test.customer.2@email.com",
			CPF:       "12345678902",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}

	tests := []struct {
		name        string
		input       dto.ListCustomersInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should list products successfully",
			input: dto.ListCustomersInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", 1, 10).
					Return(mockCustomers, int64(2), nil)
			},
			expectError: false,
		},
		{
			name: "should return error when repository fails",
			input: dto.ListCustomersInput{
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "", 1, 10).
					Return(nil, int64(0), assert.AnError)
			},
			expectError: true,
			errorType:   &domain.InternalError{},
		},
		{
			name: "should filter by name",
			input: dto.ListCustomersInput{
				Name:  "Test",
				Page:  1,
				Limit: 10,
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindAll(s.ctx, "Test", 1, 10).
					Return(mockCustomers, int64(2), nil)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			customers, total, err := s.useCase.List(s.ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, customers)
				assert.Equal(t, int64(0), total)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, customers)
				assert.Equal(t, len(mockCustomers), len(customers))
				assert.Equal(t, int64(2), total)
			}
		})
	}
}

func (s *CustomerUsecaseSuiteTest) TestCustomerUseCase_Create() {
	tests := []struct {
		name        string
		input       dto.CreateCustomerInput
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should create customer successfully",
			input: dto.CreateCustomerInput{
				Name:  "Test Customer",
				Email: "test.customer.1@email.com",
				CPF:   "123.456.789-00",
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
			input: dto.CreateCustomerInput{
				Name:  "Test Customer",
				Email: "test.customer.2@email.com",
				CPF:   "123.456.789-01",
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

			customer, err := s.useCase.Create(s.ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, customer)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, tt.input.Name, customer.Name)
				assert.Equal(t, tt.input.Email, customer.Email)
				assert.Equal(t, tt.input.CPF, customer.CPF)
			}
		})
	}
}

func (s *CustomerUsecaseSuiteTest) TestCustomerUseCase_Get() {
	currentTime := time.Now()
	mockCustomer := &entity.Customer{
		ID:        1,
		Name:      "Test Customer",
		Email:     "test.customer@email.com",
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
			name: "should get customer successfully",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(mockCustomer, nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when customer doesn't exist",
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

			customer, err := s.useCase.Get(s.ctx, dto.GetCustomerInput{
				ID: tt.id,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, customer)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, mockCustomer.ID, customer.ID)
				assert.Equal(t, mockCustomer.Name, customer.Name)
				assert.Equal(t, mockCustomer.Email, customer.Email)
				assert.Equal(t, mockCustomer.CreatedAt, customer.CreatedAt)
				assert.Equal(t, mockCustomer.UpdatedAt, customer.UpdatedAt)
			}
		})
	}
}

func (s *CustomerUsecaseSuiteTest) TestCustomerUseCase_Update() {
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
				ID:    1,
				Name:  "New Name",
				Email: "new.name@email.com",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(existingCustomer, nil)

				s.mockGateway.EXPECT().
					Update(s.ctx, gomock.Any()).
					DoAndReturn(func(_ context.Context, p *entity.Customer) error {
						assert.Equal(s.T(), "New Name", p.Name)
						assert.Equal(s.T(), "new.name@email.com", p.Email)
						return nil
					})
			},
			expectError: false,
		},
		{
			name: "should return error when customer not found",
			input: dto.UpdateCustomerInput{
				ID:    1,
				Name:  "New Name",
				Email: "new.name@email.com",
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
			input: dto.UpdateCustomerInput{
				ID:    1,
				Name:  "New Name",
				Email: "new.name@email.com",
			},
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(existingCustomer, nil)

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

			customer, err := s.useCase.Update(s.ctx, tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, customer)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, tt.input.Name, customer.Name)
				assert.Equal(t, tt.input.Email, customer.Email)
				assert.Equal(t, existingCustomer.CreatedAt, customer.CreatedAt)
			}
		})
	}
}

func (s *CustomerUsecaseSuiteTest) TestCustomerUseCase_Delete() {
	tests := []struct {
		name        string
		id          uint64
		setupMocks  func()
		expectError bool
		errorType   error
	}{
		{
			name: "should delete customer successfully",
			id:   1,
			setupMocks: func() {
				s.mockGateway.EXPECT().
					FindByID(s.ctx, uint64(1)).
					Return(&entity.Customer{ID: 1}, nil)

				s.mockGateway.EXPECT().
					Delete(s.ctx, uint64(1)).
					Return(nil)
			},
			expectError: false,
		},
		{
			name: "should return not found error when customer doesn't exist",
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
					Return(&entity.Customer{}, nil)

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

			customer, err := s.useCase.Delete(s.ctx, dto.DeleteCustomerInput{ID: tt.id})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, customer)
				if tt.errorType != nil {
					assert.IsType(t, tt.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, customer)
				assert.Equal(t, tt.id, customer.ID)
			}
		})
	}
}
