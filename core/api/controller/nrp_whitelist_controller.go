package controller

import (
	"errors"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type NrpWhitelistController interface {
	CheckWhitelist(ctx *gin.Context)
}

type nrpWhitelistController struct{}

func NewNrpWhitelist() NrpWhitelistController {
	return &nrpWhitelistController{}
}

type CheckNrpRequest struct {
	Nrp string `json:"nrp" binding:"required"`
}

func (c *nrpWhitelistController) CheckWhitelist(ctx *gin.Context) {
	var req CheckNrpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := response.NewFailed("invalid request body", err, nil)
		ctx.JSON(400, res)
		return
	}

	// Hardcoded whitelist for demo purposes
	whitelist := []string{
		"5025211014",
		"5025211015",
		"5025211016",
		"5025211244",
		"5025211155",
		"5025201088",
	}

	for _, n := range whitelist {
		if n == req.Nrp {
			res := response.NewSuccess("NRP is allowed", nil)
			ctx.JSON(200, res)
			return
		}
	}

	// If not found
	res := response.NewFailedWithCode(403, "NRP is not allowed", errors.New("nrp validation failed"))
	ctx.JSON(403, res)
}
