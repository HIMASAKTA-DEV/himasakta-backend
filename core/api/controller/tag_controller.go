package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type TagController interface {
	GetAll(ctx *gin.Context)
}

type tagController struct {
	service service.TagService
}

func NewTag(service service.TagService) TagController {
	return &tagController{service}
}

func (c *tagController) GetAll(ctx *gin.Context) {
	search := ctx.Query("search")
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx), search)
	if err != nil {
		response.NewFailed("failed get news", err).Send(ctx)
		return
	}

	response.NewSuccess("success get tags", res, m).Send(ctx)
}
