package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/gin-gonic/gin"
)

func Tag(app *gin.Engine, c controller.TagController) {
	r := app.Group("/api/v1/tags")
	{
		r.GET("", c.GetAll)
	}
}
