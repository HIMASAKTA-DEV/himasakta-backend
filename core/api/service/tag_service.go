package service

import (
	"context"
	"strings"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
)

type TagService interface {
	GetAll(ctx context.Context, metaReq meta.Meta, search string) ([]entity.Tag, meta.Meta, error)
}

type tagService struct {
	repo repository.TagRepository
}

func NewTag(repo repository.TagRepository) TagService {
	return &tagService{repo}
}

func (s *tagService) GetAll(ctx context.Context, metaReq meta.Meta, search string) ([]entity.Tag, meta.Meta, error) {
	t := strings.TrimPrefix(search, "#")
	res, metaReq, err := s.repo.GetAll(ctx, nil, metaReq, t)
	if err != nil {
		return []entity.Tag{}, metaReq, err
	}

	return res, metaReq, nil
}
