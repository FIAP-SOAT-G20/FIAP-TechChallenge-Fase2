package dto

import "github.com/gin-gonic/gin"

type ResponseWriter interface {
	JSON(statusCode int, obj any)
	XML(statusCode int, obj any)
	Error(err error) *gin.Error
}
