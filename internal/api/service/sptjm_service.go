package service

import (
	"context"

	"github.com/azkaazkun/be-samarta/internal/api/repository"
	"github.com/azkaazkun/be-samarta/internal/dto"
	"github.com/azkaazkun/be-samarta/internal/entity"
	"github.com/azkaazkun/be-samarta/internal/pkg/meta"
	"gorm.io/gorm"
)

type (
	SPTJMService interface {
		Create(ctx context.Context, req dto.CreateSPTJMRequest) (entity.SPTJM, error)
		GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.SPTJM, meta.Meta, error)
		GetById(ctx context.Context, id string) (entity.SPTJM, error)
		Update(ctx context.Context, req dto.UpdateSPTJMRequest) (entity.SPTJM, error)
		Delete(ctx context.Context, id string) error
	}

	sptjmService struct {
		sptjmRepository repository.SPTJMRepository
		db              *gorm.DB
	}
)

func NewSPTJM(sptjmRepository repository.SPTJMRepository, db *gorm.DB) SPTJMService {
	return &sptjmService{
		sptjmRepository: sptjmRepository,
		db:              db,
	}
}

func (s *sptjmService) Create(ctx context.Context, req dto.CreateSPTJMRequest) (entity.SPTJM, error) {
	sptjm := entity.SPTJM{
		SKPDName:  req.SKPDName,
		Address:   req.Address,
		PIC1Name:  req.PIC1Name,
		PIC1Email: req.PIC1Email,
		PIC1Phone: req.PIC1Phone,
		PIC2Name:  req.PIC2Name,
		PIC2Email: req.PIC2Email,
		PIC2Phone: req.PIC2Phone,
		FileURL:   req.FileURL,
	}

	newSPTJM, err := s.sptjmRepository.Create(ctx, nil, sptjm)
	if err != nil {
		return entity.SPTJM{}, err
	}

	return newSPTJM, nil
}

func (s *sptjmService) GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.SPTJM, meta.Meta, error) {
	sptjms, newMeta, err := s.sptjmRepository.GetAll(ctx, nil, metaReq)
	if err != nil {
		return nil, metaReq, err
	}

	return sptjms, newMeta, nil
}

func (s *sptjmService) GetById(ctx context.Context, id string) (entity.SPTJM, error) {
	sptjm, err := s.sptjmRepository.GetById(ctx, nil, id)
	if err != nil {
		return entity.SPTJM{}, err
	}

	return sptjm, nil
}

func (s *sptjmService) Update(ctx context.Context, req dto.UpdateSPTJMRequest) (entity.SPTJM, error) {
	sptjm, err := s.sptjmRepository.GetById(ctx, nil, req.ID)
	if err != nil {
		return entity.SPTJM{}, err
	}

	sptjm.SKPDName = req.SKPDName
	sptjm.Address = req.Address
	sptjm.PIC1Name = req.PIC1Name
	sptjm.PIC1Email = req.PIC1Email
	sptjm.PIC1Phone = req.PIC1Phone
	sptjm.PIC2Name = req.PIC2Name
	sptjm.PIC2Email = req.PIC2Email
	sptjm.PIC2Phone = req.PIC2Phone

	if req.FileURL != "" {
		sptjm.FileURL = req.FileURL
	}

	updatedSPTJM, err := s.sptjmRepository.Update(ctx, nil, sptjm)
	if err != nil {
		return entity.SPTJM{}, err
	}

	return updatedSPTJM, nil
}

func (s *sptjmService) Delete(ctx context.Context, id string) error {
	sptjm, err := s.sptjmRepository.GetById(ctx, nil, id)
	if err != nil {
		return err
	}

	if err := s.sptjmRepository.Delete(ctx, nil, sptjm); err != nil {
		return err
	}

	return nil
}
