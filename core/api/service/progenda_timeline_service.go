package service

import (
	"context"
	"errors"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgendaTimelineService interface {
	Create(ctx context.Context, progendaId string, req dto.ProgendaTimelineRequest) (entity.ProgendaTimeline, error)
	GetAll(ctx context.Context, progendaId string) ([]entity.ProgendaTimeline, error)
	Update(ctx context.Context, id string, req dto.ProgendaTimelineRequest) (entity.ProgendaTimeline, error)
	Delete(ctx context.Context, id string) error
}

type progendaTimelineService struct {
	repo         repository.ProgendaTimelineRepository
	progendaRepo repository.ProgendaRepository
}

func NewProgendaTimeline(repo repository.ProgendaTimelineRepository, progendaRepo repository.ProgendaRepository) ProgendaTimelineService {
	return &progendaTimelineService{repo, progendaRepo}
}

func (s *progendaTimelineService) Create(ctx context.Context, progendaId string, req dto.ProgendaTimelineRequest) (entity.ProgendaTimeline, error) {
	pid, _ := uuid.Parse(progendaId)
	if _, err := s.progendaRepo.GetById(ctx, nil, pid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.ProgendaTimeline{}, myerror.ErrNotFound
		}
		return entity.ProgendaTimeline{}, err
	}

	return s.repo.Create(ctx, nil, entity.ProgendaTimeline{
		ProgendaId: pid,
		EventName:  req.EventName,
		Date:       req.Date,
	})
}

func (s *progendaTimelineService) GetAll(ctx context.Context, progendaId string) ([]entity.ProgendaTimeline, error) {
	pid, _ := uuid.Parse(progendaId)
	return s.repo.GetAllByProgendaId(ctx, nil, pid)
}

func (s *progendaTimelineService) Update(ctx context.Context, id string, req dto.ProgendaTimelineRequest) (entity.ProgendaTimeline, error) {
	uid, _ := uuid.Parse(id)
	t, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return t, myerror.ErrNotFound
		}
		return t, err
	}

	t.EventName = req.EventName
	t.Date = req.Date
	return s.repo.Update(ctx, nil, t)
}

func (s *progendaTimelineService) Delete(ctx context.Context, id string) error {
	uid, _ := uuid.Parse(id)
	t, err := s.repo.GetById(ctx, nil, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrNotFound
		}
		return err
	}
	return s.repo.Delete(ctx, nil, t)
}
