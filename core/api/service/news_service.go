package service

import (
	"context"
	"errors"

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
	db          *gorm.DB
	newsRepo    repository.NewsRepository
	newsTagRepo repository.NewsTagRepository
	tagRepo     repository.TagRepository
}

func NewNews(db *gorm.DB, newsRepo repository.NewsRepository, newsTagRepo repository.NewsTagRepository, tagRepo repository.TagRepository) NewsService {
	return &newsService{
		db:          db,
		newsRepo:    newsRepo,
		newsTagRepo: newsTagRepo,
		tagRepo:     tagRepo,
	}
}

func (s *newsService) Create(ctx context.Context, req dto.CreateNewsRequest) (entity.News, error) {
	tx := s.db.Begin()
	newsId := uuid.New()

	hashtags, err := utils.SanitizeHashtags(req.Hashtags)
	if err != nil {
		return entity.News{}, err
	}

	tagEntities, err := utils.SplitHashTags(hashtags)
	if err != nil {
		return entity.News{}, err
	}
	finalTags, err := s.tagRepo.BulkAdd(ctx, tx, tagEntities)
	if err != nil {
		tx.Rollback()
		return entity.News{}, err
	}

	var newsTag []entity.NewsTag

	for _, tag := range finalTags {
		newsTag = append(newsTag, entity.NewsTag{
			NewsId: newsId,
			TagId:  tag.Id,
		})
	}

	if err := s.newsTagRepo.BulkCreate(ctx, tx, newsTag); err != nil {
		tx.Rollback()
		return entity.News{}, err
	}

	res, err := s.newsRepo.Create(ctx, tx, entity.News{
		Id:          newsId,
		Title:       req.Title,
		Slug:        utils.ToSlug(req.Title),
		Tagline:     req.Tagline,
		Content:     req.Content,
		Hashtags:    finalTags,
		ThumbnailId: req.ThumbnailId.ID,
		PublishedAt: req.PublishedAt,
		AuthorId:    req.AuthorId.ID,
	})
	if err != nil {
		tx.Rollback()
		return entity.News{}, err
	}

	tx.Commit()

	return res, myerror.ParseDBError(err, "news")
}

func (s *newsService) GetAll(ctx context.Context, metaReq meta.Meta, search string, tags string, title string) ([]entity.News, meta.Meta, error) {
	hashtags, err := utils.SanitizeHashtags(tags)
	if err != nil {
		return []entity.News{}, metaReq, err
	}

	tagEntities, err := utils.SplitHashTags(hashtags)
	if err != nil {
		return []entity.News{}, metaReq, err
	}
	var tagList []string
	for _, tag := range tagEntities {
		tagList = append(tagList, tag.Name)
	}

	return s.newsRepo.GetAll(ctx, nil, metaReq, search, tagList, title)
}

func (s *newsService) GetAutocompletion(ctx context.Context, query string) ([]string, error) {
	return s.newsRepo.GetAutocompletion(ctx, query)
}

func (s *newsService) GetBySlug(ctx context.Context, slugOrId string) (entity.News, error) {
	var n entity.News
	var err error

	uid, parseErr := uuid.Parse(slugOrId)
	if parseErr == nil {
		n, err = s.newsRepo.GetById(ctx, nil, uid)
	} else {
		n, err = s.newsRepo.GetBySlug(ctx, nil, slugOrId)
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
	return s.newsRepo.GetById(ctx, nil, uid)
}

func (s *newsService) Update(ctx context.Context, id string, req dto.UpdateNewsRequest) (entity.News, error) {
	tx := s.db.Begin()

	uid, _ := uuid.Parse(id)
	n, err := s.newsRepo.GetById(ctx, tx, uid)

	if err != nil {
		tx.Rollback()
		return n, err
	}

	var finalTags []entity.Tag
	if req.Hashtags != nil {
		//bulkdelete
		if err := s.newsTagRepo.DeleteByNews(ctx, tx, uid); err != nil {
			tx.Rollback()
			return entity.News{}, err
		}

		if *req.Hashtags != "" {

			//create and bulkadd entity.Tags
			hashtags, err := utils.SanitizeHashtags(*req.Hashtags)
			if err != nil {
				tx.Rollback()
				return entity.News{}, err
			}

			tagEntities, err := utils.SplitHashTags(hashtags)
			if err != nil {
				tx.Rollback()
				return entity.News{}, err
			}

			finalTags, err = s.tagRepo.BulkAdd(ctx, tx, tagEntities)
			if err != nil {
				tx.Rollback()
				return entity.News{}, err
			}
			//create and bulkcreate entity.NewsTag
			var newsTag []entity.NewsTag

			for _, tag := range finalTags {
				newsTag = append(newsTag, entity.NewsTag{
					NewsId: uid,
					TagId:  tag.Id,
				})
			}

			if err := s.newsTagRepo.BulkCreate(ctx, tx, newsTag); err != nil {
				tx.Rollback()
				return entity.News{}, err
			}
		}
		n.Hashtags = finalTags
	}

	if req.Title != nil {
		n.Title = *req.Title
		n.Slug = utils.ToSlug(*req.Title)
	}
	if req.Tagline != nil {
		n.Tagline = *req.Tagline
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
	if req.AuthorId.Valid {
		n.AuthorId = req.AuthorId.ID
		n.Author = nil
	}

	res, err := s.newsRepo.Update(ctx, tx, n)
	if err != nil {
		tx.Rollback()
		return entity.News{}, err
	}
	tx.Commit()
	res.Hashtags = finalTags
	return res, myerror.ParseDBError(err, "news")
}

func (s *newsService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	n, err := s.newsRepo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.newsRepo.Delete(ctx, nil, n)
}
