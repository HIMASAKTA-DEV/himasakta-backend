package controller

import (
	"net/http"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type GlobalSettingController interface {
	GetWebSettings(c *gin.Context)
	UpdateWebSettings(c *gin.Context)
}

type globalSettingController struct {
	svc service.GlobalSettingService
}

func NewGlobalSetting(svc service.GlobalSettingService) GlobalSettingController {
	return &globalSettingController{svc}
}

func (ctrl *globalSettingController) GetWebSettings(c *gin.Context) {
	settings, err := ctrl.svc.GetWebSettings(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewFailed("failed get web settings", err))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccess("success get web settings", settings))
}

func (ctrl *globalSettingController) UpdateWebSettings(c *gin.Context) {
	var req dto.WebSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewFailed("invalid request body", err))
		return
	}

	if err := ctrl.svc.UpdateWebSettings(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewFailed("failed update web settings", err))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccess("success update web settings", nil))
}
