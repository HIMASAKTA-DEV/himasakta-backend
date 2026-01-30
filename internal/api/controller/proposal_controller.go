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
	ProposalController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	proposalController struct {
		proposalService service.ProposalService
	}
)

func NewProposal(proposalService service.ProposalService) ProposalController {
	return &proposalController{
		proposalService: proposalService,
	}
}

func (c *proposalController) Create(ctx *gin.Context) {
	var req dto.CreateProposalRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.CreateProposalRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	result, err := c.proposalService.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create proposal", err).Send(ctx)
		return
	}

	response.NewSuccess("success create proposal", result).Send(ctx)
}

func (c *proposalController) GetAll(ctx *gin.Context) {
	result, metaRes, err := c.proposalService.GetAll(ctx.Request.Context(), meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get all proposal", err).Send(ctx)
		return
	}

	response.NewSuccess("success get all proposal", result, metaRes).Send(ctx)
}

func (c *proposalController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.proposalService.GetById(ctx.Request.Context(), id)
	if err != nil {
		response.NewFailed("failed get detail proposal", err).Send(ctx)
		return
	}

	response.NewSuccess("success get detail proposal", result).Send(ctx)
}

func (c *proposalController) Update(ctx *gin.Context) {
	var req dto.UpdateProposalRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.UpdateProposalRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	req.ID = ctx.Param("id")
	result, err := c.proposalService.Update(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed update proposal", err).Send(ctx)
		return
	}

	response.NewSuccess("success update proposal", result).Send(ctx)
}

func (c *proposalController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.proposalService.Delete(ctx.Request.Context(), id); err != nil {
		response.NewFailed("failed delete proposal", err).Send(ctx)
		return
	}

	response.NewSuccess("success delete proposal", nil).Send(ctx)
}
