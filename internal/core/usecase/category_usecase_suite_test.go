package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	mockport "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port/mocks"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/usecase"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type CategoryUsecaseSuiteTest struct {
	suite.Suite
	// conn    *sql.DB
	// mock    sqlmock.Sqlmock
	// handler handler
	mockCategories []*entity.Category
	mockGateway    *mockport.MockCategoryGateway
	useCase        port.CategoryUseCase
	ctx            context.Context
}

func (s *CategoryUsecaseSuiteTest) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.mockGateway = mockport.NewMockCategoryGateway(ctrl)
	s.useCase = usecase.NewCategoryUseCase(s.mockGateway)
	s.ctx = context.Background()
	currentTime := time.Now()
	s.mockCategories = []*entity.Category{
		{
			ID:        1,
			Name:      "Foods",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
		{
			ID:        2,
			Name:      "Beverages",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}
}

// func (ts *CategoryUsecaseSuiteTest) AfterTest(_, _ string) {
// 	// assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
// }

func TestCategoryUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(CategoryUsecaseSuiteTest))
}
