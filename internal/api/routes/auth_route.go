package routes

import (
	"github.com/azkaazkun/be-samarta/internal/api/controller"
	"github.com/azkaazkun/be-samarta/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Auth(app *gin.Engine, authcontroller controller.AuthController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/auth")
	{
		routes.POST("/login", authcontroller.Login)
		routes.POST("/register", authcontroller.Register)
		routes.POST("/forget", authcontroller.ForgetPassword)
		routes.POST("/change", authcontroller.ChangePassword)
		routes.GET("/verify-email", authcontroller.VerifyEmail)
		routes.GET("/refresh-token", authcontroller.RefreshToken)
		routes.GET("/me", middleware.Authenticate(), authcontroller.Me)
	}
}
