package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type DepartmentController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type departmentController struct {
	service service.DepartmentService
}

func NewDepartment(s service.DepartmentService) DepartmentController {
	return &departmentController{s}
}

func (c *departmentController) Create(ctx *gin.Context) {
	var req dto.CreateDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create department", err).Send(ctx)
		return
	}
	response.NewSuccessCreated("success create department", res).Send(ctx)
}

func (c *departmentController) GetAll(ctx *gin.Context) {
	name := ctx.Query("name")
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx), name)
	if err != nil {
		response.NewFailed("failed get departments", err).Send(ctx)
		return
	}
	response.NewSuccess("success get departments", res, m).Send(ctx)
}

func (c *departmentController) GetById(ctx *gin.Context) {
	res, err := c.service.GetByIdContent(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed get department", err).Send(ctx)
		return
	}
	response.NewSuccess("success get department", res).Send(ctx)
}

func (c *departmentController) Update(ctx *gin.Context) {
	var req dto.UpdateDepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("id"), req)
	if err != nil {
		response.NewFailed("failed update department", err).Send(ctx)
		return
	}
	response.NewSuccess("success update department", res).Send(ctx)
}

func (c *departmentController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed delete department", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete department", nil).Send(ctx)
}
