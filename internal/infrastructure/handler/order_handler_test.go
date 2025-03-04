package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *OrderHandlerSuiteTest) TestOrderHandler_List() {
	tests := []struct {
		name        string
		query       string
		setupMocks  func()
		checkResult func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:  "success",
			query: "/orders",
			setupMocks: func() {
				s.mockController.EXPECT().List(gomock.Any(), gomock.Any(), dto.ListOrdersInput{
					StatusExclude: []valueobject.OrderStatus{valueobject.CANCELLED, valueobject.COMPLETED},
					Page:          1,
					Limit:         10,
					Sort:          "status:d,created_at",
				}).Return([]byte(`{"data":["a":1]}`), nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Contains(t, res.Body.String(), `{"data":["a":1]}`)
			},
		},
		{
			name:  "success - with query",
			query: "/orders?customer_id=1&status=OPEN,PENDING",
			setupMocks: func() {
				s.mockController.EXPECT().List(gomock.Any(), gomock.Any(), dto.ListOrdersInput{
					CustomerID:    1,
					Status:        []valueobject.OrderStatus{valueobject.OPEN, valueobject.PENDING},
					StatusExclude: []valueobject.OrderStatus{valueobject.CANCELLED, valueobject.COMPLETED},
					Page:          1,
					Limit:         10,
					Sort:          "status:d,created_at",
				}).Return([]byte(`{"data":[]}`), nil)
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, res.Code)
				assert.Contains(t, res.Body.String(), `{"data":[]}`)
			},
		},
		{
			name:       "invalid query - customer_id",
			query:      "/orders?customer_id=invalid",
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Contains(t, res.Body.String(), `{"code":400,"message":"invalid parameter"}`)
			},
		},
		{
			name:       "invalid query - status",
			query:      "/orders?status=invalid",
			setupMocks: func() {},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, res.Code)
				assert.Contains(t, res.Body.String(), `{"code":400,"message":"invalid parameter"}`)
			},
		},
		{
			name:  "controller error",
			query: "/orders",
			setupMocks: func() {
				s.mockController.EXPECT().List(gomock.Any(), gomock.Any(), dto.ListOrdersInput{
					StatusExclude: []valueobject.OrderStatus{valueobject.CANCELLED, valueobject.COMPLETED},
					Page:          1,
					Limit:         10,
					Sort:          "status:d,created_at",
				}).Return(nil, domain.NewInternalError(nil))
			},
			checkResult: func(t *testing.T, res *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, res.Code)
				assert.Contains(t, res.Body.String(), `{"code":500,"message":"internal server error"}`)
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			// Arrange
			tt.setupMocks()
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tt.query, nil)

			// Act
			s.router.ServeHTTP(w, req)

			// Assert
			tt.checkResult(t, w)
		})
	}
}
