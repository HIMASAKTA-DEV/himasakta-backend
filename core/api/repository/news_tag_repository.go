package repository

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsTagRepository interface {
	BulkCreate(ctx context.Context, tx *gorm.DB, nt []entity.NewsTag) error
	DeleteByNews(ctx context.Context, tx *gorm.DB, newsId uuid.UUID) error
}

type newsTagRepository struct {
	db *gorm.DB
}

func NewNewsTag(db *gorm.DB) NewsTagRepository {
	return &newsTagRepository{db}
}

func (r *newsTagRepository) BulkCreate(ctx context.Context, tx *gorm.DB, nt []entity.NewsTag) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Create(&nt).Error; err != nil {
		return err
	}

	return nil
}

func (r *newsTagRepository) DeleteByNews(ctx context.Context, tx *gorm.DB, newsId uuid.UUID) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).
		Where("news_id = ?", newsId).
		Delete(&entity.NewsTag{}).
		Error; err != nil {
		return err
	}

	return nil
}
