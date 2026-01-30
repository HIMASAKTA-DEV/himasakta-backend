package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/azkaazkun/be-samarta/internal/api/repository"
	"github.com/azkaazkun/be-samarta/internal/dto"
	"github.com/azkaazkun/be-samarta/internal/entity"
	"github.com/azkaazkun/be-samarta/internal/pkg/meta"
	"gorm.io/gorm"
)

type (
	SSHService interface {
		Create(ctx context.Context, req dto.CreateSSHRequest) (entity.SSH, error)
		GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.SSH, meta.Meta, error)
		GetById(ctx context.Context, id string) (entity.SSH, error)
		Update(ctx context.Context, req dto.UpdateSSHRequest) (entity.SSH, error)
		Delete(ctx context.Context, id string) error
	}

	sshService struct {
		sshRepository      repository.SSHRepository
		itemRepository     repository.ItemRepository
		proposalRepository repository.ProposalRepository
		db                 *gorm.DB
	}
)

func NewSSH(sshRepository repository.SSHRepository, itemRepository repository.ItemRepository, proposalRepository repository.ProposalRepository, db *gorm.DB) SSHService {
	return &sshService{
		sshRepository:      sshRepository,
		itemRepository:     itemRepository,
		proposalRepository: proposalRepository,
		db:                 db,
	}
}

func (s *sshService) Create(ctx context.Context, req dto.CreateSSHRequest) (entity.SSH, error) {
	item, err := s.itemRepository.GetById(ctx, nil, req.ItemID)
	if err != nil {
		return entity.SSH{}, err
	}

	proposalID, err := uuid.Parse(req.ProposalID)
	if err != nil {
		return entity.SSH{}, err
	}

	newSSH, err := s.sshRepository.Create(ctx, nil, entity.SSH{
		ProposalID:     proposalID,
		ItemID:         item.ID,
		UraianKomponen: req.UraianKomponen,
		Spesifikasi:    req.Spesifikasi,
		Satuan:         req.Satuan,
		HargaSatuan:    req.HargaSatuan,
		Rekening:       req.Rekening,
		Kel:            req.Kel,
		StatusKomponen: req.StatusKomponen,
		StatusSIPD:     req.StatusSIPD,
		DasarHarga:     req.DasarHarga,
	})
	if err != nil {
		return entity.SSH{}, err
	}

	return newSSH, nil
}

func (s *sshService) GetAll(ctx context.Context, metaReq meta.Meta) ([]entity.SSH, meta.Meta, error) {
	sshList, newMeta, err := s.sshRepository.GetAll(ctx, nil, metaReq, "Item", "Proposal")
	if err != nil {
		return nil, metaReq, err
	}

	return sshList, newMeta, nil
}

func (s *sshService) GetById(ctx context.Context, id string) (entity.SSH, error) {
	ssh, err := s.sshRepository.GetById(ctx, nil, id)
	if err != nil {
		return entity.SSH{}, err
	}

	return ssh, nil
}

func (s *sshService) Update(ctx context.Context, req dto.UpdateSSHRequest) (entity.SSH, error) {
	ssh, err := s.sshRepository.GetById(ctx, nil, req.ID)
	if err != nil {
		return entity.SSH{}, err
	}

	if req.ProposalID != "" {
		proposalID, err := uuid.Parse(req.ProposalID)
		if err != nil {
			return entity.SSH{}, err
		}
		ssh.ProposalID = proposalID
	}
	if req.ItemID != "" {
		item, err := s.itemRepository.GetById(ctx, nil, req.ItemID)
		if err != nil {
			return entity.SSH{}, err
		}
		ssh.ItemID = item.ID
	}
	if req.UraianKomponen != "" {
		ssh.UraianKomponen = req.UraianKomponen
	}
	if req.Spesifikasi != "" {
		ssh.Spesifikasi = req.Spesifikasi
	}
	if req.Satuan != "" {
		ssh.Satuan = req.Satuan
	}
	if req.HargaSatuan != 0 {
		ssh.HargaSatuan = req.HargaSatuan
	}
	if req.Rekening != "" {
		ssh.Rekening = req.Rekening
	}
	if req.Kel != 0 {
		ssh.Kel = req.Kel
	}
	if req.StatusKomponen != "" {
		ssh.StatusKomponen = req.StatusKomponen
	}
	if req.StatusSIPD != "" {
		ssh.StatusSIPD = req.StatusSIPD
	}
	if req.DasarHarga != "" {
		ssh.DasarHarga = req.DasarHarga
	}

	updatedSSH, err := s.sshRepository.Update(ctx, nil, ssh)
	if err != nil {
		return entity.SSH{}, err
	}

	return updatedSSH, nil
}

func (s *sshService) Delete(ctx context.Context, id string) error {
	ssh, err := s.sshRepository.GetById(ctx, nil, id)
	if err != nil {
		return err
	}

	if err := s.sshRepository.Delete(ctx, nil, ssh); err != nil {
		return err
	}

	return nil
}
