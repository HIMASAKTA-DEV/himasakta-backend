package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
)

type CabinetInfoService interface {
	Create(ctx context.Context, req dto.CreateCabinetInfoRequest) (entity.CabinetInfo, error)
	GetAll(ctx context.Context, metaReq meta.Meta, period string) ([]entity.CabinetInfo, meta.Meta, error)
	GetById(ctx context.Context, id string) (entity.CabinetInfo, error)
	Update(ctx context.Context, id string, req dto.UpdateCabinetInfoRequest) (entity.CabinetInfo, error)
	Delete(ctx context.Context, id string) error
}

type cabinetInfoService struct {
	repo repository.CabinetInfoRepository
}

func NewCabinetInfo(repo repository.CabinetInfoRepository) CabinetInfoService {
	return &cabinetInfoService{repo}
}

func (s *cabinetInfoService) Create(ctx context.Context, req dto.CreateCabinetInfoRequest) (entity.CabinetInfo, error) {
	return s.repo.Create(ctx, nil, entity.CabinetInfo{
		Visi:     req.Visi,
		Misi:     req.Misi,
		Tagline:  req.Tagline,
		Period:   req.Period,
		LogoId:   req.LogoId,
		IsActive: req.IsActive,
	})
}

func (s *cabinetInfoService) GetAll(ctx context.Context, metaReq meta.Meta, period string) ([]entity.CabinetInfo, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq, period)
}

func (s *cabinetInfoService) GetById(ctx context.Context, id string) (entity.CabinetInfo, error) {
	uid, _ := uuid.Parse(id)
	return s.repo.GetById(ctx, nil, uid)
}

func (s *cabinetInfoService) Update(ctx context.Context, id string, req dto.UpdateCabinetInfoRequest) (entity.CabinetInfo, error) {
	uid, _ := uuid.Parse(id)
	ci, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return ci, err
	}

	if req.Visi != "" {
		ci.Visi = req.Visi
	}
	if req.Misi != "" {
		ci.Misi = req.Misi
	}
	if req.Tagline != "" {
		ci.Tagline = req.Tagline
	}
	if req.Period != "" {
		ci.Period = req.Period
	}
	ci.LogoId = req.LogoId
	if req.IsActive != nil {
		ci.IsActive = *req.IsActive
	}

	return s.repo.Update(ctx, nil, ci)
}

func (s *cabinetInfoService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	ci, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, ci)
}
