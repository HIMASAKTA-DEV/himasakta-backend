package repository

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimelineRepository interface {
	Create(ctx context.Context, tx *gorm.DB, tl entity.Timeline) (entity.Timeline, error)
	BulkCreate(ctx context.Context, tx *gorm.DB, tls []entity.Timeline) error
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Timeline, error)
	GetByProgenda(ctx context.Context, tx *gorm.DB, id uuid.UUID) ([]entity.Timeline, error)
	BulkUpdate(ctx context.Context, tx *gorm.DB, tl []entity.Timeline) ([]entity.Timeline, error)
	Delete(ctx context.Context, tx *gorm.DB, tl entity.Timeline) error
	BulkDelete(ctx context.Context, tx *gorm.DB, tl []entity.Timeline) error
}

type timelineRepository struct {
	db *gorm.DB
}

func NewTimeline(db *gorm.DB) TimelineRepository {
	return &timelineRepository{db}
}

func (r *timelineRepository) Create(ctx context.Context, tx *gorm.DB, tl entity.Timeline) (entity.Timeline, error) {
	db := r.db
	if tx != nil {
		db = tx
	}
	if err := db.WithContext(ctx).Create(&tl).Error; err != nil {
		return entity.Timeline{}, err
	}
	return tl, nil
}

func (r *timelineRepository) BulkCreate(ctx context.Context, tx *gorm.DB, tls []entity.Timeline) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	if err := db.WithContext(ctx).Create(&tls).Error; err != nil {
		return err
	}
	return nil
}

func (r *timelineRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Timeline, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	var tl entity.Timeline
	if err := db.WithContext(ctx).First(&tl, "id = ?", id).Error; err != nil {
		return entity.Timeline{}, err
	}
	return tl, nil
}

func (r *timelineRepository) GetByProgenda(ctx context.Context, tx *gorm.DB, id uuid.UUID) ([]entity.Timeline, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	var tl []entity.Timeline
	if err := db.WithContext(ctx).Where(&tl, "ProgendaId = ?", id).Order("date asc").Error; err != nil {
		return nil, err
	}
	return tl, nil
}

func (r *timelineRepository) BulkUpdate(ctx context.Context, tx *gorm.DB, tl []entity.Timeline) ([]entity.Timeline, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Save(tl).Error; err != nil {
		return []entity.Timeline{}, err
	}
	return tl, nil
}

func (r *timelineRepository) Delete(ctx context.Context, tx *gorm.DB, tl entity.Timeline) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Delete(&tl).Error; err != nil {
		return err
	}
	return nil
}

func (r *timelineRepository) BulkDelete(ctx context.Context, tx *gorm.DB, tl []entity.Timeline) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Delete(&tl).Error; err != nil {
		return err
	}
	return nil
}
