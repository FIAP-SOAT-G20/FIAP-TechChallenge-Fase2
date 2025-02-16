package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type CustomerHandler struct {
	controller *controller.CustomerController
}

type ListCustomersQueryRequest struct {
	Name  string `form:"name" example:"John Doe"`
	Page  int    `form:"page,default=1" example:"1"`
	Limit int    `form:"limit,default=10" example:"10"`
}

type CreateCustomerBodyRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100" example:"Produto A"`
	Email string `json:"email" validate:"required,email" example:"test.customer.1@email.com"`
	CPF   string `json:"cpf" validate:"required" example:"123.456.789-00"`
}
type UpdateCustomerUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateCustomerBodyRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100" example:"Produto A"`
	Email string `json:"email" validate:"required,email" example:"test.customer.1@email.com"`
}

type GetCustomerUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type DeleteCustomerUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

func NewCustomerHandler(controller *controller.CustomerController) *CustomerHandler {
	return &CustomerHandler{controller: controller}
}

func (h *CustomerHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListCustomers)
	router.POST("/", h.CreateCustomer)
	router.GET("/:id", h.GetCustomer)
	router.PUT("/:id", h.UpdateCustomer)
	router.DELETE("/:id", h.DeleteCustomer)
}

// ListCustomers godoc
//
//	@Summary		List customers
//	@Description	List all customers
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int										false	"Page number"		default(1)
//	@Param			limit	query		int										false	"Items per page"	default(10)
//	@Param			name	query		string									false	"Filter by name"
//	@Success		200		{object}	presenter.CustomerJsonPaginatedResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorResponse				"Bad Request"
//	@Failure		500		{object}	middleware.ErrorResponse				"Internal Server Error"
//	@Router			/customers [get]
func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	var query ListCustomersQueryRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidQueryParams))
		return
	}

	input := dto.ListCustomersInput{
		Name:   query.Name,
		Page:   query.Page,
		Limit:  query.Limit,
		Writer: c,
	}

	err := h.controller.ListCustomers(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// CreateCustomer godoc
//
//	@Summary		Create customer
//	@Description	Creates a new customer
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			customer	body		CreateCustomerBodyRequest		true	"Customer data"
//	@Success		201			{object}	presenter.CustomerJsonResponse	"Created"
//	@Failure		400			{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		500			{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req CreateCustomerBodyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.CreateCustomerInput{
		Name:   req.Name,
		Email:  req.Email,
		CPF:    req.CPF,
		Writer: c,
	}

	err := h.controller.CreateCustomer(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// GetCustomer godoc
//
//	@Summary		Get customer
//	@Description	Search for a customer by ID
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"Customer ID"
//	@Success		200	{object}	presenter.CustomerJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse		"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/customers/{id} [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	var req GetCustomerUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetCustomerInput{
		ID:     req.ID,
		Writer: c,
	}

	err := h.controller.GetCustomer(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// UpdateCustomer godoc
//
//	@Summary		Update customer
//	@Description	Update an existing customer
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int								true	"Customer ID"
//	@Param			customer	body		UpdateCustomerBodyRequest		true	"Customer data"
//	@Success		200			{object}	presenter.CustomerJsonResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		404			{object}	middleware.ErrorResponse		"Not Found"
//	@Failure		500			{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	var reqUri UpdateCustomerUriRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var req UpdateCustomerBodyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdateCustomerInput{
		ID:     reqUri.ID,
		Name:   req.Name,
		Email:  req.Email,
		Writer: c,
	}

	err := h.controller.UpdateCustomer(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// DeleteCustomer godoc
//
//	@Summary		Delete customer
//	@Description	Deletes a customer by ID
//	@Tags			customers
//	@Produce		json
//	@Param			id	path		int	true	"Customer ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	var req DeleteCustomerUriRequest
	if err := c.ShouldBindUri(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteCustomerInput{
		ID:     req.ID,
		Writer: c,
	}

	if err := h.controller.DeleteCustomer(c.Request.Context(), input); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
