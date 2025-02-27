package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type CategoryControllerMock struct {
	mock.Mock
}

func NewCategoryControllerMock() *CategoryControllerMock {
	return &CategoryControllerMock{}
}

func (m *CategoryControllerMock) Register(router *gin.RouterGroup) {
	m.Called(router)
}

func (m *CategoryControllerMock) GroupRouterPattern() string {
	args := m.Called()
	return args.String(0)
}

func (m *CategoryControllerMock) CreateCategory(ctx *gin.Context) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": args.Error(0).Error()})
		return
	}
	ctx.JSON(201, args.Get(1))
}

func (m *CategoryControllerMock) GetCategory(ctx *gin.Context) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		ctx.AbortWithStatusJSON(404, gin.H{"error": args.Error(0).Error()})
		return
	}
	ctx.JSON(200, args.Get(1))
}

func (m *CategoryControllerMock) ListCategories(ctx *gin.Context) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": args.Error(0).Error()})
		return
	}
	ctx.JSON(200, args.Get(1))
}

func (m *CategoryControllerMock) UpdateCategory(ctx *gin.Context) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		ctx.AbortWithStatusJSON(404, gin.H{"error": args.Error(0).Error()})
		return
	}
	ctx.JSON(200, args.Get(1))
}

func (m *CategoryControllerMock) DeleteCategory(ctx *gin.Context) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		ctx.AbortWithStatusJSON(404, gin.H{"error": args.Error(0).Error()})
		return
	}
	ctx.Status(204)
} 