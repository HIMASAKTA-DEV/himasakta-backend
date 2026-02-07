package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, tx *gorm.DB, user entity.User, preloads ...string) (entity.User, error)
		GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, preloads ...string) ([]entity.User, meta.Meta, error)
		GetById(ctx context.Context, tx *gorm.DB, userId string, preloads ...string) (entity.User, error)
		GetByEmail(ctx context.Context, tx *gorm.DB, email string, preloads ...string) (entity.User, error)
		Update(ctx context.Context, tx *gorm.DB, user entity.User, preloads ...string) (entity.User, error)
		Delete(ctx context.Context, tx *gorm.DB, user entity.User) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUser(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, user entity.User, preloads ...string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetAll(ctx context.Context, tx *gorm.DB, metaReq meta.Meta, preloads ...string) ([]entity.User, meta.Meta, error) {
	if tx == nil {
		tx = r.db
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	var users []entity.User

	tx = tx.WithContext(ctx).Model(entity.User{})

	if err := WithFilters(tx, &metaReq,
		AddModels(entity.User{})).Find(&users).Error; err != nil {
		return nil, metaReq, err
	}

	return users, metaReq, nil
}

func (r *userRepository) GetById(ctx context.Context, tx *gorm.DB, userId string, preloads ...string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	var user entity.User
	if err := tx.WithContext(ctx).Take(&user, "id = ?", userId).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, tx *gorm.DB, email string, preloads ...string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	var user entity.User
	if err := tx.WithContext(ctx).Take(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, myerror.New("user not found", http.StatusNotFound)
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, tx *gorm.DB, user entity.User, preloads ...string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	if err := tx.WithContext(ctx).Save(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, tx *gorm.DB, user entity.User) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

