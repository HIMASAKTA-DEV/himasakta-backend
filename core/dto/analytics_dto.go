package dto

import "time"

type VisitorVisitRequest struct {
	VisitorId string `json:"visitor_id" binding:"required,uuid"`
}

type AnalyticsStatsResponse struct {
	VisitorCount            int64                `json:"VisitorCount"`
	NewsCount               int64                `json:"NewsCount"`
	DepartmentCount         int64                `json:"DepartmentCount"`
	ActiveProgendaCount      int64                `json:"ActiveProgendaCount"`
	ActiveMonthlyEventCount int64                `json:"ActiveMonthlyEventCount"`
	ActiveAnggotaCount      int64                `json:"ActiveAnggotaCount"`
	ActiveGalleryCount      int64                `json:"ActiveGalleryCount"`
	ActiveNRPWhitelistCount int64                `json:"ActiveNRPWhitelistCount"`
	NewVisitorsGraph        []VisitorsGraphPoint `json:"NewVisitorsGraph"`
}

type VisitorsGraphPoint struct {
	Timestamp time.Time `json:"timestamp"`
	Count     int64     `json:"count"`
}
