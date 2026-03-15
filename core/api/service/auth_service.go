package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myjwt "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	UpdateAuth(ctx context.Context, req dto.LoginRequest) error
}

type authService struct {
	jwtService        myjwt.JWT
	globalSettingRepo repository.GlobalSettingRepository
}

func NewAuth(jwtService myjwt.JWT, globalSettingRepo repository.GlobalSettingRepository) AuthService {
	return &authService{
		jwtService:        jwtService,
		globalSettingRepo: globalSettingRepo,
	}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	rawAuth, err := s.globalSettingRepo.GetByKey(ctx, "auth")
	if err != nil {
		return dto.LoginResponse{}, err
	}

	var auth dto.LoginRequest
	err = json.Unmarshal([]byte(rawAuth.Value), &auth)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if req.Username != auth.Username {
		return dto.LoginResponse{}, errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(req.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid username or password")
	}

	token, err := s.jwtService.CreateToken("superadmin", "superadmin", "superadmin")
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: token,
	}, nil
}

func (s *authService) UpdateAuth(ctx context.Context, req dto.LoginRequest) error {
	newUser := req.Username
	newPass := req.Password

	hash, _ := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)

	err := s.globalSettingRepo.Upsert(ctx, entity.GlobalSetting{
		Key: "auth",
		Value: `{
			"username": "` + newUser + `",
			"password": "` + string(hash) + `"

		}`,
	})

	if err != nil {
		return err
	}

	return nil
}
