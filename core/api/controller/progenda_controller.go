package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type ProgendaController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type progendaController struct {
	service service.ProgendaService
}

func NewProgenda(s service.ProgendaService) ProgendaController {
	return &progendaController{s}
}

func (c *progendaController) Create(ctx *gin.Context) {
	var req dto.CreateProgendaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create progenda", err).Send(ctx)
		return
	}
	response.NewSuccessCreated("success create progenda", res).Send(ctx)
}

func (c *progendaController) GetAll(ctx *gin.Context) {
	search := ctx.Query("search")
	deptId := ctx.Query("department_id")
	name := ctx.Query("name")
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx), search, deptId, name)
	if err != nil {
		response.NewFailed("failed get progendas", err).Send(ctx)
		return
	}
	response.NewSuccess("success get progendas", res, m).Send(ctx)
}

func (c *progendaController) GetById(ctx *gin.Context) {
	res, err := c.service.GetById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed get progenda", err).Send(ctx)
		return
	}
	response.NewSuccess("success get progenda", res).Send(ctx)
}

func (c *progendaController) Update(ctx *gin.Context) {
	var req dto.UpdateProgendaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err)
		return
	}
	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("id"), req)
	if err != nil {
		response.NewFailed("failed update progenda", err).Send(ctx)
		return
	}
	response.NewSuccess("success update progenda", res).Send(ctx)
}

func (c *progendaController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed delete progenda", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete progenda", nil).Send(ctx)
}
