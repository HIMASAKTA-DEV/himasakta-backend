package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/azkaazkun/be-samarta/internal/entity"
	myerror "github.com/azkaazkun/be-samarta/internal/pkg/error"
	"github.com/azkaazkun/be-samarta/internal/pkg/meta"
	"gorm.io/gorm"
)

type (
	SPTJMRepository interface {
		Create(ctx context.Context, tx *gorm.DB, sptjm entity.SPTJM) (entity.SPTJM, error)
		GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.SPTJM, meta.Meta, error)
		GetById(ctx context.Context, tx *gorm.DB, id string) (entity.SPTJM, error)
		Update(ctx context.Context, tx *gorm.DB, sptjm entity.SPTJM) (entity.SPTJM, error)
		Delete(ctx context.Context, tx *gorm.DB, sptjm entity.SPTJM) error
	}

	sptjmRepository struct {
		db *gorm.DB
	}
)

func NewSPTJM(db *gorm.DB) SPTJMRepository {
	return &sptjmRepository{db}
}

func (r *sptjmRepository) Create(ctx context.Context, tx *gorm.DB, sptjm entity.SPTJM) (entity.SPTJM, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&sptjm).Error; err != nil {
		return entity.SPTJM{}, err
	}

	return sptjm, nil
}

func (r *sptjmRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.SPTJM, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}

	var sptjms []entity.SPTJM

	tx = tx.WithContext(ctx).Model(entity.SPTJM{})

	if err := WithFilters(tx, &metaReq,
		AddModels(entity.SPTJM{})).Find(&sptjms).Error; err != nil {
		return nil, metaReq, err
	}

	return sptjms, metaReq, nil
}

func (r *sptjmRepository) GetById(ctx context.Context, tx *gorm.DB, id string) (entity.SPTJM, error) {
	if tx == nil {
		tx = r.db
	}

	var sptjm entity.SPTJM
	if err := tx.WithContext(ctx).Take(&sptjm, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.SPTJM{}, myerror.New("sptjm not found", http.StatusNotFound)
		}
		return entity.SPTJM{}, err
	}

	return sptjm, nil
}

func (r *sptjmRepository) Update(ctx context.Context, tx *gorm.DB, sptjm entity.SPTJM) (entity.SPTJM, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&sptjm).Error; err != nil {
		return entity.SPTJM{}, err
	}

	return sptjm, nil
}

func (r *sptjmRepository) Delete(ctx context.Context, tx *gorm.DB, sptjm entity.SPTJM) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&sptjm).Error; err != nil {
		return err
	}

	return nil
}
