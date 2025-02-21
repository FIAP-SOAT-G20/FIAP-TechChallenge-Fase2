package presenter

type ResponseWriter interface {
	JSON(statusCode int, obj any)
	XML(statusCode int, obj any)
}
