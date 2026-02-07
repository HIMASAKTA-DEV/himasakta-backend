package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	"github.com/gin-gonic/gin"
)

func Gallery(app *gin.Engine, controller controller.GalleryController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/gallery")
	{
		routes.GET("", controller.GetAll)
		routes.GET("/:id", controller.GetById)

		// Protected routes (Admin only)
		protected := routes.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("", controller.Create)
			protected.PUT("/:id", controller.Update)
			protected.DELETE("/:id", controller.Delete)
		}
	}
}
