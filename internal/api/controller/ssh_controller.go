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
	SSHController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	sshController struct {
		sshService service.SSHService
	}
)

func NewSSH(sshService service.SSHService) SSHController {
	return &sshController{
		sshService: sshService,
	}
}

func (c *sshController) Create(ctx *gin.Context) {
	var req dto.CreateSSHRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.CreateSSHRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	result, err := c.sshService.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create ssh", err).Send(ctx)
		return
	}

	response.NewSuccess("success create ssh", result).Send(ctx)
}

func (c *sshController) GetAll(ctx *gin.Context) {
	result, metaRes, err := c.sshService.GetAll(ctx.Request.Context(), meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get all ssh", err).Send(ctx)
		return
	}

	response.NewSuccess("success get all ssh", result, metaRes).Send(ctx)
}

func (c *sshController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.sshService.GetById(ctx.Request.Context(), id)
	if err != nil {
		response.NewFailed("failed get detail ssh", err).Send(ctx)
		return
	}

	response.NewSuccess("success get detail ssh", result).Send(ctx)
}

func (c *sshController) Update(ctx *gin.Context) {
	var req dto.UpdateSSHRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.UpdateSSHRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	req.ID = ctx.Param("id")
	result, err := c.sshService.Update(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed update ssh", err).Send(ctx)
		return
	}

	response.NewSuccess("success update ssh", result).Send(ctx)
}

func (c *sshController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.sshService.Delete(ctx.Request.Context(), id); err != nil {
		response.NewFailed("failed delete ssh", err).Send(ctx)
		return
	}

	response.NewSuccess("success delete ssh", nil).Send(ctx)
}
