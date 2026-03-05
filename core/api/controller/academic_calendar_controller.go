package controller

import (
	"strconv"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type AcademicCalendarController interface {
	GetCalendar(ctx *gin.Context)
}

type academicCalendarController struct {
	service service.AcademicCalendarService
}

func NewAcademicCalendar(s service.AcademicCalendarService) AcademicCalendarController {
	return &academicCalendarController{service: s}
}

func (c *academicCalendarController) GetCalendar(ctx *gin.Context) {
	month, _ := strconv.Atoi(ctx.Query("month"))
	year, _ := strconv.Atoi(ctx.Query("year"))

	res, err := c.service.GetCalendar(ctx.Request.Context(), month, year)
	if err != nil {
		response.NewFailed("failed to fetch academic calendar", err).Send(ctx)
		return
	}
	response.NewSuccess("success fetch academic calendar", res).Send(ctx)
}
