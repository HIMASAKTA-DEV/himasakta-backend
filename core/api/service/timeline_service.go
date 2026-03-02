package service

import (
	"context"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/google/uuid"
)

type TimelineService interface {
	Create(ctx context.Context, progendaId uuid.UUID, req dto.CreateTimelineRequest) (entity.Timeline, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateTimelineRequest) (entity.Timeline, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type timelineService struct {
	repo repository.TimelineRepository
}

func NewTimeline(repo repository.TimelineRepository) TimelineService {
	return &timelineService{repo}
}

func (s *timelineService) Create(ctx context.Context, progendaId uuid.UUID, req dto.CreateTimelineRequest) (entity.Timeline, error) {
	res, err := s.repo.Create(ctx, nil, entity.Timeline{
		ProgendaId: &progendaId,
		Date:       req.Date,
		Info:       req.Info,
		Link:       req.Link,
	})
	return res, myerror.ParseDBError(err, "timeline")
}

func (s *timelineService) Update(ctx context.Context, id uuid.UUID, req dto.UpdateTimelineRequest) (entity.Timeline, error) {
	tl, err := s.repo.GetById(ctx, nil, id)
	if err != nil {
		return entity.Timeline{}, err
	}

	if req.Date != nil {
		tl.Date = *req.Date
	}
	if req.Info != nil {
		tl.Info = *req.Info
	}
	if req.Link != nil {
		tl.Link = *req.Link
	}

	res, err := s.repo.BulkUpdate(ctx, nil, []entity.Timeline{tl})
	if err != nil || len(res) == 0 {
		return entity.Timeline{}, myerror.ParseDBError(err, "timeline")
	}
	return res[0], nil
}

func (s *timelineService) Delete(ctx context.Context, id uuid.UUID) error {
	tl, err := s.repo.GetById(ctx, nil, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, nil, tl)
}
