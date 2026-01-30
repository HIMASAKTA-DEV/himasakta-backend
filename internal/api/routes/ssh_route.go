package routes

import (
	"github.com/azkaazkun/be-samarta/internal/api/controller"
	"github.com/azkaazkun/be-samarta/internal/entity"
	"github.com/azkaazkun/be-samarta/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SSH(app *gin.Engine, sshController controller.SSHController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/ssh")
	{
		routes.GET("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), sshController.GetAll)
		routes.GET("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), sshController.GetById)
		routes.POST("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), sshController.Create)
		routes.PUT("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), sshController.Update)
		routes.DELETE("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), sshController.Delete)
	}
}
