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
	SSHRepository interface {
		Create(ctx context.Context, tx *gorm.DB, ssh entity.SSH) (entity.SSH, error)
		GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, preloads ...string) ([]entity.SSH, meta.Meta, error)
		GetById(ctx context.Context, tx *gorm.DB, id string) (entity.SSH, error)
		Update(ctx context.Context, tx *gorm.DB, ssh entity.SSH) (entity.SSH, error)
		Delete(ctx context.Context, tx *gorm.DB, ssh entity.SSH) error
	}

	sshRepository struct {
		db *gorm.DB
	}
)

func NewSSH(db *gorm.DB) SSHRepository {
	return &sshRepository{db}
}

func (r *sshRepository) Create(ctx context.Context, tx *gorm.DB, ssh entity.SSH) (entity.SSH, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&ssh).Error; err != nil {
		return entity.SSH{}, err
	}

	return ssh, nil
}

func (r *sshRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, preloads ...string) ([]entity.SSH, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}

	var sshList []entity.SSH

	tx = tx.WithContext(ctx).Model(entity.SSH{})

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	if err := WithFilters(tx, &metaReq,
		AddModels(entity.SSH{})).Find(&sshList).Error; err != nil {
		return nil, metaReq, err
	}

	return sshList, metaReq, nil
}

func (r *sshRepository) GetById(ctx context.Context, tx *gorm.DB, id string) (entity.SSH, error) {
	if tx == nil {
		tx = r.db
	}

	var ssh entity.SSH
	if err := tx.WithContext(ctx).Take(&ssh, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.SSH{}, myerror.New("ssh not found", http.StatusNotFound)
		}
		return entity.SSH{}, err
	}

	return ssh, nil
}

func (r *sshRepository) Update(ctx context.Context, tx *gorm.DB, ssh entity.SSH) (entity.SSH, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&ssh).Error; err != nil {
		return entity.SSH{}, err
	}

	return ssh, nil
}

func (r *sshRepository) Delete(ctx context.Context, tx *gorm.DB, ssh entity.SSH) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&ssh).Error; err != nil {
		return err
	}

	return nil
}
