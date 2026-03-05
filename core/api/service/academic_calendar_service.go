package service

import (
	"context"
	"sort"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/meta"
)

type AcademicCalendarService interface {
	GetCalendar(ctx context.Context, month, year int) ([]dto.AcademicCalendarItem, error)
}

type academicCalendarService struct {
	monthlyEventRepo repository.MonthlyEventRepository
	timelineRepo     repository.TimelineRepository
}

func NewAcademicCalendar(me repository.MonthlyEventRepository, t repository.TimelineRepository) AcademicCalendarService {
	return &academicCalendarService{
		monthlyEventRepo: me,
		timelineRepo:     t,
	}
}

func (s *academicCalendarService) GetCalendar(ctx context.Context, month, year int) ([]dto.AcademicCalendarItem, error) {
	var results []dto.AcademicCalendarItem

	// 1. Get Monthly Events
	events, _, err := s.monthlyEventRepo.GetAll(ctx, nil, meta.Meta{Limit: 1000}, "")
	if err == nil {
		for _, e := range events {
			// Filter by month/year if provided
			if month > 0 && year > 0 {
				if int(e.Month.Month()) != month || e.Month.Year() != year {
					continue
				}
			}

			results = append(results, dto.AcademicCalendarItem{
				Id:          e.Id,
				Title:       e.Title,
				Type:        "monthly_event",
				Date:        e.Month,
				Description: e.Description,
				Link:        e.Link,
			})
		}
	}

	// 2. Get Timelines
	timelines, err := s.timelineRepo.GetAll(ctx, nil)
	if err == nil {
		for _, t := range timelines {
			// Filter by month/year if provided
			if month > 0 && year > 0 {
				if int(t.Date.Month()) != month || t.Date.Year() != year {
					continue
				}
			}

			progendaName := ""
			if t.Progenda != nil {
				progendaName = t.Progenda.Name
			}
			results = append(results, dto.AcademicCalendarItem{
				Id:           t.Id,
				Title:        t.Info,
				Type:         "timeline",
				Date:         t.Date,
				Link:         t.Link,
				ProgendaId:   t.ProgendaId,
				ProgendaName: progendaName,
			})
		}
	}

	// 3. Sort by Date
	sort.Slice(results, func(i, j int) bool {
		return results[i].Date.Before(results[j].Date)
	})

	return results, nil
}
