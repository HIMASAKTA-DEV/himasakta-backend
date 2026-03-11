package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func Analytics(app *gin.Engine, c controller.AnalyticsController, m middleware.Middleware) {
	r := app.Group("/api/v1/analytics")
	{
		r.POST("/visit", c.HandleVisit)
		r.GET("", c.GetStats)
	}
}
