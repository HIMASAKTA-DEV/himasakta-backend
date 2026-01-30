package controller

import (
	"github.com/azkaazkun/be-samarta/internal/api/service"
	"github.com/azkaazkun/be-samarta/internal/dto"
	myerror "github.com/azkaazkun/be-samarta/internal/pkg/error"
	"github.com/azkaazkun/be-samarta/internal/pkg/meta"
	"github.com/azkaazkun/be-samarta/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	ItemController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	itemController struct {
		itemService service.ItemService
	}
)

func NewItem(itemService service.ItemService) ItemController {
	return &itemController{
		itemService: itemService,
	}
}

func (c *itemController) Create(ctx *gin.Context) {
	var req dto.CreateItemRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.CreateItemRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	result, err := c.itemService.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create item", err).Send(ctx)
		return
	}

	response.NewSuccess("success create item", result).Send(ctx)
}

func (c *itemController) GetAll(ctx *gin.Context) {
	result, metaRes, err := c.itemService.GetAll(ctx.Request.Context(), meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get all item", err).Send(ctx)
		return
	}

	response.NewSuccess("success get all item", result, metaRes).Send(ctx)
}

func (c *itemController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.itemService.GetById(ctx.Request.Context(), id)
	if err != nil {
		response.NewFailed("failed get detail item", err).Send(ctx)
		return
	}

	response.NewSuccess("success get detail item", result).Send(ctx)
}

func (c *itemController) Update(ctx *gin.Context) {
	var req dto.UpdateItemRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.UpdateItemRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	req.ID = ctx.Param("id")
	result, err := c.itemService.Update(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed update item", err).Send(ctx)
		return
	}

	response.NewSuccess("success update item", result).Send(ctx)
}

func (c *itemController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.itemService.Delete(ctx.Request.Context(), id); err != nil {
		response.NewFailed("failed delete item", err).Send(ctx)
		return
	}

	response.NewSuccess("success delete item", nil).Send(ctx)
}
