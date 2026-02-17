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
	GetGroupedByRank(ctx context.Context, metaReq meta.Meta) ([]dto.GroupedMemberResponse, error)
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
		Index:        req.Index,
		DepartmentId: req.DepartmentId,
		PhotoId:      req.PhotoId,
		CabinetId:    req.CabinetId,
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
	if req.RoleId != nil {
		m.RoleId = req.RoleId
	}
	if req.Index != 0 {
		m.Index = req.Index
	}
	if req.CabinetId != nil {
		m.CabinetId = req.CabinetId
	}
	m.DepartmentId = req.DepartmentId
	m.PhotoId = req.PhotoId

	return s.repo.Update(ctx, nil, m)
}

func (s *memberService) GetGroupedByRank(ctx context.Context, metaReq meta.Meta) ([]dto.GroupedMemberResponse, error) {
	members, _, err := s.repo.GetAll(ctx, nil, metaReq, "")
	if err != nil {
		return nil, err
	}

	var response []dto.GroupedMemberResponse
	if len(members) == 0 {
		return response, nil
	}

	// Grouping logic assuming members are sorted by Role.Index
	var currentGroup *dto.GroupedMemberResponse

	for _, m := range members {
		rankName := ""
		if m.Role != nil {
			rankName = m.Role.Rank
		}

		if currentGroup == nil || currentGroup.Rank != rankName {
			if currentGroup != nil {
				response = append(response, *currentGroup)
			}
			currentGroup = &dto.GroupedMemberResponse{
				Rank:    rankName,
				Members: []entity.Member{},
			}
		}
		currentGroup.Members = append(currentGroup.Members, m)
	}
	if currentGroup != nil {
		response = append(response, *currentGroup)
	}

	return response, nil
}

func (s *memberService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	m, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, m)
}
