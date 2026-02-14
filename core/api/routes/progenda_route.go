package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func Progenda(app *gin.Engine, c controller.ProgendaController, tc controller.ProgendaTimelineController, m middleware.Middleware) {
	r := app.Group("/api/v1/progenda")
	{
		r.GET("", c.GetAll)
		r.GET("/:id", c.GetById)

		p := r.Group("")
		p.Use(m.AuthMiddleware())
		{
			p.POST("", c.Create)
			p.PUT("/:id", c.Update)
			p.DELETE("/:id", c.Delete)
		}

		// Nested ProgendaTimeline routes
		t := r.Group("/:progendaId/timeline")
		{
			t.GET("", tc.GetAll)

			tp := t.Group("")
			tp.Use(m.AuthMiddleware())
			{
				tp.POST("", tc.Create)
				tp.PUT("/:timelineId", tc.Update)
				tp.DELETE("/:timelineId", tc.Delete)
			}
		}
	}
}
