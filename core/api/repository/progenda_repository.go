package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgendaRepository interface {
	Create(ctx context.Context, tx *gorm.DB, p entity.Progenda) (entity.Progenda, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, search string, departmentId *uuid.UUID, name string) ([]entity.Progenda, meta.Meta, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Progenda, error)
	Update(ctx context.Context, tx *gorm.DB, p entity.Progenda) (entity.Progenda, error)
	Delete(ctx context.Context, tx *gorm.DB, p entity.Progenda) error
}

type progendaRepository struct {
	db *gorm.DB
}

func NewProgenda(db *gorm.DB) ProgendaRepository {
	return &progendaRepository{db}
}

func (r *progendaRepository) Create(ctx context.Context, tx *gorm.DB, p entity.Progenda) (entity.Progenda, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return p, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (r *progendaRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, search string, departmentId *uuid.UUID, name string) ([]entity.Progenda, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}
	var progendas []entity.Progenda
	tx = tx.WithContext(ctx).Model(&entity.Progenda{}).Preload("Department").Preload("Thumbnail").Preload("Feeds").Preload("Timelines")

	if name != "" {
		tx = tx.Where("name = ?", name)
	}

	if search != "" {
		tx = tx.Where("name ILIKE ?", "%"+search+"%")
	}

	if metaReq.SortBy == "" {
		tx = tx.Order("created_at DESC")
	}

	if err := WithFilters(tx, &metaReq, AddModels(entity.Progenda{})).Find(&progendas).Error; err != nil {
		return nil, metaReq, err
	}
	return progendas, metaReq, nil
}

func (r *progendaRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Progenda, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Progenda{}, fmt.Errorf("database connection is nil")
	}
	var p entity.Progenda
	if err := tx.WithContext(ctx).Preload("Department").Preload("Thumbnail").Preload("Feeds").Preload("Timelines").First(&p, "id = ?", id).Error; err != nil {
		return entity.Progenda{}, err
	}
	return p, nil
}

func (r *progendaRepository) Update(ctx context.Context, tx *gorm.DB, p entity.Progenda) (entity.Progenda, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Progenda{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&p).Error; err != nil {
		return entity.Progenda{}, err
	}
	return p, nil
}

func (r *progendaRepository) Delete(ctx context.Context, tx *gorm.DB, p entity.Progenda) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}
	return tx.WithContext(ctx).Delete(&p).Error
}
