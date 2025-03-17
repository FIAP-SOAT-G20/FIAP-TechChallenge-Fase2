package port

import "github.com/gin-gonic/gin"

type SignInController interface {
	// HTTP Methods
	SignIn(ctx *gin.Context)

	// Router Methods
	Register(router *gin.RouterGroup)
	GroupRouterPattern() string
}
