package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/request"
)

type CustomerHandler struct {
	controller *controller.CustomerController
}

func NewCustomerHandler(controller *controller.CustomerController) *CustomerHandler {
	return &CustomerHandler{controller: controller}
}

func (h *CustomerHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.List)
	router.POST("/", h.Create)
	router.GET("/:id", h.Get)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
}

// List godoc
//
//	@Summary		List customers
//	@Description	List all customers
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string									false	"Filter by name"
//	@Param			page	query		int										false	"Page number"		default(1)
//	@Param			limit	query		int										false	"Items per page"	default(10)
//	@Success		200		{object}	presenter.CustomerJsonPaginatedResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorJsonResponse			"Bad Request"
//	@Failure		500		{object}	middleware.ErrorJsonResponse			"Internal Server Error"
//	@Router			/customers [get]
func (h *CustomerHandler) List(c *gin.Context) {
	var query request.ListCustomersQueryRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidQueryParams))
		return
	}

	input := dto.ListCustomersInput{
		Name:  query.Name,
		Page:  query.Page,
		Limit: query.Limit,
	}
	h.controller.Presenter = presenter.NewCustomerJsonPresenter(c)
	err := h.controller.List(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Create godoc
//
//	@Summary		Create customer
//	@Description	Creates a new customer
//	@Tags			customers, sign-up
//	@Accept			json
//	@Produce		json
//	@Param			customer	body		request.CreateCustomerBodyRequest	true	"Customer data"
//	@Success		201			{object}	presenter.CustomerJsonResponse		"Created"
//	@Failure		400			{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		500			{object}	middleware.ErrorJsonResponse		"Internal Server Error"
//	@Router			/customers [post]
func (h *CustomerHandler) Create(c *gin.Context) {
	var body request.CreateCustomerBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.CreateCustomerInput{
		Name:  body.Name,
		Email: body.Email,
		CPF:   body.CPF,
	}
	h.controller.Presenter = presenter.NewCustomerJsonPresenter(c)
	err := h.controller.Create(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Get godoc
//
//	@Summary		Get customer
//	@Description	Search for a customer by ID
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"Customer ID"
//	@Success		200	{object}	presenter.CustomerJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/customers/{id} [get]
func (h *CustomerHandler) Get(c *gin.Context) {
	var uri request.GetCustomerUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetCustomerInput{
		ID: uri.ID,
	}
	h.controller.Presenter = presenter.NewCustomerJsonPresenter(c)
	err := h.controller.Get(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Update godoc
//
//	@Summary		Update customer
//	@Description	Update an existing customer
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int									true	"Customer ID"
//	@Param			customer	body		request.UpdateCustomerBodyRequest	true	"Customer data"
//	@Success		200			{object}	presenter.CustomerJsonResponse		"OK"
//	@Failure		400			{object}	middleware.ErrorJsonResponse		"Bad Request"
//	@Failure		404			{object}	middleware.ErrorJsonResponse		"Not Found"
//	@Failure		500			{object}	middleware.ErrorJsonResponse		"Internal Server Error"
//	@Router			/customers/{id} [put]
func (h *CustomerHandler) Update(c *gin.Context) {
	var uri request.UpdateCustomerUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var body request.UpdateCustomerBodyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	input := dto.UpdateCustomerInput{
		ID:    uri.ID,
		Name:  body.Name,
		Email: body.Email,
	}
	h.controller.Presenter = presenter.NewCustomerJsonPresenter(c)
	err := h.controller.Update(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Delete godoc
//
//	@Summary		Delete customer
//	@Description	Deletes a customer by ID
//	@Tags			customers
//	@Produce		json
//	@Param			id	path		int								true	"Customer ID"
//	@Success		200	{object}	presenter.CustomerJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/customers/{id} [delete]
func (h *CustomerHandler) Delete(c *gin.Context) {
	var uri request.DeleteCustomerUriRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteCustomerInput{
		ID: uri.ID,
	}
	h.controller.Presenter = presenter.NewCustomerJsonPresenter(c)
	if err := h.controller.Delete(c.Request.Context(), input); err != nil {
		_ = c.Error(err)
		return
	}
}
