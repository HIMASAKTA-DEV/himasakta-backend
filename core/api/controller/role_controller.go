package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type RoleController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type roleController struct {
	service service.RoleService
}

func NewRole(s service.RoleService) RoleController {
	return &roleController{s}
}

func (c *roleController) Create(ctx *gin.Context) {
	var req dto.CreateRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create role", err).Send(ctx)
		return
	}
	response.NewSuccessCreated("success create role", res).Send(ctx)
}

func (c *roleController) GetAll(ctx *gin.Context) {
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get roles", err).Send(ctx)
		return
	}
	response.NewSuccess("success get roles", res, m).Send(ctx)
}

func (c *roleController) GetById(ctx *gin.Context) {
	res, err := c.service.GetById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed get role", err).Send(ctx)
		return
	}
	response.NewSuccess("success get role", res).Send(ctx)
}

func (c *roleController) Update(ctx *gin.Context) {
	var req dto.UpdateRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("id"), req)
	if err != nil {
		response.NewFailed("failed update role", err).Send(ctx)
		return
	}
	response.NewSuccess("success update role", res).Send(ctx)
}

func (c *roleController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed delete role", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete role", nil).Send(ctx)
}
