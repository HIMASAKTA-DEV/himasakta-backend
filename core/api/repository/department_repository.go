package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(ctx context.Context, tx *gorm.DB, department entity.Department) (entity.Department, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, name string) ([]entity.Department, meta.Meta, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Department, error)
	GetByName(ctx context.Context, tx *gorm.DB, name string) (entity.Department, error)
	Update(ctx context.Context, tx *gorm.DB, department entity.Department) (entity.Department, error)
	Delete(ctx context.Context, tx *gorm.DB, department entity.Department) error
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartment(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db}
}

func (r *departmentRepository) Create(ctx context.Context, tx *gorm.DB, d entity.Department) (entity.Department, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return d, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Create(&d).Error; err != nil {
		return d, err
	}
	return d, nil
}

func (r *departmentRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, name string) ([]entity.Department, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}
	var departments []entity.Department
	tx = tx.WithContext(ctx).Model(&entity.Department{}).Preload("Logo")

	if name != "" {
		tx = tx.Where("name = ?", name)
	}

	if err := WithFilters(tx, &metaReq, AddModels(entity.Department{})).Find(&departments).Error; err != nil {
		return nil, metaReq, err
	}
	return departments, metaReq, nil
}

func (r *departmentRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Department, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Department{}, fmt.Errorf("database connection is nil")
	}
	var d entity.Department
	if err := tx.WithContext(ctx).Preload("Logo").Take(&d, "id = ?", id).Error; err != nil {
		return entity.Department{}, err
	}
	return d, nil
}

func (r *departmentRepository) GetByName(ctx context.Context, tx *gorm.DB, name string) (entity.Department, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Department{}, fmt.Errorf("database connection is nil")
	}
	var d entity.Department
	if err := tx.WithContext(ctx).Preload("Logo").Take(&d, "name = ?", name).Error; err != nil {
		return entity.Department{}, err
	}
	return d, nil
}

func (r *departmentRepository) Update(ctx context.Context, tx *gorm.DB, d entity.Department) (entity.Department, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Department{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&d).Error; err != nil {
		return entity.Department{}, err
	}
	return d, nil
}

func (r *departmentRepository) Delete(ctx context.Context, tx *gorm.DB, d entity.Department) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}
	return tx.WithContext(ctx).Delete(&d).Error
}
