package repository

import (
	"context"

	"github.com/azkaazkun/be-samarta/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, tx *gorm.DB, refreshToken entity.RefreshToken) (entity.RefreshToken, error)
	GetByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (entity.RefreshToken, error)
	GetAllByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]entity.RefreshToken, error)
	GetByUserIDAndUserAgent(ctx context.Context, tx *gorm.DB, userID uuid.UUID, userAgent string) (entity.RefreshToken, error)
	GetByRefreshTokenHash(ctx context.Context, tx *gorm.DB, refreshTokenHash string) (entity.RefreshToken, error)
	Delete(ctx context.Context, tx *gorm.DB, id uuid.UUID) error
	DeleteByUserIDAndUserAgent(ctx context.Context, tx *gorm.DB, userID uuid.UUID, userAgent string) error
}

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{
		db: db,
	}
}

func (r *refreshTokenRepository) Create(ctx context.Context, tx *gorm.DB, refreshToken entity.RefreshToken) (entity.RefreshToken, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&refreshToken).Error; err != nil {
		return entity.RefreshToken{}, err
	}

	return refreshToken, nil
}

func (r *refreshTokenRepository) GetByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (entity.RefreshToken, error) {
	if tx == nil {
		tx = r.db
	}

	var refreshToken entity.RefreshToken
	if err := tx.WithContext(ctx).Where("user_id = ?", userID).First(&refreshToken).Error; err != nil {
		return entity.RefreshToken{}, err
	}

	return refreshToken, nil
}

func (r *refreshTokenRepository) GetAllByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]entity.RefreshToken, error) {
	if tx == nil {
		tx = r.db
	}

	var refreshTokens []entity.RefreshToken
	if err := tx.WithContext(ctx).Where("user_id = ?", userID).Find(&refreshTokens).Error; err != nil {
		return nil, err
	}

	return refreshTokens, nil
}

func (r *refreshTokenRepository) GetByUserIDAndUserAgent(ctx context.Context, tx *gorm.DB, userID uuid.UUID, userAgent string) (entity.RefreshToken, error) {
	if tx == nil {
		tx = r.db
	}

	var refreshToken entity.RefreshToken
	if err := tx.WithContext(ctx).Where("user_id = ? AND user_agent = ?", userID, userAgent).First(&refreshToken).Error; err != nil {
		return entity.RefreshToken{}, err
	}

	return refreshToken, nil
}

func (r *refreshTokenRepository) GetByRefreshTokenHash(ctx context.Context, tx *gorm.DB, refreshTokenHash string) (entity.RefreshToken, error) {
	if tx == nil {
		tx = r.db
	}

	var refreshToken entity.RefreshToken
	if err := tx.WithContext(ctx).Where("refresh_token_hash = ?", refreshTokenHash).First(&refreshToken).Error; err != nil {
		return entity.RefreshToken{}, err
	}

	return refreshToken, nil
}

func (r *refreshTokenRepository) DeleteByUserIDAndUserAgent(ctx context.Context, tx *gorm.DB, userID uuid.UUID, userAgent string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("user_id = ? AND user_agent = ?", userID, userAgent).Unscoped().Delete(&entity.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *refreshTokenRepository) Delete(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Unscoped().Delete(&entity.RefreshToken{}, id).Error; err != nil {
		return err
	}

	return nil
}
