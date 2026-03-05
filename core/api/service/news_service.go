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
	GetAll(ctx context.Context, metaReq meta.Meta, search string, tags string, title string) ([]entity.News, meta.Meta, error)
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
	hashtags, err := utils.SanitizeHashtags(req.Hashtags)
	if err != nil {
		return entity.News{}, err
	}

	res, err := s.repo.Create(ctx, nil, entity.News{
		Title:       req.Title,
		Slug:        utils.ToSlug(req.Title),
		Tagline:     req.Tagline,
		Hashtags:    hashtags,
		Content:     req.Content,
		ThumbnailId: req.ThumbnailId.ID,
		PublishedAt: req.PublishedAt,
		AuthorId:    req.AuthorId,
	})
	return res, myerror.ParseDBError(err, "news")
}

func (s *newsService) GetAll(ctx context.Context, metaReq meta.Meta, search string, tags string, title string) ([]entity.News, meta.Meta, error) {
	var tagsList []string
	if tags != "" {
		tagsList = strings.Split(tags, ",")
	}
	return s.repo.GetAll(ctx, nil, metaReq, search, tagsList, title)
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

	if req.Title != nil {
		n.Title = *req.Title
		n.Slug = utils.ToSlug(*req.Title)
	}
	if req.Tagline != nil {
		n.Tagline = *req.Tagline
	}
	if req.Hashtags != nil {
		hashtags, err := utils.SanitizeHashtags(*req.Hashtags)
		if err != nil {
			return n, err
		}
		n.Hashtags = hashtags
	}
	if req.Content != nil {
		n.Content = *req.Content
	}
	if req.ThumbnailId.Valid {
		n.ThumbnailId = req.ThumbnailId.ID
		n.Thumbnail = nil
	}
	if req.PublishedAt != nil {
		n.PublishedAt = *req.PublishedAt
	}
	if req.AuthorId != nil {
		n.AuthorId = req.AuthorId
		n.Author = nil
	}

	res, err := s.repo.Update(ctx, nil, n)
	return res, myerror.ParseDBError(err, "news")
}

func (s *newsService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	n, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, n)
}
