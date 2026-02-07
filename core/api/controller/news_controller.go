package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type NewsController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Autocompletion(ctx *gin.Context)
}

type newsController struct {
	service service.NewsService
}

func NewNews(s service.NewsService) NewsController {
	return &newsController{s}
}

func (c *newsController) Create(ctx *gin.Context) {
	var req dto.CreateNewsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create news", err).Send(ctx)
		return
	}
	response.NewSuccess("success create news", res).Send(ctx)
}

func (c *newsController) GetAll(ctx *gin.Context) {
	search := ctx.Query("search")
	category := ctx.Query("category")
	title := ctx.Query("title")
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx), search, category, title)
	if err != nil {
		response.NewFailed("failed get news", err).Send(ctx)
		return
	}
	response.NewSuccess("success get news", res, m).Send(ctx)
}

func (c *newsController) Autocompletion(ctx *gin.Context) {
	search := ctx.Query("search")
	res, err := c.service.GetAutocompletion(ctx.Request.Context(), search)
	if err != nil {
		response.NewFailed("failed get autocompletion", err).Send(ctx)
		return
	}
	response.NewSuccess("success get autocompletion", res).Send(ctx)
}

func (c *newsController) GetById(ctx *gin.Context) {
	res, err := c.service.GetById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed get news", err).Send(ctx)
		return
	}
	response.NewSuccess("success get news", res).Send(ctx)
}

func (c *newsController) Update(ctx *gin.Context) {
	var req dto.UpdateNewsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("id"), req)
	if err != nil {
		response.NewFailed("failed update news", err).Send(ctx)
		return
	}
	response.NewSuccess("success update news", res).Send(ctx)
}

func (c *newsController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed delete news", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete news", nil).Send(ctx)
}
