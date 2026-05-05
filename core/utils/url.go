package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetBaseURL(ctx *gin.Context) string {
	scheme := "http"
	if ctx.Request.TLS != nil || ctx.GetHeader("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	return scheme + "://" + ctx.Request.Host
}

func ResolveImageURL(baseURL, url string) string {
	if url == "" || !strings.HasPrefix(url, "/api/static/") {
		return url
	}
	return baseURL + url
}
