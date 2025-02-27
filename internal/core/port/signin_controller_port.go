package port

import "github.com/gin-gonic/gin"

type SignInControllerPort interface {
	// HTTP Methods
	SignIn(ctx *gin.Context)

	// Router Methods
	Register(router *gin.RouterGroup)
	GroupRouterPattern() string
} 