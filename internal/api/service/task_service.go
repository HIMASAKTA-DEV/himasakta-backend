package service

import (
	"context"

	"github.com/Flexoo-Academy/Golang-Template/internal/api/repository"
	"github.com/Flexoo-Academy/Golang-Template/internal/dto"
	"github.com/Flexoo-Academy/Golang-Template/internal/entity"
	"github.com/Flexoo-Academy/Golang-Template/internal/pkg/meta"
)

type (
	TaskService interface {
		Create(ctx context.Context, req dto.CreateTaskRequest) (entity.Task, error)
		GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.Task, meta.Meta, error)
		GetById(ctx context.Context, taskId string) (entity.Task, error)
		Update(ctx context.Context, taskId string, req dto.UpdateTaskRequest) (entity.Task, error)
		Delete(ctx context.Context, taskId string) error
	}

	taskService struct {
		taskRepo repository.TaskRepository
	}
)

func NewTask(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo}
}

func (s *taskService) Create(ctx context.Context, req dto.CreateTaskRequest) (entity.Task, error) {
	taskCreateResult, err := s.taskRepo.Create(ctx, nil, entity.Task{
		PhotoUrl:    req.PhotoUrl,
		Description: req.Description,
		Deadline:    req.Deadline,
		Status:      entity.TaskStatus(req.Status),
	})
	if err != nil {
		return entity.Task{}, err
	}

	return taskCreateResult, nil
}

func (s *taskService) GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.Task, meta.Meta, error) {
	return s.taskRepo.GetAll(ctx, nil, metaReq)
}

func (s *taskService) GetById(ctx context.Context, taskId string) (entity.Task, error) {
	return s.taskRepo.GetById(ctx, nil, taskId)
}

func (s *taskService) Update(ctx context.Context, taskId string, req dto.UpdateTaskRequest) (entity.Task, error) {
	task, err := s.taskRepo.GetById(ctx, nil, taskId)
	if err != nil {
		return entity.Task{}, err
	}

	task.PhotoUrl = req.PhotoUrl
	task.Description = req.Description
	task.Deadline = req.Deadline
	task.Status = entity.TaskStatus(req.Status)

	updateTaskResult, err := s.taskRepo.Update(ctx, nil, task)
	if err != nil {
		return entity.Task{}, err
	}

	return updateTaskResult, nil
}

func (s *taskService) Delete(ctx context.Context, taskId string) error {
	task, err := s.taskRepo.GetById(ctx, nil, taskId)
	if err != nil {
		return err
	}

	return s.taskRepo.Delete(ctx, nil, task)
}
