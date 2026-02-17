package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
)

type GalleryService interface {
	Create(ctx context.Context, req dto.CreateGalleryRequest) (entity.Gallery, error)
	GetAll(ctx context.Context, metaReq meta.Meta, caption string) ([]entity.Gallery, meta.Meta, error)
	GetById(ctx context.Context, id string) (entity.Gallery, error)
	Update(ctx context.Context, id string, req dto.UpdateGalleryRequest) (entity.Gallery, error)
	Delete(ctx context.Context, id string) error
}

type galleryService struct {
	galleryRepo repository.GalleryRepository
}

func NewGallery(galleryRepo repository.GalleryRepository) GalleryService {
	return &galleryService{galleryRepo}
}

func (s *galleryService) Create(ctx context.Context, req dto.CreateGalleryRequest) (entity.Gallery, error) {
	return s.galleryRepo.Create(ctx, nil, entity.Gallery{
		ImageUrl:     req.ImageUrl,
		Caption:      req.Caption,
		Category:     req.Category,
		DepartmentId: req.DepartmentId,
		ProgendaId:   req.ProgendaId,
		CabinetId:    req.CabinetId,
	})
}

func (s *galleryService) GetAll(ctx context.Context, metaReq meta.Meta, caption string) ([]entity.Gallery, meta.Meta, error) {
	return s.galleryRepo.GetAll(ctx, nil, metaReq, caption)
}

func (s *galleryService) GetById(ctx context.Context, id string) (entity.Gallery, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.Gallery{}, err
	}
	return s.galleryRepo.GetById(ctx, nil, uid)
}

func (s *galleryService) Update(ctx context.Context, id string, req dto.UpdateGalleryRequest) (entity.Gallery, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.Gallery{}, err
	}

	gallery, err := s.galleryRepo.GetById(ctx, nil, uid)
	if err != nil {
		return entity.Gallery{}, err
	}

	if req.ImageUrl != "" {
		gallery.ImageUrl = req.ImageUrl
	}
	if req.Caption != "" {
		gallery.Caption = req.Caption
	}
	if req.Category != "" {
		gallery.Category = req.Category
	}
	if req.DepartmentId != nil {
		gallery.DepartmentId = req.DepartmentId
	}
	if req.ProgendaId != nil {
		gallery.ProgendaId = req.ProgendaId
	}
	if req.CabinetId != nil {
		gallery.CabinetId = req.CabinetId
	}

	return s.galleryRepo.Update(ctx, nil, gallery)
}

func (s *galleryService) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	gallery, err := s.galleryRepo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}

	return s.galleryRepo.Delete(ctx, nil, gallery)
}
