package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	mailer "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/email"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	myjwt "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/jwt"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"

	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
)

type (
	AuthService interface {
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
		Register(ctx context.Context, req dto.RegisterRequest) (dto.UserResponse, error)
		VerifyEmail(ctx context.Context, token string) error
		RefreshToken(ctx context.Context, refreshToken string) (dto.LoginResponse, error)
		Logout(ctx context.Context, req dto.LogoutRequest) error
		ForgetPassword(ctx context.Context, req dto.ForgetPasswordRequest) error
		ChangePassword(ctx context.Context, req dto.ChangePasswordRequest) error
		GetMe(ctx context.Context, userId string) (dto.GetMe, error)
	}

	authService struct {
		userRepository         repository.UserRepository
		refreshTokenRepository repository.RefreshTokenRepository
		mailService            mailer.Mailer
		db                     *gorm.DB
	}
)

func NewAuth(userRepository repository.UserRepository,
	refreshTokenRepository repository.RefreshTokenRepository,
	mailService mailer.Mailer,
	db *gorm.DB) AuthService {
	return &authService{
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
		mailService:            mailService,
		db:                     db,
	}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if !user.IsVerified {
		return dto.LoginResponse{}, myerror.New("user is not verify", http.StatusUnauthorized)
	}

	checkPassword, err := utils.CheckPassword(user.Password, []byte(req.Password))
	if !checkPassword || err != nil {
		return dto.LoginResponse{}, myerror.New("email or password invalid", http.StatusBadRequest)
	}

	token, err := myjwt.GenerateToken(map[string]string{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    string(user.Role),
	}, 24*time.Hour)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	resp := dto.LoginResponse{
		AccessToken: token,
		Role:        string(user.Role),
	}

	if req.RememberMe {
		_ = s.refreshTokenRepository.DeleteByUserIDAndUserAgent(ctx, nil, user.ID, req.UserAgent)

		refreshTokenJTI, err := utils.GenerateRandomString(32)
		if err != nil {
			return dto.LoginResponse{}, err
		}

		refreshToken, err := myjwt.GenerateToken(map[string]string{
			"user_id": user.ID.String(),
			"role":    string(user.Role),
			"type":    "refresh",
			"jti":     refreshTokenJTI,
		}, 30*24*time.Hour)
		if err != nil {
			return dto.LoginResponse{}, err
		}

		_, err = s.refreshTokenRepository.Create(ctx, nil, entity.RefreshToken{
			UserID:       user.ID,
			RefreshToken: refreshToken,
			UserAgent:    req.UserAgent,
			IP:           req.IP,
			ExpiresAt:    time.Now().Add(30 * 24 * time.Hour), // 30 days
		})
		if err != nil {
			return dto.LoginResponse{}, err
		}

		resp.RefreshToken = &refreshToken
	}

	return resp, nil
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.UserResponse, error) {
	// check existing email
	_, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err == nil {
		return dto.UserResponse{}, myerror.New("email already exists", http.StatusConflict)
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user := entity.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashedPassword,
		Role:       entity.RoleUser,
		IsVerified: false,
	}

	newUser, err := s.userRepository.Create(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Generate verification token
	token, err := myjwt.GenerateToken(map[string]string{
		"user_id": newUser.ID.String(),
		"email":   newUser.Email,
		"type":    "verification",
	}, 24*time.Hour)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Send verification email
	verifyLink := fmt.Sprintf("%s/verify-email/%s", os.Getenv("APP_URL"), token)
	if err := s.mailService.MakeMail("./core/pkg/email/template/verification_email.html", map[string]any{
		"Name":   newUser.Name,
		"Verify": verifyLink,
	}).Send(newUser.Email, "Verify Your Account").Error; err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:    newUser.ID.String(),
		Name:  newUser.Name,
		Email: newUser.Email,
		Role:  string(newUser.Role),
	}, nil
}

func (s *authService) VerifyEmail(ctx context.Context, token string) error {
	claims, err := myjwt.GetPayloadInsideToken(token)
	if err != nil {
		return myerror.New("invalid token", http.StatusBadRequest)
	}

	if claims["type"] != "verification" {
		return myerror.New("invalid token type", http.StatusBadRequest)
	}

	user, err := s.userRepository.GetByEmail(ctx, nil, claims["email"])
	if err != nil {
		return err
	}

	if user.IsVerified {
		return nil
	}

	user.IsVerified = true
	_, err = s.userRepository.Update(ctx, nil, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (dto.LoginResponse, error) {
	// Parse JWT
	claims, err := myjwt.GetPayloadInsideToken(refreshToken)
	if err != nil {
		return dto.LoginResponse{}, myerror.New("refresh token invalid", http.StatusUnauthorized)
	}

	// Validate token type
	if claims["type"] != "refresh" {
		return dto.LoginResponse{}, myerror.New("invalid token type", http.StatusUnauthorized)
	}

	userID := claims["user_id"]

	// Find matching token directly from DB
	foundToken, err := s.refreshTokenRepository.GetByRefreshToken(ctx, nil, refreshToken)
	if err != nil {
		return dto.LoginResponse{}, myerror.New("refresh token invalid", http.StatusUnauthorized)
	}

	if foundToken.UserID.String() != userID {
		return dto.LoginResponse{}, myerror.New("refresh token invalid", http.StatusUnauthorized)
	}

	if foundToken.ExpiresAt.Before(time.Now()) {
		_ = s.refreshTokenRepository.Delete(ctx, nil, foundToken.ID)
		return dto.LoginResponse{}, myerror.New("refresh token expired", http.StatusUnauthorized)
	}

	user, err := s.userRepository.GetById(ctx, nil, userID)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// rotate token
	newJTI, err := utils.GenerateRandomString(32)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	newRefreshToken, err := myjwt.GenerateToken(map[string]string{
		"user_id": user.ID.String(),
		"role":    string(user.Role),
		"type":    "refresh",
		"jti":     newJTI,
	}, 30*24*time.Hour)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// create new token with same expiration
	_, err = s.refreshTokenRepository.Create(ctx, nil, entity.RefreshToken{
		UserID:       user.ID,
		RefreshToken: newRefreshToken,
		UserAgent:    foundToken.UserAgent,
		IP:           foundToken.IP,
		ExpiresAt:    foundToken.ExpiresAt,
	})
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// generate access token
	token, err := myjwt.GenerateToken(map[string]string{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    string(user.Role),
	}, 24*time.Hour)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		AccessToken:  token,
		RefreshToken: &newRefreshToken,
		Role:         string(user.Role),
	}, nil
}

func (s *authService) Logout(ctx context.Context, req dto.LogoutRequest) error {
	// Parse JWT
	claims, err := myjwt.GetPayloadInsideToken(req.RefreshToken)
	if err != nil {
		return myerror.New("refresh token invalid", http.StatusUnauthorized)
	}

	// Validate token type
	if claims["type"] != "refresh" {
		return myerror.New("invalid token type", http.StatusUnauthorized)
	}

	userID := claims["user_id"]

	// Get all refresh tokens for user
	// Find matching token directly from DB
	foundToken, err := s.refreshTokenRepository.GetByRefreshToken(ctx, nil, req.RefreshToken)
	if err != nil {
		return myerror.New("refresh token invalid", http.StatusUnauthorized)
	}

	if foundToken.UserID.String() != userID {
		return myerror.New("refresh token invalid", http.StatusUnauthorized)
	}

	if err := s.refreshTokenRepository.Delete(ctx, nil, foundToken.ID); err != nil {
		return err
	}

	return nil
}

func (s *authService) ForgetPassword(ctx context.Context, req dto.ForgetPasswordRequest) error {
	user, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err != nil {
		return err
	}

	if !user.IsVerified {
		return errors.New("user not verified")
	}

	token, err := myjwt.GenerateToken(map[string]string{
		"user_id": user.ID.String(),
		"email":   user.Email,
	}, 24*time.Hour)
	if err != nil {
		return err
	}

	// generate token
	token = fmt.Sprintf("%s/reset-password/%s", os.Getenv("APP_URL"), token)
	if err := s.mailService.MakeMail("./core/pkg/email/template/forget_password_email.html", map[string]any{
		"Fullname": user.Name,
		"Link":     token,
	}).Send(user.Email, "Forget Password").Error; err != nil {
		return err
	}

	return nil
}

func (s *authService) ChangePassword(ctx context.Context, req dto.ChangePasswordRequest) error {
	user, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	_, err = s.userRepository.Update(ctx, nil, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) GetMe(ctx context.Context, userId string) (dto.GetMe, error) {
	user, err := s.userRepository.GetById(ctx, nil, userId)
	if err != nil {
		return dto.GetMe{}, err
	}

	return dto.GetMe{
		PersonalInfo: dto.PersonalInfo{
			ID:    userId,
			Name:  user.Name,
			Email: user.Email,
			Role:  string(user.Role),
		},
	}, nil
}
