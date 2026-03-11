package repository

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GlobalSettingRepository interface {
	GetByKey(ctx context.Context, key string) (entity.GlobalSetting, error)
	Upsert(ctx context.Context, setting entity.GlobalSetting) error
}

type globalSettingRepository struct {
	db *gorm.DB
}

func NewGlobalSetting(db *gorm.DB) GlobalSettingRepository {
	return &globalSettingRepository{db}
}

func (r *globalSettingRepository) GetByKey(ctx context.Context, key string) (entity.GlobalSetting, error) {
	var setting entity.GlobalSetting
	err := r.db.WithContext(ctx).First(&setting, "key = ?", key).Error
	return setting, err
}

func (r *globalSettingRepository) Upsert(ctx context.Context, setting entity.GlobalSetting) error {
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"value": setting.Value}),
	}).Create(&setting).Error
}
