package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgendaTimelineRepository interface {
	Create(ctx context.Context, tx *gorm.DB, t entity.ProgendaTimeline) (entity.ProgendaTimeline, error)
	GetAllByProgendaId(ctx context.Context, tx *gorm.DB, progendaId uuid.UUID) ([]entity.ProgendaTimeline, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.ProgendaTimeline, error)
	Update(ctx context.Context, tx *gorm.DB, t entity.ProgendaTimeline) (entity.ProgendaTimeline, error)
	Delete(ctx context.Context, tx *gorm.DB, t entity.ProgendaTimeline) error
}

type progendaTimelineRepository struct {
	db *gorm.DB
}

func NewProgendaTimeline(db *gorm.DB) ProgendaTimelineRepository {
	return &progendaTimelineRepository{db}
}

func (r *progendaTimelineRepository) Create(ctx context.Context, tx *gorm.DB, t entity.ProgendaTimeline) (entity.ProgendaTimeline, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return t, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (r *progendaTimelineRepository) GetAllByProgendaId(ctx context.Context, tx *gorm.DB, progendaId uuid.UUID) ([]entity.ProgendaTimeline, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	var timelines []entity.ProgendaTimeline
	if err := tx.WithContext(ctx).Where("progenda_id = ?", progendaId).Order("date ASC").Find(&timelines).Error; err != nil {
		return nil, err
	}
	return timelines, nil
}

func (r *progendaTimelineRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.ProgendaTimeline, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.ProgendaTimeline{}, fmt.Errorf("database connection is nil")
	}
	var t entity.ProgendaTimeline
	if err := tx.WithContext(ctx).Take(&t, "id = ?", id).Error; err != nil {
		return entity.ProgendaTimeline{}, err
	}
	return t, nil
}

func (r *progendaTimelineRepository) Update(ctx context.Context, tx *gorm.DB, t entity.ProgendaTimeline) (entity.ProgendaTimeline, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.ProgendaTimeline{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		return entity.ProgendaTimeline{}, err
	}
	return t, nil
}

func (r *progendaTimelineRepository) Delete(ctx context.Context, tx *gorm.DB, t entity.ProgendaTimeline) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}
	return tx.WithContext(ctx).Delete(&t).Error
}
