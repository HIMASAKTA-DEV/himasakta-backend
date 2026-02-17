package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GalleryRepository interface {
	Create(ctx context.Context, tx *gorm.DB, gallery entity.Gallery) (entity.Gallery, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, caption string) ([]entity.Gallery, meta.Meta, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Gallery, error)
	Update(ctx context.Context, tx *gorm.DB, gallery entity.Gallery) (entity.Gallery, error)
	Delete(ctx context.Context, tx *gorm.DB, gallery entity.Gallery) error
}

type galleryRepository struct {
	db *gorm.DB
}

func NewGallery(db *gorm.DB) GalleryRepository {
	return &galleryRepository{db}
}

func (r *galleryRepository) Create(ctx context.Context, tx *gorm.DB, gallery entity.Gallery) (entity.Gallery, error) {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return gallery, fmt.Errorf("database connection is nil")
	}

	if err := tx.WithContext(ctx).Create(&gallery).Error; err != nil {
		return gallery, err
	}

	return gallery, nil
}

func (r *galleryRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, caption string) ([]entity.Gallery, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}
	var galleries []entity.Gallery
	tx = tx.WithContext(ctx).Model(&entity.Gallery{})

	if metaReq.SortBy == "" {
		tx = tx.Order("created_at DESC")
	}

	if err := WithFilters(tx, &metaReq, AddModels(entity.Gallery{})).Find(&galleries).Error; err != nil {
		return nil, metaReq, err
	}

	return galleries, metaReq, nil
}

func (r *galleryRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Gallery, error) {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return entity.Gallery{}, fmt.Errorf("database connection is nil")
	}

	var gallery entity.Gallery
	if err := tx.WithContext(ctx).Take(&gallery, "id = ?", id).Error; err != nil {
		return entity.Gallery{}, err
	}

	return gallery, nil
}

func (r *galleryRepository) Update(ctx context.Context, tx *gorm.DB, gallery entity.Gallery) (entity.Gallery, error) {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return entity.Gallery{}, fmt.Errorf("database connection is nil")
	}

	if err := tx.WithContext(ctx).Save(&gallery).Error; err != nil {
		return entity.Gallery{}, err
	}

	return gallery, nil
}

func (r *galleryRepository) Delete(ctx context.Context, tx *gorm.DB, gallery entity.Gallery) error {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}

	if err := tx.WithContext(ctx).Delete(&gallery).Error; err != nil {
		return err
	}

	return nil
}
