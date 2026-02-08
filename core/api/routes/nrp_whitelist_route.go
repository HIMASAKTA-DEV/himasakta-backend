package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/gin-gonic/gin"
)

func NrpWhitelist(route *gin.Engine, controller controller.NrpWhitelistController) {
	routes := route.Group("/api/v1/nrp-whitelist")
	{
		routes.GET("", controller.GetWhitelist)
	}
}
