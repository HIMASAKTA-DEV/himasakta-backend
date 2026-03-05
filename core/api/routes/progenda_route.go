package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func Progenda(app *gin.Engine, c controller.ProgendaController, timelineCtrl controller.TimelineController, m middleware.Middleware) {
	r := app.Group("/api/v1/progenda")
	{
		r.GET("", c.GetAll)
		r.GET("/:id", c.GetById)

		p := r.Group("")
		p.Use(m.AuthMiddleware(), m.OnlyAllow("superadmin"))
		{
			p.POST("", c.Create)
			p.PUT("/:id", c.Update)
			p.DELETE("/:id", c.Delete)

			// Individual timeline CRUD
			p.POST("/:id/timeline", timelineCtrl.Create)
			p.PUT("/timeline/:timelineId", timelineCtrl.Update)
			p.DELETE("/timeline/:timelineId", timelineCtrl.Delete)
		}
	}
}
