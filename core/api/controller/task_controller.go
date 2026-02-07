package controller

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	TaskController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	taskController struct {
		taskService service.TaskService
	}
)

func NewTask(taskService service.TaskService) TaskController {
	return &taskController{
		taskService: taskService,
	}
}

func (c *taskController) Create(ctx *gin.Context) {
	var req dto.CreateTaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.CreateTaskRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	result, err := c.taskService.Create(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed create task", err).Send(ctx)
		return
	}

	response.NewSuccess("success create task", result).Send(ctx)
}

func (c *taskController) GetAll(ctx *gin.Context) {
	users, metaRes, err := c.taskService.GetAll(ctx.Request.Context(), meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get all users", err).Send(ctx)
		return
	}

	response.NewSuccess("success get all users", users, metaRes).Send(ctx)
}

func (c *taskController) GetById(ctx *gin.Context) {
	userId := ctx.Param("id")
	result, err := c.taskService.GetById(ctx.Request.Context(), userId)
	if err != nil {
		response.NewFailed("failed get detail user", err).Send(ctx)
		return
	}

	response.NewSuccess("success get detail user", result).Send(ctx)
}

func (c *taskController) Update(ctx *gin.Context) {
	taskId := ctx.Param("id")
	var req dto.UpdateTaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		err = myerror.GetErrBodyRequest(err, dto.UpdateTaskRequest{})
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	result, err := c.taskService.Update(ctx.Request.Context(), taskId, req)
	if err != nil {
		response.NewFailed("failed update task", err).Send(ctx)
		return
	}

	response.NewSuccess("success update task", result).Send(ctx)
}

func (c *taskController) Delete(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := c.taskService.Delete(ctx.Request.Context(), userId)
	if err != nil {
		response.NewFailed("failed delete user", err).Send(ctx)
		return
	}

	response.NewSuccess("success delete user", nil).Send(ctx)
}
