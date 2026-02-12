package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type NrpWhitelistController interface {
	CheckWhitelist(ctx *gin.Context)
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type nrpWhitelistController struct {
	service service.NrpWhitelistService
}

func NewNrpWhitelist(s service.NrpWhitelistService) NrpWhitelistController {
	return &nrpWhitelistController{s}
}

func (c *nrpWhitelistController) CheckWhitelist(ctx *gin.Context) {
	var req dto.CheckNrpWhitelistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err, nil).Send(ctx)
		return
	}

	res, err := c.service.Check(ctx.Request.Context(), req)
	if err != nil {
		// If not found
		response.NewFailedWithCode(403, "NRP is not allowed", err)
	}
	response.NewSuccess("NRP is allowed", res)

}

func (c *nrpWhitelistController) Create(ctx *gin.Context) {
	var req dto.CreateNrpWhitelistRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err, nil).Send(ctx)
		return
	}

	res, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create nrp whitelist", err).Send(ctx)
		return
	}
	response.NewSuccess("success create nrp whitelist", res).Send(ctx)
}

func (c *nrpWhitelistController) GetAll(ctx *gin.Context) {
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get all nrp whitelist", err).Send(ctx)
		return
	}
	response.NewSuccess("success get all nrp whitelist", res, m).Send(ctx)
}

func (c *nrpWhitelistController) Update(ctx *gin.Context) {
	var req dto.UpdateNrpWhitelistRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}

	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("id"), req)
	if err != nil {
		response.NewFailed("failed update nrp whitelist", err).Send(ctx)
		return
	}
	response.NewSuccess("success update nrp whitelist", res).Send(ctx)
}

func (c *nrpWhitelistController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("nrp"))
	if err != nil {
		response.NewFailed("failed delete cabinet info", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete cabinet info", nil).Send(ctx)
}
