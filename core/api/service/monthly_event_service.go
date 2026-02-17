package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
	"github.com/google/uuid"
)

type MonthlyEventService interface {
	Create(ctx context.Context, req dto.CreateMonthlyEventRequest) (entity.MonthlyEvent, error)
	GetAll(ctx context.Context, metaReq meta.Meta, title string) ([]entity.MonthlyEvent, meta.Meta, error)
	GetThisMonth(ctx context.Context) ([]entity.MonthlyEvent, error)
	GetById(ctx context.Context, id string) (entity.MonthlyEvent, error)
	Update(ctx context.Context, id string, req dto.UpdateMonthlyEventRequest) (entity.MonthlyEvent, error)
	Delete(ctx context.Context, id string) error
}

type monthlyEventService struct {
	repo repository.MonthlyEventRepository
}

func NewMonthlyEvent(repo repository.MonthlyEventRepository) MonthlyEventService {
	return &monthlyEventService{repo}
}

func (s *monthlyEventService) Create(ctx context.Context, req dto.CreateMonthlyEventRequest) (entity.MonthlyEvent, error) {
	return s.repo.Create(ctx, nil, entity.MonthlyEvent{
		Title:       req.Title,
		ThumbnailId: req.ThumbnailId,
		Description: req.Description,
		Month:       req.Month,
		Link:        req.Link,
	})
}

func (s *monthlyEventService) GetAll(ctx context.Context, metaReq meta.Meta, title string) ([]entity.MonthlyEvent, meta.Meta, error) {
	return s.repo.GetAll(ctx, nil, metaReq, title)
}

func (s *monthlyEventService) GetThisMonth(ctx context.Context) ([]entity.MonthlyEvent, error) {
	return s.repo.GetThisMonth(ctx, nil)
}

func (s *monthlyEventService) GetById(ctx context.Context, id string) (entity.MonthlyEvent, error) {
	uid, _ := uuid.Parse(id)
	return s.repo.GetById(ctx, nil, uid)
}

func (s *monthlyEventService) Update(ctx context.Context, id string, req dto.UpdateMonthlyEventRequest) (entity.MonthlyEvent, error) {
	uid, _ := uuid.Parse(id)
	e, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return e, err
	}

	if req.Title != "" {
		e.Title = req.Title
	}
	if req.ThumbnailId != nil {
		e.ThumbnailId = req.ThumbnailId
		e.Thumbnail = nil
	}
	e.Description = req.Description
	if req.Month != nil {
		e.Month = *req.Month
	}
	e.Link = req.Link

	return s.repo.Update(ctx, nil, e)
}

func (s *monthlyEventService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	e, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, e)
}
