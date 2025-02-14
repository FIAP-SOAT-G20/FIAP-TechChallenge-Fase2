package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type CustomerHandler struct {
	controller *controller.CustomerController
}

type CustomerRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100" example:"Produto A"`
	Email string `json:"email" validate:"required,email" example:""test.customer.1@email.com"`
	CPF   string `json:"cpf" validate:"required" example:"123.456.789-00"`
}

func (p *CustomerRequest) Validate() error {
	return GetValidator().Struct(p)
}

type CustomerListRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100" example:"John Doe"`
	Page  int    `json:"page" validate:"required,gte=1" example:"1"`
	Limit int    `json:"limit" validate:"required,gte=1,lte=100" example:"10"`
}

func (p *CustomerListRequest) Validate() error {
	return GetValidator().Struct(p)
}

func NewCustomerHandler(controller *controller.CustomerController) *CustomerHandler {
	return &CustomerHandler{controller: controller}
}

func (h *CustomerHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListCustomers)
	router.POST("/", h.CreateCustomer)
	router.GET("/:id", h.GetCustomer)
	// router.PUT("/:id", h.UpdateCustomer)
	// router.DELETE("/:id", h.DeleteCustomer)
}

// ListCustomers godoc
//
//	@Summary		List customers
//	@Description	List all customers
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int										false	"Page number"		default(1)
//	@Param			limit		query		int										false	"Items per page"	default(10)
//	@Param			name		query		string									false	"Filter by name"
//	@Success		200			{object}	presenter.CustomerJsonPaginatedResponse	"OK"
//	@Failure		400			{object}	middleware.ErrorResponse				"Bad Request"
//	@Failure		500			{object}	middleware.ErrorResponse				"Internal Server Error"
//	@Router			/customers [get]
func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	input := dto.ListCustomersInput{
		Name:   c.Query("name"),
		Page:   page,
		Limit:  limit,
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
//	@Param			customer	body		CustomerRequest					true	"Customer data"
//	@Success		201		{object}	presenter.CustomerJsonResponse	"Created"
//	@Failure		400		{object}	middleware.ErrorResponse		"Bad Request"
//	@Failure		500		{object}	middleware.ErrorResponse		"Internal Server Error"
//	@Router			/customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req CustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	// validate request
	if err := req.Validate(); err != nil {
		_ = c.Error(domain.NewInvalidInputError(err.Error()))
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetCustomerInput{
		ID:     id,
		Writer: c,
	}

	err = h.controller.GetCustomer(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// // UpdateCustomer godoc
// //
// //	@Summary		Update customer
// //	@Description	Update an existing customer
// //	@Tags			customers
// //	@Accept			json
// //	@Produce		json
// //	@Param			id		path		int								true	"Customer ID"
// //	@Param			customer	body		CustomerRequest					true	"Customer data"
// //	@Success		200		{object}	presenter.CustomerJsonResponse	"OK"
// //	@Failure		400		{object}	middleware.ErrorResponse		"Bad Request"
// //	@Failure		404		{object}	middleware.ErrorResponse		"Not Found"
// //	@Failure		500		{object}	middleware.ErrorResponse		"Internal Server Error"
// //	@Router			/customers/{id} [put]
// func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
// 		return
// 	}

// 	var req CustomerRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
// 		return
// 	}

// 	if err := req.Validate(); err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(err.Error()))
// 		return
// 	}

// 	input := dto.UpdateCustomerInput{
// 		ID:          id,
// 		Name:        req.Name,
// 		Description: req.Description,
// 		Price:       req.Price,
// 		CategoryID:  req.CategoryID,
// 		Writer:      c,
// 	}

// 	err = h.controller.UpdateCustomer(c.Request.Context(), input)
// 	if err != nil {
// 		_ = c.Error(err)
// 		return
// 	}
// }

// // DeleteCustomer godoc
// //
// //	@Summary		Delete customer
// //	@Description	Deletes a customer by ID
// //	@Tags			customers
// //	@Produce		json
// //	@Param			id	path		int	true	"Customer ID"
// //	@Success		204	{object}	nil
// //	@Failure		400	{object}	middleware.ErrorResponse	"Bad Request"
// //	@Failure		404	{object}	middleware.ErrorResponse	"Not Found"
// //	@Failure		500	{object}	middleware.ErrorResponse	"Internal Server Error"
// //	@Router			/customers/{id} [delete]
// func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
// 		return
// 	}

// 	input := dto.DeleteCustomerInput{
// 		ID:     id,
// 		Writer: c,
// 	}

// 	if err := h.controller.DeleteCustomer(c.Request.Context(), input); err != nil {
// 		_ = c.Error(err)
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }
