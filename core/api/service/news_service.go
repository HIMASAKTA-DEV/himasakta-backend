package service

import (
	"context"
	"errors"
	"strings"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsService interface {
	Create(ctx context.Context, req dto.CreateNewsRequest) (entity.News, error)
	GetAll(ctx context.Context, metaReq meta.Meta, search string, category string, title string) ([]entity.News, meta.Meta, error)
	GetBySlug(ctx context.Context, slugOrId string) (entity.News, error)
	Update(ctx context.Context, id string, req dto.UpdateNewsRequest) (entity.News, error)
	Delete(ctx context.Context, id string) error
	GetAutocompletion(ctx context.Context, query string) ([]string, error)
}

type newsService struct {
	repo repository.NewsRepository
}

func NewNews(repo repository.NewsRepository) NewsService {
	return &newsService{repo}
}

func (s *newsService) Create(ctx context.Context, req dto.CreateNewsRequest) (entity.News, error) {
	return s.repo.Create(ctx, nil, entity.News{
		Title:       req.Title,
		Slug:        utils.ToSlug(req.Title),
		Tagline:     req.Tagline,
		Hashtags:    req.Hashtags,
		Content:     req.Content,
		ThumbnailId: req.ThumbnailId,
		PublishedAt: req.PublishedAt,
	})
}

func (s *newsService) GetAll(ctx context.Context, metaReq meta.Meta, search string, category string, title string) ([]entity.News, meta.Meta, error) {
	var categories []string
	if category != "" {
		categories = strings.Split(category, ",")
	}
	return s.repo.GetAll(ctx, nil, metaReq, search, categories, title)
}

func (s *newsService) GetAutocompletion(ctx context.Context, query string) ([]string, error) {
	return s.repo.GetAutocompletion(ctx, query)
}

func (s *newsService) GetBySlug(ctx context.Context, slugOrId string) (entity.News, error) {
	var n entity.News
	var err error

	uid, parseErr := uuid.Parse(slugOrId)
	if parseErr == nil {
		n, err = s.repo.GetById(ctx, nil, uid)
	} else {
		n, err = s.repo.GetBySlug(ctx, nil, slugOrId)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return n, myerror.ErrNotFound
		}
		return n, err
	}
	return n, nil
}

func (s *newsService) GetById(ctx context.Context, id string) (entity.News, error) {
	uid, _ := uuid.Parse(id)
	return s.repo.GetById(ctx, nil, uid)
}

func (s *newsService) Update(ctx context.Context, id string, req dto.UpdateNewsRequest) (entity.News, error) {
	uid, _ := uuid.Parse(id)
	n, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return n, err
	}

	if req.Title != "" {
		n.Title = req.Title
		n.Slug = utils.ToSlug(req.Title)
	}
	n.Tagline = req.Tagline
	n.Hashtags = req.Hashtags
	if req.Content != "" {
		n.Content = req.Content
	}
	if req.ThumbnailId != nil {
		n.ThumbnailId = req.ThumbnailId
		n.Thumbnail = nil
	}
	if req.PublishedAt != nil {
		n.PublishedAt = *req.PublishedAt
	}

	return s.repo.Update(ctx, nil, n)
}

func (s *newsService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	n, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, n)
}
