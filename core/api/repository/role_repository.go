package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(ctx context.Context, tx *gorm.DB, r entity.Role) (entity.Role, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, name string) ([]entity.Role, meta.Meta, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Role, error)
	Update(ctx context.Context, tx *gorm.DB, r entity.Role) (entity.Role, error)
	Delete(ctx context.Context, tx *gorm.DB, r entity.Role) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Create(ctx context.Context, tx *gorm.DB, role entity.Role) (entity.Role, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return role, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Create(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (r *roleRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, name string) ([]entity.Role, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}
	var roles []entity.Role
	tx = tx.WithContext(ctx).Model(&entity.Role{})

	if name != "" {
		tx = tx.Where("name ILIKE ?", "%"+name+"%")
	}

	if err := WithFilters(tx, &metaReq, AddModels(entity.Role{})).Find(&roles).Error; err != nil {
		return nil, metaReq, err
	}
	return roles, metaReq, nil
}

func (r *roleRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Role, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Role{}, fmt.Errorf("database connection is nil")
	}
	var role entity.Role
	if err := tx.WithContext(ctx).Take(&role, "id = ?", id).Error; err != nil {
		return entity.Role{}, err
	}
	return role, nil
}

func (r *roleRepository) Update(ctx context.Context, tx *gorm.DB, role entity.Role) (entity.Role, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Role{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&role).Error; err != nil {
		return entity.Role{}, err
	}
	return role, nil
}

func (r *roleRepository) Delete(ctx context.Context, tx *gorm.DB, role entity.Role) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}
	return tx.WithContext(ctx).Delete(&role).Error
}
