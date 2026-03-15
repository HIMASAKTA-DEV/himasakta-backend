package repository

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagRepository interface {
	BulkAdd(ctx context.Context, tx *gorm.DB, tag []entity.Tag) ([]entity.Tag, error)
	GetAll(ctx context.Context, tx *gorm.DB, meta meta.Meta, search string) ([]entity.Tag, meta.Meta, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTag(db *gorm.DB) TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) BulkAdd(ctx context.Context, tx *gorm.DB, tags []entity.Tag) ([]entity.Tag, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	if len(tags) == 0 {
		return []entity.Tag{}, nil
	}

	if err := db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&tags).Error; err != nil {
		return []entity.Tag{}, err
	}

	var tagNames []string
	for _, t := range tags {
		tagNames = append(tagNames, t.Name)
	}
	var finalTags []entity.Tag

	if err := db.Where("name IN ?", tagNames).Find(&finalTags).Error; err != nil {
		return []entity.Tag{}, err
	}

	return finalTags, nil
}

func (r *tagRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, search string) ([]entity.Tag, meta.Meta, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	// fmt.Println("repo db:", r.db)
	var tags []entity.Tag
	db = db.WithContext(ctx).Model(&entity.Tag{})

	if search != "" {
		db = db.Where("name ILIKE ?", "%"+search+"%")
	}

	if err := WithFilters(db, &metaReq, AddModels(entity.Tag{})).Find(&tags).Error; err != nil {
		return nil, metaReq, err
	}
	return tags, metaReq, nil
}
