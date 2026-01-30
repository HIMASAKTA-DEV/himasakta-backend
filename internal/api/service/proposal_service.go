package service

import (
	"context"

	"github.com/azkaazkun/be-samarta/internal/api/repository"
	"github.com/azkaazkun/be-samarta/internal/dto"
	"github.com/azkaazkun/be-samarta/internal/entity"
	"github.com/azkaazkun/be-samarta/internal/pkg/meta"
	"gorm.io/gorm"
)

type (
	ProposalService interface {
		Create(ctx context.Context, req dto.CreateProposalRequest) (entity.Proposal, error)
		GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.Proposal, meta.Meta, error)
		GetById(ctx context.Context, id string) (entity.Proposal, error)
		Update(ctx context.Context, req dto.UpdateProposalRequest) (entity.Proposal, error)
		Delete(ctx context.Context, id string) error
	}

	proposalService struct {
		proposalRepository repository.ProposalRepository
		db                 *gorm.DB
	}
)

func NewProposal(proposalRepository repository.ProposalRepository, db *gorm.DB) ProposalService {
	return &proposalService{
		proposalRepository: proposalRepository,
		db:                 db,
	}
}

func (s *proposalService) Create(ctx context.Context, req dto.CreateProposalRequest) (entity.Proposal, error) {
	newProposal, err := s.proposalRepository.Create(ctx, nil, entity.Proposal{
		No:            req.No,
		ComponentName: req.ComponentName,
		Specification: req.Specification,
		Unit:          req.Unit,
		UnitPrice:     req.UnitPrice,
		AccountCode:   req.AccountCode,
	})
	if err != nil {
		return entity.Proposal{}, err
	}

	return newProposal, nil
}

func (s *proposalService) GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.Proposal, meta.Meta, error) {
	proposalList, newMeta, err := s.proposalRepository.GetAll(ctx, nil, metaReq)
	if err != nil {
		return nil, metaReq, err
	}

	return proposalList, newMeta, nil
}

func (s *proposalService) GetById(ctx context.Context, id string) (entity.Proposal, error) {
	proposal, err := s.proposalRepository.GetById(ctx, nil, id)
	if err != nil {
		return entity.Proposal{}, err
	}

	return proposal, nil
}

func (s *proposalService) Update(ctx context.Context, req dto.UpdateProposalRequest) (entity.Proposal, error) {
	proposal, err := s.proposalRepository.GetById(ctx, nil, req.ID)
	if err != nil {
		return entity.Proposal{}, err
	}

	if req.No != "" {
		proposal.No = req.No
	}
	if req.ComponentName != "" {
		proposal.ComponentName = req.ComponentName
	}
	if req.Specification != "" {
		proposal.Specification = req.Specification
	}
	if req.Unit != "" {
		proposal.Unit = req.Unit
	}
	if req.UnitPrice != 0 {
		proposal.UnitPrice = req.UnitPrice
	}
	if req.AccountCode != "" {
		proposal.AccountCode = req.AccountCode
	}

	updatedProposal, err := s.proposalRepository.Update(ctx, nil, proposal)
	if err != nil {
		return entity.Proposal{}, err
	}

	return updatedProposal, nil
}

func (s *proposalService) Delete(ctx context.Context, id string) error {
	proposal, err := s.proposalRepository.GetById(ctx, nil, id)
	if err != nil {
		return err
	}

	if err := s.proposalRepository.Delete(ctx, nil, proposal); err != nil {
		return err
	}

	return nil
}
