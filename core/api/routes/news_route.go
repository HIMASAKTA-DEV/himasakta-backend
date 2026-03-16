package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func News(app *gin.Engine, c controller.NewsController, m middleware.Middleware) {
	r := app.Group("/api/v1/news")
	{
		r.GET("", c.GetAll)
		r.GET("/autocompletion", c.Autocompletion)
		r.GET("/s/:slug", c.GetById)
		r.GET("/tags", c.GetAllTags)

		p := r.Group("")
		p.Use(m.AuthMiddleware(), m.OnlyAllow("superadmin"))
		{
			p.POST("", c.Create)
			p.PUT("/:id", c.Update)
			p.DELETE("/:id", c.Delete)
		}
	}
}
