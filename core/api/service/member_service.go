package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
)

type MemberService interface {
	Create(ctx context.Context, req dto.CreateMemberRequest) (entity.Member, error)
	GetAll(ctx context.Context, metaReq meta.Meta, name string) ([]entity.Member, meta.Meta, error)
	GetById(ctx context.Context, id string) (entity.Member, error)
	Update(ctx context.Context, id string, req dto.UpdateMemberRequest) (entity.Member, error)
	Delete(ctx context.Context, id string) error
}

type memberService struct {
	repo repository.MemberRepository
}

func NewMember(repo repository.MemberRepository) MemberService {
	return &memberService{repo}
}

func (s *memberService) Create(ctx context.Context, req dto.CreateMemberRequest) (entity.Member, error) {
	return s.repo.Create(ctx, nil, entity.Member{
		Nrp:          req.Nrp,
		Name:         req.Name,
		Role:         req.Role,
		DepartmentId: req.DepartmentId,
		PhotoId:      req.PhotoId,
		Period:       req.Period,
	})
}

func (s *memberService) GetAll(ctx context.Context, metaReq meta.Meta, name string) ([]entity.Member, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq, name)
}

func (s *memberService) GetById(ctx context.Context, id string) (entity.Member, error) {
	uid, _ := uuid.Parse(id)
	return s.repo.GetById(ctx, nil, uid)
}

func (s *memberService) Update(ctx context.Context, id string, req dto.UpdateMemberRequest) (entity.Member, error) {
	uid, _ := uuid.Parse(id)
	m, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return m, err
	}

	if req.Nrp != "" {
		m.Nrp = req.Nrp
	}
	if req.Name != "" {
		m.Name = req.Name
	}
	if req.Role != "" {
		m.Role = req.Role
	}
	if req.DepartmentId != nil {
		m.DepartmentId = req.DepartmentId
		m.Department = nil
	}
	if req.PhotoId != nil {
		m.PhotoId = req.PhotoId
		m.Photo = nil
	}
	if req.Period != "" {
		m.Period = req.Period
	}

	return s.repo.Update(ctx, nil, m)
}

func (s *memberService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	m, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, m)
}
