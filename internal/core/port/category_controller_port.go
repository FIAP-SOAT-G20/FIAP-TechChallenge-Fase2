package port

import (
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
	ListCategories(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	Register(router *gin.RouterGroup)
	GroupRouterPattern() string
}
