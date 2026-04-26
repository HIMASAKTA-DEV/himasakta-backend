package service

import (
	"context"
	"errors"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"
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
	res, err := s.repo.Create(ctx, nil, entity.Department{
		Name:            req.Name,
		Slug:            utils.ToSlug(req.Name),
		Description:     req.Description,
		LogoId:          req.LogoId.ID,
		InstagramLink:   req.InstagramLink,
		YoutubeLink:     req.YoutubeLink,
		TwitterLink:     req.TwitterLink,
		LinkedinLink:    req.LinkedinLink,
		TiktokLink:      req.TiktokLink,
		BankSoalLink:    req.BankSoalLink,
		SilabusLink:     req.SilabusLink,
		BankRefLink:     req.BankRefLink,
		LeaderId:        req.LeaderId.ID,
	})
	return res, myerror.ParseDBError(err, "department")
}

func (s *departmentService) GetAll(ctx context.Context, metaReq meta.Meta, name string) ([]entity.Department, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq, name)
}

func (s *departmentService) GetByIdContent(ctx context.Context, idOrSlug string) (entity.Department, error) {
	var d entity.Department
	var err error

	uid, parseErr := uuid.Parse(idOrSlug)
	if parseErr == nil {
		d, err = s.repo.GetById(ctx, nil, uid)
	} else {
		d, err = s.repo.GetBySlug(ctx, nil, idOrSlug)
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

	if req.Name != nil {
		d.Name = *req.Name
		d.Slug = utils.ToSlug(*req.Name)
	}
	if req.Description != nil {
		d.Description = *req.Description
	}
	if req.LogoId.Valid {
		d.LogoId = req.LogoId.ID
		d.Logo = nil
	}
	if req.InstagramLink != nil {
		d.InstagramLink = *req.InstagramLink
	}
	if req.YoutubeLink != nil {
		d.YoutubeLink = *req.YoutubeLink
	}
	if req.TwitterLink != nil {
		d.TwitterLink = *req.TwitterLink
	}
	if req.LinkedinLink != nil {
		d.LinkedinLink = *req.LinkedinLink
	}
	if req.TiktokLink != nil {
		d.TiktokLink = *req.TiktokLink
	}
	if req.BankSoalLink != nil {
		d.BankSoalLink = *req.BankSoalLink
	}
	if req.SilabusLink != nil {
		d.SilabusLink = *req.SilabusLink
	}
	if req.BankRefLink != nil {
		d.BankRefLink = *req.BankRefLink
	}
	if req.LeaderId.Valid {
		d.LeaderId = req.LeaderId.ID
		d.Leader = nil
	}

	res, err := s.repo.Update(ctx, nil, d)
	return res, myerror.ParseDBError(err, "department")
}

func (s *departmentService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	d, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, d)
}
