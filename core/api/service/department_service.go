package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
)

type DepartmentService interface {
	Create(ctx context.Context, req dto.CreateDepartmentRequest) (entity.Department, error)
	GetAll(ctx context.Context, metaReq meta.Meta, name string) ([]entity.Department, meta.Meta, error)
	GetByIdContent(ctx context.Context, idOrName string) (entity.Department, error)
	Update(ctx context.Context, id string, req dto.UpdateDepartmentRequest) (entity.Department, error)
	Delete(ctx context.Context, id string) error
}

type departmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartment(repo repository.DepartmentRepository) DepartmentService {
	return &departmentService{repo}
}

func (s *departmentService) Create(ctx context.Context, req dto.CreateDepartmentRequest) (entity.Department, error) {
	return s.repo.Create(ctx, nil, entity.Department{
		Name:            req.Name,
		Description:     req.Description,
		LogoId:          req.LogoId,
		SocialMediaLink: req.SocialMediaLink,
		BankSoalLink:    req.BankSoalLink,
		SilabusLink:     req.SilabusLink,
		BankRefLink:     req.BankRefLink,
	})
}

func (s *departmentService) GetAll(ctx context.Context, metaReq meta.Meta, name string) ([]entity.Department, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq, name)
}

func (s *departmentService) GetByIdContent(ctx context.Context, idOrName string) (entity.Department, error) {
	uid, err := uuid.Parse(idOrName)
	if err == nil {
		return s.repo.GetById(ctx, nil, uid)
	}
	return s.repo.GetByName(ctx, nil, idOrName)
}

func (s *departmentService) GetById(ctx context.Context, id string) (entity.Department, error) {
	uid, _ := uuid.Parse(id)
	return s.repo.GetById(ctx, nil, uid)
}

func (s *departmentService) Update(ctx context.Context, id string, req dto.UpdateDepartmentRequest) (entity.Department, error) {
	uid, _ := uuid.Parse(id)
	d, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return d, err
	}

	if req.Name != "" {
		d.Name = req.Name
	}
	d.Description = req.Description
	d.LogoId = req.LogoId
	d.SocialMediaLink = req.SocialMediaLink
	d.BankSoalLink = req.BankSoalLink
	d.SilabusLink = req.SilabusLink
	d.BankRefLink = req.BankRefLink

	return s.repo.Update(ctx, nil, d)
}

func (s *departmentService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	d, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, d)
}
