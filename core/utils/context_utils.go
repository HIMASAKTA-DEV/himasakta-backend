package utils

import (
	"net/http"

	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/gin-gonic/gin"
)

func GetUserIdFromCtx(ctx *gin.Context) (string, error) {
	user, exists := ctx.Get("user_id")
	if !exists {
		return "", myerror.New("user id not found", http.StatusInternalServerError)
	}

	userId, ok := user.(string)
	if !ok {
		return "", myerror.New("invalid user id", http.StatusInternalServerError)
	}

	return userId, nil
}
