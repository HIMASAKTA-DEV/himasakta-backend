package repository

import (
	"context"
	"time"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/dto"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AnalyticsRepository interface {
	UpsertVisitor(ctx context.Context, visitor entity.Visitor) error
	GetStats(ctx context.Context) (dto.AnalyticsStatsResponse, error)
	GetNewVisitorsGraph(ctx context.Context, limit time.Duration) ([]dto.VisitorsGraphPoint, error)
	CountNewVisitorsByIp(ctx context.Context, ip string, since time.Time) (int64, error)
	CheckVisitorExists(ctx context.Context, id uuid.UUID) (bool, error)
}

type analyticsRepository struct {
	db *gorm.DB
}

func NewAnalyticsRepository(db *gorm.DB) AnalyticsRepository {
	return &analyticsRepository{db: db}
}

func (r *analyticsRepository) UpsertVisitor(ctx context.Context, visitor entity.Visitor) error {
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"last_seen_at": visitor.LastSeenAt, "client_ip": visitor.ClientIp}),
	}).Create(&visitor).Error
}

func (r *analyticsRepository) GetStats(ctx context.Context) (dto.AnalyticsStatsResponse, error) {
	var stats dto.AnalyticsStatsResponse

	r.db.WithContext(ctx).Model(&entity.Visitor{}).Count(&stats.VisitorCount)
	r.db.WithContext(ctx).Model(&entity.News{}).Count(&stats.NewsCount)
	r.db.WithContext(ctx).Model(&entity.Department{}).Count(&stats.DepartmentCount)
	r.db.WithContext(ctx).Model(&entity.Progenda{}).Count(&stats.ActiveProgendaCount) // Assuming all are active for now
	r.db.WithContext(ctx).Model(&entity.MonthlyEvent{}).Count(&stats.ActiveMonthlyEventCount)
	r.db.WithContext(ctx).Model(&entity.Member{}).Count(&stats.ActiveAnggotaCount)

	return stats, nil
}

func (r *analyticsRepository) GetNewVisitorsGraph(ctx context.Context, limit time.Duration) ([]dto.VisitorsGraphPoint, error) {
	points := make([]dto.VisitorsGraphPoint, 0)
	since := time.Now().Add(-limit)

	err := r.db.WithContext(ctx).Model(&entity.Visitor{}).
		Select("date_trunc('hour', created_at) as timestamp, count(*) as count").
		Where("created_at >= ?", since).
		Group("timestamp").
		Order("timestamp").
		Scan(&points).Error

	return points, err
}

func (r *analyticsRepository) CountNewVisitorsByIp(ctx context.Context, ip string, since time.Time) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Visitor{}).
		Where("client_ip = ? AND created_at >= ?", ip, since).
		Count(&count).Error
	return count, err
}

func (r *analyticsRepository) CheckVisitorExists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	r.db.WithContext(ctx).Model(&entity.Visitor{}).Where("id = ?", id).Count(&count)
	return count > 0, nil
}
