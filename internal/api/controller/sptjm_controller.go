package controller

import (
	"github.com/azkaazkun/be-samarta/internal/api/service"
	"github.com/azkaazkun/be-samarta/internal/dto"
	myerror "github.com/azkaazkun/be-samarta/internal/pkg/error"
	"github.com/azkaazkun/be-samarta/internal/pkg/meta"
	"github.com/azkaazkun/be-samarta/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	SPTJMController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	sptjmController struct {
		sptjmService service.SPTJMService
	}
)

func NewSPTJM(sptjmService service.SPTJMService) SPTJMController {
	return &sptjmController{
		sptjmService: sptjmService,
	}
}

func (c *sptjmController) Create(ctx *gin.Context) {
	var req dto.CreateSPTJMRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.CreateSPTJMRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	result, err := c.sptjmService.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create sptjm", err).Send(ctx)
		return
	}

	response.NewSuccess("success create sptjm", result).Send(ctx)
}

func (c *sptjmController) GetAll(ctx *gin.Context) {
	result, metaRes, err := c.sptjmService.GetAll(ctx.Request.Context(), meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get all sptjm", err).Send(ctx)
		return
	}

	response.NewSuccess("success get all sptjm", result, metaRes).Send(ctx)
}

func (c *sptjmController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.sptjmService.GetById(ctx.Request.Context(), id)
	if err != nil {
		response.NewFailed("failed get detail sptjm", err).Send(ctx)
		return
	}

	response.NewSuccess("success get detail sptjm", result).Send(ctx)
}

func (c *sptjmController) Update(ctx *gin.Context) {
	var req dto.UpdateSPTJMRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.UpdateSPTJMRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	req.ID = ctx.Param("id")
	result, err := c.sptjmService.Update(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed update sptjm", err).Send(ctx)
		return
	}

	response.NewSuccess("success update sptjm", result).Send(ctx)
}

func (c *sptjmController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.sptjmService.Delete(ctx.Request.Context(), id); err != nil {
		response.NewFailed("failed delete sptjm", err).Send(ctx)
		return
	}

	response.NewSuccess("success delete sptjm", nil).Send(ctx)
}
