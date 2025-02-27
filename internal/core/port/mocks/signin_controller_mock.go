package mocks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type SignInControllerMock struct {
	mock.Mock
}

func NewSignInControllerMock() *SignInControllerMock {
	return &SignInControllerMock{}
}

func (m *SignInControllerMock) Register(router *gin.RouterGroup) {
	m.Called(router)
}

func (m *SignInControllerMock) GroupRouterPattern() string {
	args := m.Called()
	return args.String(0)
}

func (m *SignInControllerMock) SignIn(ctx *gin.Context) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": args.Error(0).Error()})
		return
	}
	ctx.JSON(http.StatusOK, args.Get(1))
} 