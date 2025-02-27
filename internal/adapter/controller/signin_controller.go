package controller

import (
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"github.com/gin-gonic/gin"
)

type SignInController struct {
	signinUsecase port.SignInUsecasePort
}

func NewSignInController(signinUsecase port.SignInUsecasePort) port.SignInControllerPort {
	return &SignInController{
		signinUsecase: signinUsecase,
	}
}

func (sc *SignInController) Register(router *gin.RouterGroup) {
	router.POST("", sc.SignIn)
}

func (sc *SignInController) GroupRouterPattern() string {
	return "/api/v1/sign-in"
}

// SignIn godoc
//
//	@Summary		Sign in a customer
//	@Description	Sign in a customer
//	@Description	Example CPF: 123.456.789-00
//	@Description	> 2.b: ii. Identificação do Cliente via CPF
//	@Tags			sign-in
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.SignInRequest	true	"SignIn Request"
//	@Success		200		{object}	dto.SignInResponse	"Successfully signed in"
//	@Failure		400		{object}	dto.ErrorResponse	"Validation error"
//	@Failure		401		{object}	dto.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	dto.ErrorResponse	"Internal server error"
//	@Router			/api/v1/sign-in [post]
func (sc *SignInController) SignIn(ctx *gin.Context) {
	var req dto.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	customer, err := sc.signinUsecase.GetByCPF(req.CPF)
	if err != nil {
		switch err {
		case entity.ErrNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, dto.ErrorResponse{Message: "customer not found"})
		case entity.ErrInvalidInput:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, dto.NewSignInResponse(customer))
} 