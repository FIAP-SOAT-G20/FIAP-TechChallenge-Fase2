package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/presenter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
)

type StaffHandler struct {
	controller *controller.StaffController
}

type StaffRequest struct {
	Name string                `json:"name" validate:"required,min=3,max=100" example:"John Doe"`
	Role valueobject.StaffRole `json:"role" validate:"max=500" example:"COOK"`
}

func (p *StaffRequest) Validate() error {
	return GetValidator().Struct(p)
}

type StaffListRequest struct {
	Name  string                `json:"name" validate:"required,min=3,max=100" example:"John Doe"`
	Role  valueobject.StaffRole `json:"role" example:"COOK"`
	Page  int                   `json:"page" validate:"required,gte=1" example:"1"`
	Limit int                   `json:"limit" validate:"required,gte=1,lte=100" example:"10"`
}

func (p *StaffListRequest) Validate() error {
	return GetValidator().Struct(p)
}

func NewStaffHandler(controller *controller.StaffController) *StaffHandler {
	return &StaffHandler{controller: controller}
}

func (h *StaffHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.List)
	router.POST("/", h.Create)
	router.GET("/:id", h.Get)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
}

// List godoc
//
//	@Summary		List staffs
//	@Description	List all staffs
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string									false	"Filter by name"
//	@Param			role	query		string									false	"Filter by role. Available options: COOK, ATTENDANT, MANAGER"
//	@Param			page	query		int										false	"Page number"		default(1)
//	@Param			limit	query		int										false	"Items per page"	default(10)
//	@Success		200		{object}	presenter.StaffJsonPaginatedResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorJsonResponse			"Bad Request"
//	@Failure		500		{object}	middleware.ErrorJsonResponse			"Internal Server Error"
//	@Router			/staffs [get]
func (h *StaffHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	input := dto.ListStaffsInput{
		Name:  c.Query("name"),
		Role:  valueobject.ToStaffRole(c.Query("role")),
		Page:  page,
		Limit: limit,
	}

	h.controller.Presenter = presenter.NewStaffJsonPresenter(c)
	err := h.controller.List(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Create godoc
//
//	@Summary		Create staff
//	@Description	Creates a new staff
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			staff	body		StaffRequest					true	"Staff data"
//	@Success		201		{object}	presenter.StaffJsonResponse		"Created"
//	@Failure		400		{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		500		{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/staffs [post]
func (h *StaffHandler) Create(c *gin.Context) {
	var req StaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	// validate request
	if err := req.Validate(); err != nil {
		_ = c.Error(domain.NewInvalidInputError(err.Error()))
		return
	}

	input := dto.CreateStaffInput{
		Name: req.Name,
		Role: req.Role,
	}

	err := h.controller.Create(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Get godoc
//
//	@Summary		Get staff
//	@Description	Search for a staff by ID
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"Staff ID"
//	@Success		200	{object}	presenter.StaffJsonResponse		"OK"
//	@Failure		400	{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/staffs/{id} [get]
func (h *StaffHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetStaffInput{
		ID: id,
	}
	h.controller.Presenter = presenter.NewStaffJsonPresenter(c)
	err = h.controller.Get(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Update godoc
//
//	@Summary		Update staff
//	@Description	Update an existing staff
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"Staff ID"
//	@Param			staff	body		StaffRequest					true	"Staff data"
//	@Success		200		{object}	presenter.StaffJsonResponse		"OK"
//	@Failure		400		{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404		{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500		{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/staffs/{id} [put]
func (h *StaffHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	var req StaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidBody))
		return
	}

	if err := req.Validate(); err != nil {
		_ = c.Error(domain.NewInvalidInputError(err.Error()))
		return
	}

	input := dto.UpdateStaffInput{
		ID:   id,
		Name: req.Name,
		Role: req.Role,
	}
	h.controller.Presenter = presenter.NewStaffJsonPresenter(c)
	err = h.controller.Update(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// Delete godoc
//
//	@Summary		Delete staff
//	@Description	Deletes a staff by ID
//	@Tags			staffs
//	@Produce		json
//	@Param			id	path		int	true	"Staff ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	middleware.ErrorJsonResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorJsonResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorJsonResponse	"Internal Server Error"
//	@Router			/staffs/{id} [delete]
func (h *StaffHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteStaffInput{
		ID: id,
	}
	h.controller.Presenter = presenter.NewStaffJsonPresenter(c)
	if err := h.controller.Delete(c.Request.Context(), input); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
