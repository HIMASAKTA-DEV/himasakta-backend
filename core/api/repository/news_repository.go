package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(ctx context.Context, tx *gorm.DB, n entity.News) (entity.News, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, search string, categories []string, title string) ([]entity.News, meta.Meta, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.News, error)
	GetBySlug(ctx context.Context, tx *gorm.DB, slug string) (entity.News, error)
	Update(ctx context.Context, tx *gorm.DB, n entity.News) (entity.News, error)
	Delete(ctx context.Context, tx *gorm.DB, n entity.News) error
	GetAutocompletion(ctx context.Context, query string) ([]string, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNews(db *gorm.DB) NewsRepository {
	return &newsRepository{db}
}

func (r *newsRepository) Create(ctx context.Context, tx *gorm.DB, n entity.News) (entity.News, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return n, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Create(&n).Error; err != nil {
		return n, err
	}
	return n, nil
}

func (r *newsRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, search string, categories []string, title string) ([]entity.News, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}

	var news []entity.News
	tx = tx.WithContext(ctx).Model(&entity.News{}).Preload("Thumbnail")

	if search != "" {
		tx = tx.Where("title ILIKE ?", "%"+search+"%")
	}

	if metaReq.SortBy == "" {
		tx = tx.Order("created_at DESC")
	}

	if err := WithFilters(tx, &metaReq, AddModels(entity.News{})).Find(&news).Error; err != nil {
		return nil, metaReq, err
	}
	return news, metaReq, nil
}

func (r *newsRepository) GetAutocompletion(ctx context.Context, query string) ([]string, error) {
	var titles []string
	if err := r.db.WithContext(ctx).Model(&entity.News{}).Where("title ILIKE ?", query+"%").Limit(10).Pluck("title", &titles).Error; err != nil {
		return nil, err
	}
	return titles, nil
}

func (r *newsRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.News, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.News{}, fmt.Errorf("database connection is nil")
	}
	var n entity.News
	if err := tx.WithContext(ctx).Preload("Thumbnail").Take(&n, "id = ?", id).Error; err != nil {
		return entity.News{}, err
	}
	return n, nil
}

func (r *newsRepository) GetBySlug(ctx context.Context, tx *gorm.DB, slug string) (entity.News, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.News{}, fmt.Errorf("database connection is nil")
	}
	var n entity.News
	if err := tx.WithContext(ctx).Preload("Thumbnail").Take(&n, "slug = ?", slug).Error; err != nil {
		return entity.News{}, err
	}
	return n, nil
}

func (r *newsRepository) Update(ctx context.Context, tx *gorm.DB, n entity.News) (entity.News, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.News{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&n).Error; err != nil {
		return entity.News{}, err
	}
	return n, nil
}

func (r *newsRepository) Delete(ctx context.Context, tx *gorm.DB, n entity.News) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}
	return tx.WithContext(ctx).Delete(&n).Error
}
