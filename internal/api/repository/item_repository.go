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
	ItemRepository interface {
		Create(ctx context.Context, tx *gorm.DB, item entity.Item) (entity.Item, error)
		GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.Item, meta.Meta, error)
		GetById(ctx context.Context, tx *gorm.DB, id string) (entity.Item, error)
		Update(ctx context.Context, tx *gorm.DB, item entity.Item) (entity.Item, error)
		Delete(ctx context.Context, tx *gorm.DB, item entity.Item) error
		GetByAccountCode(ctx context.Context, tx *gorm.DB, code string) (entity.Item, error)
	}

	itemRepository struct {
		db *gorm.DB
	}
)

func NewItem(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) Create(ctx context.Context, tx *gorm.DB, item entity.Item) (entity.Item, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&item).Error; err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func (r *itemRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.Item, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}

	var itemList []entity.Item

	tx = tx.WithContext(ctx).Model(entity.Item{})

	if err := WithFilters(tx, &metaReq,
		AddModels(entity.Item{})).Find(&itemList).Error; err != nil {
		return nil, metaReq, err
	}

	return itemList, metaReq, nil
}

func (r *itemRepository) GetById(ctx context.Context, tx *gorm.DB, id string) (entity.Item, error) {
	if tx == nil {
		tx = r.db
	}

	var item entity.Item
	if err := tx.WithContext(ctx).Take(&item, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Item{}, myerror.New("item not found", http.StatusNotFound)
		}
		return entity.Item{}, err
	}

	return item, nil
}

func (r *itemRepository) Update(ctx context.Context, tx *gorm.DB, item entity.Item) (entity.Item, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&item).Error; err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func (r *itemRepository) Delete(ctx context.Context, tx *gorm.DB, item entity.Item) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&item).Error; err != nil {
		return err
	}

	return nil
}

func (r *itemRepository) GetByAccountCode(ctx context.Context, tx *gorm.DB, code string) (entity.Item, error) {
	if tx == nil {
		tx = r.db
	}

	var item entity.Item
	if err := tx.WithContext(ctx).Take(&item, "account_code = ?", code).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Item{}, myerror.New("item not found", http.StatusNotFound)
		}
		return entity.Item{}, err
	}

	return item, nil
}
