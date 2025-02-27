package port

import (
	"github.com/gin-gonic/gin"
)

type CategoryControllerPort interface {
	CreateCategory(ctx *gin.Context)
	GetCategory(ctx *gin.Context)
	ListCategories(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
} 