package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func NrpWhitelist(app *gin.Engine, c controller.NrpWhitelistController, m middleware.Middleware) {
	r := app.Group("/api/v1/nrp-whitelist")
	{
		r.GET("", c.GetAll)
		r.POST("", c.CheckWhitelist)

		p := r.Group("")
		p.Use(m.AuthMiddleware())
		{
			p.POST("/add", c.Create)
			p.PUT("/:id", c.Update)
			p.DELETE("/:nrp", c.Delete)
		}
	}

}
