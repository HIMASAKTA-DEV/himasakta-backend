package controller

import (
	"errors"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CabinetInfoController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetCurrentCabinet(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type cabinetInfoController struct {
	service service.CabinetInfoService
}

func NewCabinetInfo(s service.CabinetInfoService) CabinetInfoController {
	return &cabinetInfoController{s}
}

func (c *cabinetInfoController) Create(ctx *gin.Context) {
	var req dto.CreateCabinetInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create cabinet info", err).Send(ctx)
		return
	}
	response.NewSuccessCreated("success create cabinet info", res).Send(ctx)
}

func (c *cabinetInfoController) GetAll(ctx *gin.Context) {
	res, m, err := c.service.GetAll(ctx.Request.Context(), meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get cabinet infos", err).Send(ctx)
		return
	}
	response.NewSuccess("success get cabinet infos", res, m).Send(ctx)
}

func (c *cabinetInfoController) GetCurrentCabinet(ctx *gin.Context) {
	res, err := c.service.GetCurrentCabinet(ctx.Request.Context())
	if err != nil {
		// Log the error for debugging
		// mylog.Errorf("GetCurrentCabinet error: %v (type: %T)", err, err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NewFailedWithCode(404, "current cabinet info not found", err).Send(ctx)
			return
		}
		response.NewFailed("failed get current cabinet info", err).Send(ctx)
		return
	}
	response.NewSuccess("success get current cabinet info", res).Send(ctx)
}

func (c *cabinetInfoController) GetById(ctx *gin.Context) {
	res, err := c.service.GetById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed get cabinet info", err).Send(ctx)
		return
	}
	response.NewSuccess("success get cabinet info", res).Send(ctx)
}

func (c *cabinetInfoController) Update(ctx *gin.Context) {
	var req dto.UpdateCabinetInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("invalid request body", err).Send(ctx)
		return
	}
	res, err := c.service.Update(ctx.Request.Context(), ctx.Param("id"), req)
	if err != nil {
		response.NewFailed("failed update cabinet info", err).Send(ctx)
		return
	}
	response.NewSuccess("success update cabinet info", res).Send(ctx)
}

func (c *cabinetInfoController) Delete(ctx *gin.Context) {
	err := c.service.Delete(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		response.NewFailed("failed delete cabinet info", err).Send(ctx)
		return
	}
	response.NewSuccess("success delete cabinet info", nil).Send(ctx)
}
