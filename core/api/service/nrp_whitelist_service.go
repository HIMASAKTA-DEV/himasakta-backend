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

type NrpWhitelistService interface {
	Create(ctx context.Context, req dto.CreateNrpWhitelistRequest) (entity.NrpWhitelist, error)
	Check(ctx context.Context, req dto.CheckNrpWhitelistRequest) (entity.NrpWhitelist, error)
	GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.NrpWhitelist, meta.Meta, error)
	Update(ctx context.Context, id string, req dto.UpdateNrpWhitelistRequest) (entity.NrpWhitelist, error)
	Delete(ctx context.Context, nrp string) error
}

type nrpWhitelistService struct {
	repo repository.NrpWhitelistRepository
}

func NewNrpWhitelist(repo repository.NrpWhitelistRepository) NrpWhitelistService {
	return &nrpWhitelistService{repo}
}
func (s *nrpWhitelistService) Create(ctx context.Context, req dto.CreateNrpWhitelistRequest) (entity.NrpWhitelist, error) {
	// Check if already exist
	existing, err := s.repo.GetByNrp(ctx, nil, entity.NrpWhitelist{Nrp: req.Nrp})
	if err == nil && existing.Nrp != "" {
		return entity.NrpWhitelist{}, myerror.ErrAlreadyExist
	}

	return s.repo.Create(ctx, nil, entity.NrpWhitelist{
		Nrp:  req.Nrp,
		Name: req.Name,
	})
}

func (s *nrpWhitelistService) Check(ctx context.Context, req dto.CheckNrpWhitelistRequest) (entity.NrpWhitelist, error) {
	res, err := s.repo.GetByNrp(ctx, nil, entity.NrpWhitelist{
		Nrp: req.Nrp,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, myerror.ErrNotFound
		}
		return res, err
	}
	return res, nil
}

func (s *nrpWhitelistService) GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.NrpWhitelist, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq)
}

func (s *nrpWhitelistService) Update(ctx context.Context, id string, req dto.UpdateNrpWhitelistRequest) (entity.NrpWhitelist, error) {
	uid, _ := uuid.Parse(id)
	ci, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ci, myerror.ErrNotFound
		}
		return ci, err
	}
	if req.Nrp != "" {
		ci.Nrp = req.Nrp
	}
	if req.Name != "" {
		ci.Name = req.Name
	}

	return s.repo.Update(ctx, nil, ci)
}

func (s *nrpWhitelistService) Delete(ctx context.Context, nrp string) error {
	ci, err := s.repo.GetByNrp(ctx, nil, entity.NrpWhitelist{
		Nrp: nrp,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrNotFound
		}
		return err
	}
	return s.repo.Delete(ctx, nil, ci)
}
