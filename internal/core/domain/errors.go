package domain

var (
	ErrConflict           = "data conflicts with existing data"
	ErrNotFound           = "data not found"
	ErrInvalidParam       = "invalid parameter"
	ErrInvalidQueryParams = "invalid query parameters"
	ErrInvalidBody        = "invalid body"

	ErrTokenDuration = "invalid token duration format"
	ErrTokenCreation = "error creating token"
	ErrExpiredToken  = "access token has expired"
	ErrInvalidToken  = "access token is invalid"

	ErrOrderInvalidStatusTransition = "invalid status transition"
	ErrOrderWithoutProducts         = "order without products"
	ErrOrderMandatoryStaffId        = "order mandatory staff id"
	ErrOrderIdMandatory             = "order id mandatory"
	ErrOrderIsNotOnStatusOpen       = "order is not on status open"

	ErrProductIdMandatory = "product id mandatory"

	ErrPageMustBeGreaterThanZero = "page must be greater than zero"
	ErrLimitMustBeBetween1And100 = "limit must be between 1 and 100"

	ErrInternalError = "internal error"
)

type ValidationError struct {
	Message string
	Err     error
}

func (e *ValidationError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type InternalError struct {
	Message string
	Err     error
}

func (e *InternalError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

type InvalidInputError struct {
	Message string
}

func (e *InvalidInputError) Error() string {
	return e.Message
}

func NewValidationError(err error) *ValidationError {
	return &ValidationError{
		Message: "erro de validação",
		Err:     err,
	}
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
	}
}

func NewInternalError(err error) *InternalError {
	return &InternalError{
		Message: "erro interno",
		Err:     err,
	}
}

func NewInvalidInputError(message string) *InvalidInputError {
	return &InvalidInputError{
		Message: message,
	}
}
