package routes

import (
	"github.com/azkaazkun/be-samarta/internal/api/controller"
	"github.com/azkaazkun/be-samarta/internal/entity"
	"github.com/azkaazkun/be-samarta/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Item(app *gin.Engine, itemController controller.ItemController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/item")
	{
		routes.GET("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), itemController.GetAll)
		routes.GET("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), itemController.GetById)
		routes.POST("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), itemController.Create)
		routes.PUT("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), itemController.Update)
		routes.DELETE("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), itemController.Delete)
	}
}
