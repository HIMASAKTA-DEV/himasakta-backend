package service

import (
	"context"
	"encoding/json"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
)

type GlobalSettingService interface {
	GetWebSettings(ctx context.Context) (dto.WebSettings, error)
	UpdateWebSettings(ctx context.Context, settings dto.WebSettings) error
}

type globalSettingService struct {
	repo repository.GlobalSettingRepository
}

func NewGlobalSetting(repo repository.GlobalSettingRepository) GlobalSettingService {
	return &globalSettingService{repo}
}

func (s *globalSettingService) GetWebSettings(ctx context.Context) (dto.WebSettings, error) {
	var settings dto.WebSettings
	setting, err := s.repo.GetByKey(ctx, "web_settings")
	if err != nil {
		return settings, err
	}

	err = json.Unmarshal([]byte(setting.Value), &settings)
	return settings, err
}

func (s *globalSettingService) UpdateWebSettings(ctx context.Context, settings dto.WebSettings) error {
	val, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	return s.repo.Upsert(ctx, entity.GlobalSetting{
		Key:   "web_settings",
		Value: string(val),
	})
}
