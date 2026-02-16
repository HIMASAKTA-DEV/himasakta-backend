package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type MemberController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type memberController struct {
	service service.MemberService
}

func NewMember(s service.MemberService) MemberController {
	return &memberController{s}
}

func (c *memberController) Create(ctx *gin.Context) {
	var req dto.CreateMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create member", err).Send(ctx)
		return
	}
	response.NewSuccessCreated("success create member", res).Send(ctx)
}

func (c *memberController) GetAll(ctx *gin.Context) {
	if ctx.Query("groupby") == "rank" {
		res, err := c.service.GetGroupedByRank(ctx.Request.Context(), meta.New(ctx))
		if err != nil {
			response.NewFailed("failed get members grouped by rank", err).Send(ctx)
			return
		}
		response.NewSuccess("success get members grouped by rank", res).Send(ctx)
		return
	}

	name := ctx.Query("name")
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx), name)
	if err != nil {
		response.NewFailed("failed get members", err).Send(ctx)
		return
	}
	response.NewSuccess("success get members", res, m).Send(ctx)
}

func (c *memberController) GetById(ctx *gin.Context) {
	res, err := c.service.GetById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed get member", err).Send(ctx)
		return
	}
	response.NewSuccess("success get member", res).Send(ctx)
}

func (c *memberController) Update(ctx *gin.Context) {
	var req dto.UpdateMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("id"), req)
	if err != nil {
		response.NewFailed("failed update member", err).Send(ctx)
		return
	}
	response.NewSuccess("success update member", res).Send(ctx)
}

func (c *memberController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed delete member", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete member", nil).Send(ctx)
}
