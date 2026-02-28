package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TimelineController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type timelineController struct {
	service service.TimelineService
}

func NewTimeline(s service.TimelineService) TimelineController {
	return &timelineController{s}
}

func (c *timelineController) Create(ctx *gin.Context) {
	progendaId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.NewFailedWithCode(400, "invalid progenda id", err).Send(ctx)
		return
	}

	var req dto.CreateTimelineRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}

	res, err := c.service.Create(ctx.Request.Context(), progendaId, req)
	if err != nil {
		response.NewFailed("failed create timeline", err).Send(ctx)
		return
	}
	response.NewSuccessCreated("success create timeline", res).Send(ctx)
}

func (c *timelineController) Update(ctx *gin.Context) {
	timelineId, err := uuid.Parse(ctx.Param("timelineId"))
	if err != nil {
		response.NewFailedWithCode(400, "invalid timeline id", err).Send(ctx)
		return
	}

	var req dto.UpdateTimelineRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}

	res, err := c.service.Update(ctx.Request.Context(), timelineId, req)
	if err != nil {
		response.NewFailed("failed update timeline", err).Send(ctx)
		return
	}
	response.NewSuccess("success update timeline", res).Send(ctx)
}

func (c *timelineController) Delete(ctx *gin.Context) {
	timelineId, err := uuid.Parse(ctx.Param("timelineId"))
	if err != nil {
		response.NewFailedWithCode(400, "invalid timeline id", err).Send(ctx)
		return
	}

	err = c.service.Delete(ctx.Request.Context(), timelineId)
	if err != nil {
		response.NewFailed("failed delete timeline", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete timeline", nil).Send(ctx)
}
