package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupSignInTest() (*gin.Engine, *mocks.SignInUsecaseMock) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockUsecase := new(mocks.SignInUsecaseMock)
	controller := NewSignInController(mockUsecase)
	
	group := r.Group(controller.GroupRouterPattern())
	controller.Register(group)
	
	return r, mockUsecase
}

func TestSignInController_SignIn(t *testing.T) {
	r, mockUsecase := setupSignInTest()

	t.Run("success", func(t *testing.T) {
		req := dto.SignInRequest{CPF: "12345678900"}
		expectedCustomer := &entity.Customer{
			ID:   1,
			Name: "Test Customer",
			CPF:  req.CPF,
		}

		mockUsecase.On("GetByCPF", req.CPF).Return(expectedCustomer, nil)

		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/sign-in", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("invalid request", func(t *testing.T) {
		req := dto.SignInRequest{} // CPF vazio
		
		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/sign-in", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("not found", func(t *testing.T) {
		req := dto.SignInRequest{CPF: "99999999999"}
		mockUsecase.On("GetByCPF", req.CPF).Return(nil, entity.ErrNotFound)

		body, _ := json.Marshal(req)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/sign-in", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, request)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockUsecase.AssertExpectations(t)
	})
} 