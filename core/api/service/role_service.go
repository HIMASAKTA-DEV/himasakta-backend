package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
)

type RoleService interface {
	Create(ctx context.Context, req dto.CreateRoleRequest) (entity.Role, error)
	GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.Role, meta.Meta, error)
	GetById(ctx context.Context, id string) (entity.Role, error)
	Update(ctx context.Context, id string, req dto.UpdateRoleRequest) (entity.Role, error)
	Delete(ctx context.Context, id string) error
}

type roleService struct {
	repo repository.RoleRepository
}

func NewRole(repo repository.RoleRepository) RoleService {
	return &roleService{repo}
}

func (s *roleService) Create(ctx context.Context, req dto.CreateRoleRequest) (entity.Role, error) {
	return s.repo.Create(ctx, nil, entity.Role{
		Rank:        req.Rank,
		Index:       req.Index,
		Description: req.Description,
	})
}

func (s *roleService) GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.Role, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq)
}

func (s *roleService) GetById(ctx context.Context, id string) (entity.Role, error) {
	uid, _ := uuid.Parse(id)
	return s.repo.GetById(ctx, nil, uid)
}

func (s *roleService) Update(ctx context.Context, id string, req dto.UpdateRoleRequest) (entity.Role, error) {
	uid, _ := uuid.Parse(id)
	r, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return r, err
	}

	if req.Rank != "" {
		r.Rank = req.Rank
	}
	if req.Index != 0 {
		r.Index = req.Index
	}
	if req.Description != "" {
		r.Description = req.Description
	}

	return s.repo.Update(ctx, nil, r)
}

func (s *roleService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	r, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, r)
}
