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
	GetAllGrouped(ctx context.Context) ([]dto.MemberGroupResponse, error)
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
		RoleId:       req.RoleId,
		DepartmentId: req.DepartmentId,
		PhotoId:      req.PhotoId,
		CabinetId:    req.CabinetId,
		Index:        req.Index,
	})
}

func (s *memberService) GetAll(ctx context.Context, metaReq meta.Meta, name string) ([]entity.Member, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq, name)
}

func (s *memberService) GetAllGrouped(ctx context.Context) ([]dto.MemberGroupResponse, error) {
	// Fetch all members (large limit)
	mReq := meta.Meta{Page: 1, Limit: 1000} 
	members, _, err := s.repo.GetAll(ctx, nil, mReq, "")
	if err != nil {
		return nil, err
	}

	var groups []dto.MemberGroupResponse
	var currentGroup *dto.MemberGroupResponse
	
	for _, m := range members {
		if m.Role == nil {
			continue 
		}

		if currentGroup == nil || currentGroup.Role.Id != m.Role.Id {
			if currentGroup != nil {
				groups = append(groups, *currentGroup)
			}
			currentGroup = &dto.MemberGroupResponse{
				Role:    *m.Role,
				Members: []entity.Member{},
			}
		}
		currentGroup.Members = append(currentGroup.Members, m)
	}
	if currentGroup != nil {
		groups = append(groups, *currentGroup)
	}
	
	return groups, nil
}

func (s *memberService) GetById(ctx context.Context, id string) (entity.Member, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.Member{}, err
	}
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
	if req.RoleId != nil {
		m.RoleId = req.RoleId
		m.Role = nil
	}
	if req.DepartmentId != nil {
		m.DepartmentId = req.DepartmentId
		m.Department = nil
	}
	if req.PhotoId != nil {
		m.PhotoId = req.PhotoId
		m.Photo = nil
	}
	if req.CabinetId != nil {
		m.CabinetId = req.CabinetId
	}
	m.Index = req.Index

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
