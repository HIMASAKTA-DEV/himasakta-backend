package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
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
	res, err := s.repo.Create(ctx, nil, entity.Member{
		Nrp:          req.Nrp,
		Name:         req.Name,
		RoleId:       req.RoleId.ID,
		DepartmentId: req.DepartmentId.ID,
		PhotoId:      req.PhotoId.ID,
		CabinetId:    req.CabinetId.ID,
		Index:        req.Index,
	})
	return res, myerror.ParseDBError(err, "member")
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

	if req.Nrp != nil {
		m.Nrp = *req.Nrp
	}
	if req.Name != nil {
		m.Name = *req.Name
	}
	if req.RoleId.Valid {
		m.RoleId = req.RoleId.ID
		m.Role = nil
	}
	if req.DepartmentId.Valid {
		m.DepartmentId = req.DepartmentId.ID
		m.Department = nil
	}
	if req.PhotoId.Valid {
		m.PhotoId = req.PhotoId.ID
		m.Photo = nil
	}
	if req.CabinetId.Valid {
		m.CabinetId = req.CabinetId.ID
	}
	if req.Index != nil {
		m.Index = *req.Index
	}

	res, err := s.repo.Update(ctx, nil, m)
	return res, myerror.ParseDBError(err, "member")
}

func (s *memberService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	m, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, m)
}
