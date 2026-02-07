package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	myjwt "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/jwt"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"
	"github.com/gin-gonic/gin"
)

type (
	AuthController interface {
		Login(ctx *gin.Context)
		Register(ctx *gin.Context)
		VerifyEmail(ctx *gin.Context)
		RefreshToken(ctx *gin.Context)
		Logout(ctx *gin.Context)
		ForgetPassword(ctx *gin.Context)
		ChangePassword(ctx *gin.Context)
		Me(ctx *gin.Context)
	}

	authController struct {
		authService service.AuthService
	}
)

func NewAuth(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.LoginRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	req.UserAgent = ctx.Request.UserAgent()
	req.IP = ctx.ClientIP()

	result, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed login", err).Send(ctx)
		return
	}

	response.NewSuccess("success login", result).Send(ctx)
}

func (c *authController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.RegisterRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	result, err := c.authService.Register(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed register", err).Send(ctx)
		return
	}

	response.NewSuccess("success register", result).Send(ctx)
}

func (c *authController) VerifyEmail(ctx *gin.Context) {
	token := ctx.Query("token")

	if err := c.authService.VerifyEmail(ctx.Request.Context(), token); err != nil {
		response.NewFailed("failed verify email", err).Send(ctx)
		return
	}

	response.NewSuccess("success verify email", nil).Send(ctx)
}

func (c *authController) RefreshToken(ctx *gin.Context) {
	token := ctx.Query("token")
	result, err := c.authService.RefreshToken(ctx.Request.Context(), token)
	if err != nil {
		response.NewFailed("failed refresh token", err).Send(ctx)
		return
	}

	response.NewSuccess("success refresh token", result).Send(ctx)
}

func (c *authController) Logout(ctx *gin.Context) {
	var req dto.LogoutRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.LogoutRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	if err := c.authService.Logout(ctx.Request.Context(), req); err != nil {
		response.NewFailed("failed logout", err).Send(ctx)
		return
	}

	response.NewSuccess("success logout", nil).Send(ctx)
}

func (c *authController) ForgetPassword(ctx *gin.Context) {
	var req dto.ForgetPasswordRequest

	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.ForgetPasswordRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	if err := c.authService.ForgetPassword(ctx, req); err != nil {
		response.NewFailed("failed forget password", err).Send(ctx)
		return
	}

	response.NewSuccess("success forget password", nil).Send(ctx)
}

func (c *authController) ChangePassword(ctx *gin.Context) {
	var req dto.ChangePasswordRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.ChangePasswordRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	token := ctx.Query("token")
	if token == "" {
		response.NewFailed("failed change password", myerror.ErrBodyRequest).Send(ctx)
		return
	}

	claims, err := myjwt.GetPayloadInsideToken(token)
	if err != nil {
		response.NewFailed("failed change password", err).Send(ctx)
		return
	}

	req.Email = claims["email"]
	if err := c.authService.ChangePassword(ctx, req); err != nil {
		response.NewFailed("failed change password", err).Send(ctx)
		return
	}

	response.NewSuccess("success change password", nil).Send(ctx)
}

func (c *authController) Me(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		response.NewFailed("failed get user id", err).Send(ctx)
		return
	}

	res, err := c.authService.GetMe(ctx.Request.Context(), userId)
	if err != nil {
		response.NewFailed("failed get me", err).Send(ctx)
		return
	}

	response.NewSuccess("success get me", res).Send(ctx)
}
