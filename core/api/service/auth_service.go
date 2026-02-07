package service

import (
	"context"
	"errors"
	"os"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	myjwt "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/jwt"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
}

type authService struct {
	jwtService myjwt.JWT
}

func NewAuth(jwtService myjwt.JWT) AuthService {
	return &authService{jwtService: jwtService}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	adminUser := os.Getenv("ADMIN_USERNAME")
	adminPass := os.Getenv("ADMIN_PASSWORD")

	if req.Username != adminUser || req.Password != adminPass {
		return dto.LoginResponse{}, errors.New("invalid username or password")
	}

	token, err := s.jwtService.CreateToken("superadmin", "superadmin")
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: token,
	}, nil
}
