package routes

import (
	"github.com/azkaazkun/be-samarta/internal/api/controller"
	"github.com/azkaazkun/be-samarta/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SPTJM(app *gin.Engine, sptjmController controller.SPTJMController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/sptjm")
	{
		routes.GET("", middleware.Authenticate(), sptjmController.GetAll)
		routes.GET("/:id", middleware.Authenticate(), sptjmController.GetById)
		routes.POST("", middleware.Authenticate(), sptjmController.Create)
		routes.PUT("/:id", middleware.Authenticate(), sptjmController.Update)
		routes.DELETE("/:id", middleware.Authenticate(), sptjmController.Delete)
	}
}
