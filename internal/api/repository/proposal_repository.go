package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/azkaazkun/be-samarta/internal/entity"
	myerror "github.com/azkaazkun/be-samarta/internal/pkg/error"
	"github.com/azkaazkun/be-samarta/internal/pkg/meta"
	"gorm.io/gorm"
)

type (
	ProposalRepository interface {
		Create(ctx context.Context, tx *gorm.DB, proposal entity.Proposal) (entity.Proposal, error)
		GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.Proposal, meta.Meta, error)
		GetById(ctx context.Context, tx *gorm.DB, id string) (entity.Proposal, error)
		Update(ctx context.Context, tx *gorm.DB, proposal entity.Proposal) (entity.Proposal, error)
		Delete(ctx context.Context, tx *gorm.DB, proposal entity.Proposal) error
	}

	proposalRepository struct {
		db *gorm.DB
	}
)

func NewProposal(db *gorm.DB) ProposalRepository {
	return &proposalRepository{db}
}

func (r *proposalRepository) Create(ctx context.Context, tx *gorm.DB, proposal entity.Proposal) (entity.Proposal, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&proposal).Error; err != nil {
		return entity.Proposal{}, err
	}

	return proposal, nil
}

func (r *proposalRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta) ([]entity.Proposal, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}

	var proposalList []entity.Proposal

	tx = tx.WithContext(ctx).Model(entity.Proposal{})

	if err := WithFilters(tx, &metaReq,
		AddModels(entity.Proposal{})).Find(&proposalList).Error; err != nil {
		return nil, metaReq, err
	}

	return proposalList, metaReq, nil
}

func (r *proposalRepository) GetById(ctx context.Context, tx *gorm.DB, id string) (entity.Proposal, error) {
	if tx == nil {
		tx = r.db
	}

	var proposal entity.Proposal
	if err := tx.WithContext(ctx).Take(&proposal, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Proposal{}, myerror.New("proposal not found", http.StatusNotFound)
		}
		return entity.Proposal{}, err
	}

	return proposal, nil
}

func (r *proposalRepository) Update(ctx context.Context, tx *gorm.DB, proposal entity.Proposal) (entity.Proposal, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&proposal).Error; err != nil {
		return entity.Proposal{}, err
	}

	return proposal, nil
}

func (r *proposalRepository) Delete(ctx context.Context, tx *gorm.DB, proposal entity.Proposal) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&proposal).Error; err != nil {
		return err
	}

	return nil
}
