package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type ProgendaTimelineController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type progendaTimelineController struct {
	service service.ProgendaTimelineService
}

func NewProgendaTimeline(s service.ProgendaTimelineService) ProgendaTimelineController {
	return &progendaTimelineController{s}
}

func (c *progendaTimelineController) Create(ctx *gin.Context) {
	var req dto.ProgendaTimelineRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Create(ctx.Request.Context(), ctx.Param("progendaId"), req)
	if err != nil {
		response.NewFailed("failed create timeline", err).Send(ctx)
		return
	}
	response.NewSuccessCreated("success create timeline", res).Send(ctx)
}

func (c *progendaTimelineController) GetAll(ctx *gin.Context) {
	res, err := c.service.GetAll(ctx.Request.Context(), ctx.Param("progendaId"))
	if err != nil {
		response.NewFailed("failed get timelines", err).Send(ctx)
		return
	}
	response.NewSuccess("success get timelines", res).Send(ctx)
}

func (c *progendaTimelineController) Update(ctx *gin.Context) {
	var req dto.ProgendaTimelineRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("timelineId"), req)
	if err != nil {
		response.NewFailed("failed update timeline", err).Send(ctx)
		return
	}
	response.NewSuccess("success update timeline", res).Send(ctx)
}

func (c *progendaTimelineController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("timelineId"))
	if err != nil {
		response.NewFailed("failed delete timeline", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete timeline", nil).Send(ctx)
}
