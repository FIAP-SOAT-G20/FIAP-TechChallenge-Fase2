package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/controller"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/adapter/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
)

type StaffHandler struct {
	controller *controller.StaffController
}

type StaffRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100" example:"Nome do funcionario"`
	Role string `json:"role" validate:"max=500" example:"Cargo do funcionario"`
}

func (p *StaffRequest) Validate() error {
	return GetValidator().Struct(p)
}

type StaffListRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100" example:"Funcionario"`
	Role  string `json:"role" example:"COOK"`
	Page  int    `json:"page" validate:"required,gte=1" example:"1"`
	Limit int    `json:"limit" validate:"required,gte=1,lte=100" example:"10"`
}

func (p *StaffListRequest) Validate() error {
	return GetValidator().Struct(p)
}

func NewStaffHandler(controller *controller.StaffController) *StaffHandler {
	return &StaffHandler{controller: controller}
}

func (h *StaffHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListStaffs)
	router.POST("/", h.CreateStaff)
	router.GET("/:id", h.GetStaff)
	router.PUT("/:id", h.UpdateStaff)
	router.DELETE("/:id", h.DeleteStaff)
}

// ListStaffs godoc
//
//	@Summary		List staffs
//	@Description	List all staffs
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int										false	"Page number"		default(1)
//	@Param			limit	query		int										false	"Items per page"	default(10)
//	@Param			name	query		string									false	"Filter by name"
//	@Param			role	query		string									false	"Filter by role. Available options: COOK, ATTENDANT, MANAGER"
//	@Success		200		{object}	presenter.StaffJsonPaginatedResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorResponse				"Bad Request"
//	@Failure		500		{object}	middleware.ErrorResponse				"Internal Server Error"
//	@Router			/staffs [get]
func (h *StaffHandler) ListStaffs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	input := dto.ListStaffsInput{
		Name:   c.Query("name"),
		Role:   c.Query("role"),
		Page:   page,
		Limit:  limit,
		Writer: c,
	}

	err := h.controller.ListStaffs(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// CreateStaff godoc
//
//	@Summary		Create staff
//	@Description	Creates a new staff
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			staff	body		StaffRequest				true	"Staff data"
//	@Success		201		{object}	presenter.StaffJsonResponse	"Created"
//	@Failure		400		{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		500		{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/staffs [post]
func (h *StaffHandler) CreateStaff(c *gin.Context) {
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
		Name:   req.Name,
		Role:   req.Role,
		Writer: c,
	}

	err := h.controller.CreateStaff(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// GetStaff godoc
//
//	@Summary		Get staff
//	@Description	Search for a staff by ID
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int							true	"Staff ID"
//	@Success		200	{object}	presenter.StaffJsonResponse	"OK"
//	@Failure		400	{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/staffs/{id} [get]
func (h *StaffHandler) GetStaff(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.GetStaffInput{
		ID:     id,
		Writer: c,
	}

	err = h.controller.GetStaff(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// UpdateStaff godoc
//
//	@Summary		Update staff
//	@Description	Update an existing staff
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Staff ID"
//	@Param			staff	body		StaffRequest				true	"Staff data"
//	@Success		200		{object}	presenter.StaffJsonResponse	"OK"
//	@Failure		400		{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404		{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500		{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/staffs/{id} [put]
func (h *StaffHandler) UpdateStaff(c *gin.Context) {
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
		ID:     id,
		Name:   req.Name,
		Role:   req.Role,
		Writer: c,
	}

	err = h.controller.UpdateStaff(c.Request.Context(), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
}

// DeleteStaff godoc
//
//	@Summary		Delete staff
//	@Description	Deletes a staff by ID
//	@Tags			staffs
//	@Produce		json
//	@Param			id	path		int	true	"Staff ID"
//	@Success		204	{object}	nil
//	@Failure		400	{object}	middleware.ErrorResponse	"Bad Request"
//	@Failure		404	{object}	middleware.ErrorResponse	"Not Found"
//	@Failure		500	{object}	middleware.ErrorResponse	"Internal Server Error"
//	@Router			/staffs/{id} [delete]
func (h *StaffHandler) DeleteStaff(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(domain.NewInvalidInputError(domain.ErrInvalidParam))
		return
	}

	input := dto.DeleteStaffInput{
		ID:     id,
		Writer: c,
	}

	if err := h.controller.DeleteStaff(c.Request.Context(), input); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
