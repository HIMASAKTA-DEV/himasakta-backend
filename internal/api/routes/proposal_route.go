package routes

import (
	"github.com/azkaazkun/be-samarta/internal/api/controller"
	"github.com/azkaazkun/be-samarta/internal/entity"
	"github.com/azkaazkun/be-samarta/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Proposal(app *gin.Engine, proposalController controller.ProposalController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/proposal")
	{
		routes.GET("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), proposalController.GetAll)
		routes.GET("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), proposalController.GetById)
		routes.POST("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), proposalController.Create)
		routes.PUT("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), proposalController.Update)
		routes.DELETE("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), proposalController.Delete)
	}
}
