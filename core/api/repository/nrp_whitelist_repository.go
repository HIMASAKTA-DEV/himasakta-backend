package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NrpWhitelistRepository interface {
	Create(ctx context.Context, tx *gorm.DB, ci entity.NrpWhitelist) (entity.NrpWhitelist, error)
	GetByNrp(ctx context.Context, tx *gorm.DB, ci entity.NrpWhitelist) (entity.NrpWhitelist, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.NrpWhitelist, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.NrpWhitelist, meta.Meta, error)
	Update(ctx context.Context, tx *gorm.DB, ci entity.NrpWhitelist) (entity.NrpWhitelist, error)
	Delete(ctx context.Context, tx *gorm.DB, ci entity.NrpWhitelist) error
}

type nrpWhitelistRepository struct {
	db *gorm.DB
}

func NewNrpWhitelist(db *gorm.DB) NrpWhitelistRepository {
	return &nrpWhitelistRepository{db}
}

func (r *nrpWhitelistRepository) Create(ctx context.Context, tx *gorm.DB, ci entity.NrpWhitelist) (entity.NrpWhitelist, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return ci, fmt.Errorf("database connection is nil")
	}

	if err := tx.WithContext(ctx).Create(&ci).Error; err != nil {
		return ci, err
	}

	return ci, nil
}

func (r *nrpWhitelistRepository) GetByNrp(ctx context.Context, tx *gorm.DB, filter entity.NrpWhitelist) (entity.NrpWhitelist, error) {
	if tx == nil {
		tx = r.db

	}
	if tx == nil {
		return filter, fmt.Errorf("database connection is nil")
	}

	var result entity.NrpWhitelist
	if err := tx.WithContext(ctx).Where("Nrp = ?", filter.Nrp).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (r *nrpWhitelistRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.NrpWhitelist, error) {
	if tx == nil {
		tx = r.db
	}

	if tx == nil {
		return entity.NrpWhitelist{}, fmt.Errorf("database connection is nil")
	}

	var result entity.NrpWhitelist
	if err := tx.WithContext(ctx).Take(&result, "id = ?", id).Error; err != nil {
		return entity.NrpWhitelist{}, err
	}
	return result, nil

}

func (r *nrpWhitelistRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.NrpWhitelist, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}
	var result []entity.NrpWhitelist
	tx = tx.WithContext(ctx).Model(&entity.NrpWhitelist{})

	if err := WithFilters(tx, &metaReq, AddModels(entity.NrpWhitelist{})).Find(&result).Error; err != nil {
		return nil, metaReq, err
	}
	return result, metaReq, nil
}

func (r *nrpWhitelistRepository) Update(ctx context.Context, tx *gorm.DB, ci entity.NrpWhitelist) (entity.NrpWhitelist, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.NrpWhitelist{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&ci).Error; err != nil {
		return entity.NrpWhitelist{}, err
	}
	return ci, nil
}

func (r *nrpWhitelistRepository) Delete(ctx context.Context, tx *gorm.DB, ci entity.NrpWhitelist) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}

	if err := tx.WithContext(ctx).Delete(&ci).Error; err != nil {
		return err
	}
	return nil
}
