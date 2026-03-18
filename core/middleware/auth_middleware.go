package middleware

import (
	"log"
	"net/http"
	"strings"

	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	MESSAGE_FAILED_VERIFY_TOKEN = "failed to verify token"
	MESSAGE_USER_NOT_AUTHORIZED = "user not authorized"
	MESSAGE_API_IS_LOCKED       = "api is now locked"
)

var (
	ErrTokenInvalid    = myerror.New("token invalid", http.StatusUnauthorized)
	ErrTokenNotFound   = myerror.New("token not found", http.StatusUnauthorized)
	ErrTokenExpired    = myerror.New("token expired", http.StatusUnauthorized)
	ErrRoleNotAllowed  = myerror.New("role not allowed", http.StatusForbidden)
	ErrTokenNotAllowed = myerror.New("token not allowed", http.StatusUnauthorized)
)

func (m Middleware) OnlyAllow(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole := ctx.MustGet("role").(string)

		for _, role := range roles {
			if userRole == role {
				ctx.Next()
				return
			}
		}

		log.Println(userRole)
		log.Println(roles)

		res := response.NewFailed(MESSAGE_USER_NOT_AUTHORIZED, ErrRoleNotAllowed)
		res.SendWithAbort(ctx)
	}
}

func (m Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			res := response.NewFailed(MESSAGE_FAILED_VERIFY_TOKEN, ErrTokenNotFound)
			res.SendWithAbort(ctx)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			res := response.NewFailed(MESSAGE_FAILED_VERIFY_TOKEN, ErrTokenInvalid)
			res.SendWithAbort(ctx)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)

		idToken, err := m.jwtService.GetClaims(authHeader)
		if err != nil {
			if err.Error() == "token expired" {
				res := response.NewFailed(MESSAGE_FAILED_VERIFY_TOKEN, ErrTokenExpired)
				res.SendWithAbort(ctx)
				return
			}


			log.Printf("verify token error: %v\n", err)
			res := response.NewFailed(MESSAGE_FAILED_VERIFY_TOKEN, ErrTokenInvalid)
			res.SendWithAbort(ctx)
			return
		}

		ctx.Set("token", authHeader)
		ctx.Set("payload", idToken)
		ctx.Set("user_id", idToken["user_id"])
		ctx.Set("username", idToken["username"])
		ctx.Set("role", idToken["role"])
		ctx.Next()
	}
}

func (m Middleware) OptionalAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.Next()
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			ctx.Next()
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)

		idToken, err := m.jwtService.GetClaims(authHeader)
		if err != nil {
			log.Printf("optional verify token error: %v\n", err)
			ctx.Next()
			return
		}

		ctx.Set("token", authHeader)
		ctx.Set("payload", idToken)
		ctx.Set("user_id", idToken["user_id"])
		ctx.Set("username", idToken["username"])
		ctx.Set("role", idToken["role"])
		ctx.Next()
	}
}
