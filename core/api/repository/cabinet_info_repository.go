package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CabinetInfoRepository interface {
	Create(ctx context.Context, tx *gorm.DB, ci entity.CabinetInfo) (entity.CabinetInfo, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.CabinetInfo, meta.Meta, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.CabinetInfo, error)
	GetCurrentCabinet(ctx context.Context, tx *gorm.DB) (entity.CabinetInfo, error)
	Update(ctx context.Context, tx *gorm.DB, ci entity.CabinetInfo) (entity.CabinetInfo, error)
	Delete(ctx context.Context, tx *gorm.DB, ci entity.CabinetInfo) error
}

type cabinetInfoRepository struct {
	db *gorm.DB
}

func NewCabinetInfo(db *gorm.DB) CabinetInfoRepository {
	return &cabinetInfoRepository{db}
}

func (r *cabinetInfoRepository) Create(ctx context.Context, tx *gorm.DB, ci entity.CabinetInfo) (entity.CabinetInfo, error) {
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

func (r *cabinetInfoRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.CabinetInfo, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}
	var infos []entity.CabinetInfo
	tx = tx.WithContext(ctx).Model(&entity.CabinetInfo{}).Preload("Logo").Preload("Organigram")

	if metaReq.SortBy == "" {
		metaReq.SortBy = "period_start"
		metaReq.Sort = "desc"
	}

	if err := WithFilters(tx, &metaReq, AddModels(entity.CabinetInfo{})).Find(&infos).Error; err != nil {
		return nil, metaReq, err
	}
	return infos, metaReq, nil
}

func (r *cabinetInfoRepository) GetCurrentCabinet(ctx context.Context, tx *gorm.DB) (entity.CabinetInfo, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.CabinetInfo{}, fmt.Errorf("database connection is nil")
	}
	var ci entity.CabinetInfo
	if err := tx.WithContext(ctx).Model(&entity.CabinetInfo{}).Preload("Logo").Preload("Organigram").Where("is_active = ?", true).Order("period_start desc").First(&ci).Error; err != nil {
		return entity.CabinetInfo{}, err
	}
	return ci, nil
}

func (r *cabinetInfoRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.CabinetInfo, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.CabinetInfo{}, fmt.Errorf("database connection is nil")
	}
	var ci entity.CabinetInfo
	if err := tx.WithContext(ctx).Preload("Logo").Preload("Organigram").Take(&ci, "id = ?", id).Error; err != nil {
		return entity.CabinetInfo{}, err
	}
	return ci, nil
}

func (r *cabinetInfoRepository) Update(ctx context.Context, tx *gorm.DB, ci entity.CabinetInfo) (entity.CabinetInfo, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.CabinetInfo{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&ci).Error; err != nil {
		return entity.CabinetInfo{}, err
	}
	return ci, nil
}

func (r *cabinetInfoRepository) Delete(ctx context.Context, tx *gorm.DB, ci entity.CabinetInfo) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}
	return tx.WithContext(ctx).Delete(&ci).Error
}
