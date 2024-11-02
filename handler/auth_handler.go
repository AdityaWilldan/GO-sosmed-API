package handler

import (
	"GoSosmed/dto"
	"GoSosmed/errorhandler"
	"GoSosmed/helper"
	"GoSosmed/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// implementasi dari interface authService
type authHandler struct {
	service service.AuthService
}

// contructor instance dari authHandler
func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

// methode untuk menangani pendaftaran user baru
func (h *authHandler) Register(ctx *gin.Context) {
	var register dto.ResgisterRequest

	if err := ctx.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(ctx, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "register Success fully, please login",
	})
	ctx.JSON(http.StatusCreated, res)

}

// methode untuk menangani proses login user
func (h *authHandler) Login(ctx *gin.Context) {
	var login dto.LoginRequest

	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		errorhandler.HandleError(ctx, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandleError(ctx, err)
		return
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Success Fully Login",
		Data:       result,
	})
	ctx.JSON(http.StatusOK, res)

}
