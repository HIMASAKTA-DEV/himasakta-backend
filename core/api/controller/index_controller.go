package controller

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type IndexController interface {
	Index(ctx *gin.Context)
}

type indexController struct{}

func NewIndex() IndexController {
	return &indexController{}
}

func (c *indexController) Index(ctx *gin.Context) {
	htmlPath := filepath.Join("public", "index.html")
	htmlContent, err := os.ReadFile(htmlPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to load documentation: "+err.Error())
		return
	}
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", htmlContent)
}
