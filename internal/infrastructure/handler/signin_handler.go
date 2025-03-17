package handler

import (
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/dto"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/port"
	"github.com/gin-gonic/gin"
)

type SignInHandler struct {
	signinUsecase port.SignInUsecasePort
}

func NewSignInHandler(signinUsecase port.SignInUsecasePort) *SignInHandler {
	return &SignInHandler{
		signinUsecase: signinUsecase,
	}
}

func (sh *SignInHandler) Register(router *gin.RouterGroup) {
	router.POST("", sh.SignIn)
}

func (sh *SignInHandler) GroupRouterPattern() string {
	return "/api/v1/sign-in"
}

func (sh *SignInHandler) SignIn(ctx *gin.Context) {
	var req dto.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	customer, err := sh.signinUsecase.GetByCPF(req.CPF)
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
