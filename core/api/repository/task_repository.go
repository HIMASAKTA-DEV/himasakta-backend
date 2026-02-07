package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"gorm.io/gorm"
)

type (
	TaskRepository interface {
		Create(ctx context.Context, tx *gorm.DB, task entity.Task, preloads ...string) (entity.Task, error)
		GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, preloads ...string) ([]entity.Task, meta.Meta, error)
		GetById(ctx context.Context, tx *gorm.DB, taskId string, preloads ...string) (entity.Task, error)
		Update(ctx context.Context, tx *gorm.DB, task entity.Task, preloads ...string) (entity.Task, error)
		Delete(ctx context.Context, tx *gorm.DB, task entity.Task) error
	}

	taskRepository struct {
		db *gorm.DB
	}
)

func NewTask(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(ctx context.Context, tx *gorm.DB, task entity.Task, preloads ...string) (entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return task, fmt.Errorf("database connection is nil")
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	if err := tx.WithContext(ctx).Create(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (r *taskRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, preloads ...string) ([]entity.Task, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	var tasks []entity.Task

	tx = tx.WithContext(ctx).Model(entity.Task{})

	if err := WithFilters(tx, &metaReq,
		AddModels(entity.Task{})).Find(&tasks).Error; err != nil {
		return nil, metaReq, err
	}

	return tasks, metaReq, nil
}

func (r *taskRepository) GetById(ctx context.Context, tx *gorm.DB, taskId string, preloads ...string) (entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return entity.Task{}, fmt.Errorf("database connection is nil")
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	var task entity.Task
	if err := tx.WithContext(ctx).Take(&task, "id = ?", taskId).Error; err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Update(ctx context.Context, tx *gorm.DB, task entity.Task, preloads ...string) (entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return entity.Task{}, fmt.Errorf("database connection is nil")
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	if err := tx.WithContext(ctx).Save(&task).Error; err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Delete(ctx context.Context, tx *gorm.DB, task entity.Task) error {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}

	if err := tx.WithContext(ctx).Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
