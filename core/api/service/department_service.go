package service

import (
	"context"
	"errors"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	var d entity.Department
	var err error

	uid, parseErr := uuid.Parse(idOrName)
	if parseErr == nil {
		d, err = s.repo.GetById(ctx, nil, uid)
	} else {
		d, err = s.repo.GetByName(ctx, nil, idOrName)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return d, myerror.ErrNotFound
		}
		return d, err
	}
	return d, nil
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
	if req.LogoId != nil {
		d.LogoId = req.LogoId
		d.Logo = nil
	}
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
