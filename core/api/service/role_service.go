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
	GetAll(ctx context.Context, metaReq meta.Meta, name string) ([]entity.Role, meta.Meta, error)
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
		Name:        req.Name,
		Level:       req.Level,
		Description: req.Description,
	})
}

func (s *roleService) GetAll(ctx context.Context, metaReq meta.Meta, name string) ([]entity.Role, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq, name)
}

func (s *roleService) GetById(ctx context.Context, id string) (entity.Role, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.Role{}, err
	}
	return s.repo.GetById(ctx, nil, uid)
}

func (s *roleService) Update(ctx context.Context, id string, req dto.UpdateRoleRequest) (entity.Role, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.Role{}, err
	}
	r, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return r, err
	}

	if req.Name != "" {
		r.Name = req.Name
	}
	if req.Level != 0 {
		r.Level = req.Level
	}
	if req.Description != "" {
		r.Description = req.Description
	}

	return s.repo.Update(ctx, nil, r)
}

func (s *roleService) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	r, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, r)
}
