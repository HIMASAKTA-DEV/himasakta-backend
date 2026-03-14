package service

import (
	"context"
	"sync"
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"github.com/google/uuid"
)

type AnalyticsService interface {
	TrackVisit(ctx context.Context, visitorId string, clientIp string) (int, error)
	GetStats(ctx context.Context, graphLimit string) (dto.AnalyticsStatsResponse, error)
}

type analyticsService struct {
	repo       repository.AnalyticsRepository
	cache      dto.AnalyticsStatsResponse
	cacheTime  time.Time
	cacheMutex sync.RWMutex
}

func NewAnalyticsService(repo repository.AnalyticsRepository) AnalyticsService {
	return &analyticsService{repo: repo}
}

func (s *analyticsService) TrackVisit(ctx context.Context, visitorId string, clientIp string) (int, error) {
	uid, err := uuid.Parse(visitorId)
	if err != nil {
		return 400, err
	}

	// Check if visitor exists
	exists, _ := s.repo.CheckVisitorExists(ctx, uid)

	if !exists {
		// Rate limiting for new visitors
		count, err := s.repo.CountNewVisitorsByIp(ctx, clientIp, time.Now().Add(-1*time.Hour))
		if err == nil && count >= 500 {
			return 429, nil // Too Many Requests
		}
	}

	// Async upsert
	go func() {
		bgCtx := context.Background()
		visitor := entity.Visitor{
			Id:         uid,
			ClientIp:   clientIp,
			LastSeenAt: time.Now(),
		}
		if !exists {
			visitor.CreatedAt = time.Now()
			
			// Invalidate cache for new visitor
			s.cacheMutex.Lock()
			s.cacheTime = time.Time{} // Force refresh on next GetStats
			s.cacheMutex.Unlock()
		}
		if err := s.repo.UpsertVisitor(bgCtx, visitor); err != nil {
			mylog.Errorf("Failed to upsert visitor: %v", err)
		}
	}()

	return 202, nil // Accepted
}

func (s *analyticsService) GetStats(ctx context.Context, graphLimit string) (dto.AnalyticsStatsResponse, error) {
	s.cacheMutex.RLock()
	if time.Since(s.cacheTime) < 5*time.Minute {
		defer s.cacheMutex.RUnlock()
		return s.cache, nil
	}
	s.cacheMutex.RUnlock()

	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	// Re-check cache after acquiring lock
	if time.Since(s.cacheTime) < 5*time.Minute {
		return s.cache, nil
	}

	stats, err := s.repo.GetStats(ctx)
	if err != nil {
		return stats, err
	}

	limit := 7 * 24 * time.Hour
    if graphLimit != "" {
        parsed, err := time.ParseDuration(graphLimit)
        if err == nil {
            limit = parsed
        }
    }

	graph, err := s.repo.GetNewVisitorsGraph(ctx, limit)
	if err != nil {
		return stats, err
	}
	stats.NewVisitorsGraph = graph

	s.cache = stats
	s.cacheTime = time.Now()

	return stats, nil
}
