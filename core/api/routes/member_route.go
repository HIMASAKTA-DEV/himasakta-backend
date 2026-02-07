package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func Member(app *gin.Engine, c controller.MemberController, m middleware.Middleware) {
	r := app.Group("/api/v1/member")
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
	}
}
