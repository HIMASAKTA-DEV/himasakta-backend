package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MonthlyEventRepository interface {
	Create(ctx context.Context, tx *gorm.DB, e entity.MonthlyEvent) (entity.MonthlyEvent, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, title string) ([]entity.MonthlyEvent, meta.Meta, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.MonthlyEvent, error)
	GetThisMonth(ctx context.Context, tx *gorm.DB) ([]entity.MonthlyEvent, error)
	Update(ctx context.Context, tx *gorm.DB, e entity.MonthlyEvent) (entity.MonthlyEvent, error)
	Delete(ctx context.Context, tx *gorm.DB, e entity.MonthlyEvent) error
}

type monthlyEventRepository struct {
	db *gorm.DB
}

func NewMonthlyEvent(db *gorm.DB) MonthlyEventRepository {
	return &monthlyEventRepository{db}
}

func (r *monthlyEventRepository) Create(ctx context.Context, tx *gorm.DB, e entity.MonthlyEvent) (entity.MonthlyEvent, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return e, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Create(&e).Error; err != nil {
		return e, err
	}
	return e, nil
}

func (r *monthlyEventRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, title string) ([]entity.MonthlyEvent, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}
	var events []entity.MonthlyEvent
	tx = tx.WithContext(ctx).Model(&entity.MonthlyEvent{}).Preload("Thumbnail")

	if title != "" {
		tx = tx.Where("title = ?", title)
	}

	if err := WithFilters(tx, &metaReq, AddModels(entity.MonthlyEvent{})).Find(&events).Error; err != nil {
		return nil, metaReq, err
	}
	return events, metaReq, nil
}

func (r *monthlyEventRepository) GetThisMonth(ctx context.Context, tx *gorm.DB) ([]entity.MonthlyEvent, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	nextMonth := startOfMonth.AddDate(0, 1, 0)

	var events []entity.MonthlyEvent
	if err := tx.WithContext(ctx).Preload("Thumbnail").Where("month >= ? AND month < ?", startOfMonth, nextMonth).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *monthlyEventRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.MonthlyEvent, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.MonthlyEvent{}, fmt.Errorf("database connection is nil")
	}
	var e entity.MonthlyEvent
	if err := tx.WithContext(ctx).Preload("Thumbnail").Take(&e, "id = ?", id).Error; err != nil {
		return entity.MonthlyEvent{}, err
	}
	return e, nil
}

func (r *monthlyEventRepository) Update(ctx context.Context, tx *gorm.DB, e entity.MonthlyEvent) (entity.MonthlyEvent, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.MonthlyEvent{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&e).Error; err != nil {
		return entity.MonthlyEvent{}, err
	}
	return e, nil
}

func (r *monthlyEventRepository) Delete(ctx context.Context, tx *gorm.DB, e entity.MonthlyEvent) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}
	return tx.WithContext(ctx).Delete(&e).Error
}
