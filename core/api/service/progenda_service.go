package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
)

type ProgendaService interface {
	Create(ctx context.Context, req dto.CreateProgendaRequest) (entity.Progenda, error)
	GetAll(ctx context.Context, metaReq meta.Meta, search string, departmentId string, name string) ([]entity.Progenda, meta.Meta, error)
	GetById(ctx context.Context, id string) (entity.Progenda, error)
	Update(ctx context.Context, id string, req dto.UpdateProgendaRequest) (entity.Progenda, error)
	Delete(ctx context.Context, id string) error
}

type progendaService struct {
	repo repository.ProgendaRepository
}

func NewProgenda(repo repository.ProgendaRepository) ProgendaService {
	return &progendaService{repo}
}

func (s *progendaService) Create(ctx context.Context, req dto.CreateProgendaRequest) (entity.Progenda, error) {
	return s.repo.Create(ctx, nil, entity.Progenda{
		Name:         req.Name,
		ThumbnailId:  req.ThumbnailId,
		Goal:         req.Goal,
		Description:  req.Description,
		Timeline:     req.Timeline,
		WebsiteLink:  req.WebsiteLink,
		DepartmentId: req.DepartmentId,
	})
}

func (s *progendaService) GetAll(ctx context.Context, metaReq meta.Meta, search string, departmentId string, name string) ([]entity.Progenda, meta.Meta, error) {
	var deptId *uuid.UUID
	if departmentId != "" {
		id, err := uuid.Parse(departmentId)
		if err == nil {
			deptId = &id
		}
	}
	return s.repo.GetAll(ctx, nil, metaReq, search, deptId, name)
}

func (s *progendaService) GetById(ctx context.Context, id string) (entity.Progenda, error) {
	uid, _ := uuid.Parse(id)
	return s.repo.GetById(ctx, nil, uid)
}

func (s *progendaService) Update(ctx context.Context, id string, req dto.UpdateProgendaRequest) (entity.Progenda, error) {
	uid, _ := uuid.Parse(id)
	p, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return p, err
	}

	if req.Name != "" {
		p.Name = req.Name
	}
	p.ThumbnailId = req.ThumbnailId
	p.Goal = req.Goal
	p.Description = req.Description
	p.Timeline = req.Timeline
	p.WebsiteLink = req.WebsiteLink
	p.DepartmentId = req.DepartmentId

	return s.repo.Update(ctx, nil, p)
}

func (s *progendaService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	p, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, p)
}
