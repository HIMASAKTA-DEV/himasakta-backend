package controller

import (
	"net/http"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type AnalyticsController interface {
	HandleVisit(c *gin.Context)
	GetStats(c *gin.Context)
}

type analyticsController struct {
	svc service.AnalyticsService
}

func NewAnalyticsController(svc service.AnalyticsService) AnalyticsController {
	return &analyticsController{svc: svc}
}

func (ctrl *analyticsController) HandleVisit(c *gin.Context) {
	visitorId := c.GetHeader("X-Visitor-Id")
	if visitorId == "" {
		c.JSON(http.StatusBadRequest, response.NewFailed("X-Visitor-Id header is required", nil))
		return
	}

	clientIp := c.ClientIP()
	status, err := ctrl.svc.TrackVisit(c.Request.Context(), visitorId, clientIp)
	if err != nil {
		c.JSON(status, response.NewFailed(err.Error(), nil))
		return
	}

	if status == 429 {
		c.Status(http.StatusTooManyRequests)
		return
	}

	c.Status(http.StatusAccepted)
}

func (ctrl *analyticsController) GetStats(c *gin.Context) {
	graphLimit := c.Query("graphlimit")
	stats, err := ctrl.svc.GetStats(c.Request.Context(), graphLimit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewFailed(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccess("success get analytics stats", stats))
}
