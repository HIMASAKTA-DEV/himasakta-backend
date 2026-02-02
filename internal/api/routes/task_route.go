package routes

import (
	"github.com/Flexoo-Academy/Golang-Template/internal/api/controller"
	"github.com/Flexoo-Academy/Golang-Template/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Task(app *gin.Engine, taskcontroller controller.TaskController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/task")
	{
		routes.POST("", taskcontroller.Create)
		routes.GET("", taskcontroller.GetAll)
		routes.GET("/:id", taskcontroller.GetById)
		routes.PUT("/:id", taskcontroller.Update)
		routes.DELETE("/:id", taskcontroller.Delete)
	}
}
