package middleware

import (
	"net/http"

	myerror "github.com/azkaazkun/be-samarta/internal/pkg/error"
	"github.com/azkaazkun/be-samarta/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m Middleware) OnlyDebug() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secretToken := ctx.GetHeader("secret_token")

		if secretToken != "Mint4AkseSdong!!" {
			response.NewFailed(
				"invalid secret token you are intruder",
				myerror.New("no no no no yohan", http.StatusUnauthorized),
			).Send(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
