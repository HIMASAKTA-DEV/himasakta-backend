package repository

import (
	"context"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MemberRepository interface {
	Create(ctx context.Context, tx *gorm.DB, m entity.Member) (entity.Member, error)
	GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, search string) ([]entity.Member, meta.Meta, error)
	GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Member, error)
	Update(ctx context.Context, tx *gorm.DB, m entity.Member) (entity.Member, error)
	Delete(ctx context.Context, tx *gorm.DB, m entity.Member) error
}

type memberRepository struct {
	db *gorm.DB
}

func NewMember(db *gorm.DB) MemberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) Create(ctx context.Context, tx *gorm.DB, m entity.Member) (entity.Member, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return m, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Create(&m).Error; err != nil {
		return m, err
	}
	// Fetch again to get preloads if needed, or just return
	return r.GetById(ctx, tx, m.Id)
}

func (r *memberRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, search string) ([]entity.Member, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return nil, metaReq, fmt.Errorf("database connection is nil")
	}
	var members []entity.Member
	tx = tx.WithContext(ctx).Model(&entity.Member{}).Preload("Department").Preload("Photo").Preload("Role")

	if search != "" {
		tx = tx.Where("members.name ILIKE ?", "%"+search+"%")
	}
	
	tx = tx.Joins("LEFT JOIN roles ON members.role_id = roles.id")

	if metaReq.SortBy == "" {
		if search != "" {
			tx = tx.Order(clause.Expr{SQL: "CASE WHEN members.name ILIKE ? THEN 1 WHEN members.name ILIKE ? THEN 2 ELSE 3 END ASC, roles.level ASC, members.index ASC, members.created_at DESC", Vars: []interface{}{search, search + "%"}})
		} else {
			tx = tx.Order("roles.level ASC, members.index ASC, members.created_at DESC")
		}
	} else {
		tx = tx.Order("roles.level ASC, members.index ASC")
	}

	if err := WithFilters(tx, &metaReq, AddModels(entity.Member{})).Find(&members).Error; err != nil {
		return nil, metaReq, err
	}
	return members, metaReq, nil
}

func (r *memberRepository) GetById(ctx context.Context, tx *gorm.DB, id uuid.UUID) (entity.Member, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Member{}, fmt.Errorf("database connection is nil")
	}
	var m entity.Member
	if err := tx.WithContext(ctx).Preload("Department").Preload("Photo").Preload("Role").Take(&m, "id = ?", id).Error; err != nil {
		return entity.Member{}, err
	}
	return m, nil
}

func (r *memberRepository) Update(ctx context.Context, tx *gorm.DB, m entity.Member) (entity.Member, error) {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return entity.Member{}, fmt.Errorf("database connection is nil")
	}
	if err := tx.WithContext(ctx).Save(&m).Error; err != nil {
		return entity.Member{}, err
	}
	return r.GetById(ctx, tx, m.Id)
}

func (r *memberRepository) Delete(ctx context.Context, tx *gorm.DB, m entity.Member) error {
	if tx == nil {
		tx = r.db
	}
	if tx == nil {
		return fmt.Errorf("database connection is nil")
	}
	return tx.WithContext(ctx).Delete(&m).Error
}
