package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/gin-gonic/gin"
)

func Auth(app *gin.Engine, authController controller.AuthController) {
	routes := app.Group("/api/v1/auth")
	{
		routes.POST("/login", authController.Login)
	}
}
