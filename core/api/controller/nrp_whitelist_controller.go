package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type NrpWhitelistController interface {
	GetWhitelist(ctx *gin.Context)
}

type nrpWhitelistController struct{}

func NewNrpWhitelist() NrpWhitelistController {
	return &nrpWhitelistController{}
}

func (c *nrpWhitelistController) GetWhitelist(ctx *gin.Context) {
	// Hardcoded whitelist for demo purposes
	whitelist := []string{
		"5025211014",
		"5025211015",
		"5025211016",
		"5025211244",
		"5025211155",
		"5025201088",
	}

	res := response.BuildResponse("success get whitelist", whitelist)
	ctx.JSON(200, res)
}
