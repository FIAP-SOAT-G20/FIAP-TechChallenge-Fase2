package controller

import (
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/infrastructure/handler/response"
	"github.com/gin-gonic/gin"
)

type SignInController struct {
	signinUsecase port.SignInUsecasePort
}

func NewSignInController(signinUsecase port.SignInUsecasePort) *SignInController {
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
//	@Param			request	body		request.SignInRequest	true	"SignIn Request"
//	@Success		200		{object}	response.SignInResponse	"Successfully signed in"
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		401		{object}	response.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/sign-in [post]
func (sc *SignInController) SignIn(ctx *gin.Context) {
	var req request.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	customer, err := sc.signinUsecase.GetByCPF(req.CPF)
	if err != nil {
		switch err {
		case domain.NewNotFoundError(domain.ErrNotFound):
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{Message: "customer not found"})
		case domain.NewInvalidInputError(domain.ErrInvalidInput):
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{Message: "internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, response.NewSignInResponse(customer))
}
