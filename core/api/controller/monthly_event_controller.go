package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type MonthlyEventController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetThisMonth(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type monthlyEventController struct {
	service service.MonthlyEventService
}

func NewMonthlyEvent(s service.MonthlyEventService) MonthlyEventController {
	return &monthlyEventController{s}
}

func (c *monthlyEventController) Create(ctx *gin.Context) {
	var req dto.CreateMonthlyEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create monthly event", err).Send(ctx)
		return
	}
	response.NewSuccess("success create monthly event", res).Send(ctx)
}

func (c *monthlyEventController) GetAll(ctx *gin.Context) {
	title := ctx.Query("title")
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx), title)
	if err != nil {
		response.NewFailed("failed get monthly events", err).Send(ctx)
		return
	}
	response.NewSuccess("success get monthly events", res, m).Send(ctx)
}

func (c *monthlyEventController) GetThisMonth(ctx *gin.Context) {
	res, err := c.service.GetThisMonth(ctx.Request.Context())
	if err != nil {
		response.NewFailed("failed get this month events", err).Send(ctx)
		return
	}
	response.NewSuccess("success get this month events", res).Send(ctx)
}

func (c *monthlyEventController) GetById(ctx *gin.Context) {
	res, err := c.service.GetById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed get monthly event", err).Send(ctx)
		return
	}
	response.NewSuccess("success get monthly event", res).Send(ctx)
}

func (c *monthlyEventController) Update(ctx *gin.Context) {
	var req dto.UpdateMonthlyEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("id"), req)
	if err != nil {
		response.NewFailed("failed update monthly event", err).Send(ctx)
		return
	}
	response.NewSuccess("success update monthly event", res).Send(ctx)
}

func (c *monthlyEventController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed delete monthly event", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete monthly event", nil).Send(ctx)
}
