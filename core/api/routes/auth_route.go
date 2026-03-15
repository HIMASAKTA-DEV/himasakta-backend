package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func Auth(app *gin.Engine, authController controller.AuthController, m middleware.Middleware) {
	routes := app.Group("/api/v1/auth")
	{
		routes.POST("/login", authController.Login)

		p := routes.Group("")
		p.Use(m.AuthMiddleware(), m.OnlyAllow("superadmin"))
		{
			p.POST("/update", authController.UpdateAuth)
		}
	}
}
