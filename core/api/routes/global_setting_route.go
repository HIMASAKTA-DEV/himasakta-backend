package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func GlobalSetting(app *gin.Engine, c controller.GlobalSettingController, m middleware.Middleware) {
	r := app.Group("/api/v1/settings")
	{
		r.GET("/web", c.GetWebSettings)
		
		p := r.Group("")
		p.Use(m.AuthMiddleware(), m.OnlyAllow("superadmin"))
		{
			p.PUT("/web", c.UpdateWebSettings)
		}
	}
}
