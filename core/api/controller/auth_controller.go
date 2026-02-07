package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuth(authService service.AuthService) AuthController {
	return &authController{authService: authService}
}

func (c *authController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}

	res, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailedWithCode(401, "login failed", err).Send(ctx)
		return
	}

	response.NewSuccess("login success", res).Send(ctx)
}
