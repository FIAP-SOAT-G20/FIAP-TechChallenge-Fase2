package presenter

import "github.com/gin-gonic/gin"

type ResponseWriter interface {
	JSON(statusCode int, obj any)
	XML(statusCode int, obj any)
	Error(err error) *gin.Error
	// Write(p []byte) (n int, err error)
}

type Error struct {
	Err  error
	Type ErrorType
	Meta any
}

type ErrorType uint64
