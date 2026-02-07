package middleware

import (
	"net/http"
	"time"

	myerror "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/error"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	LockApiMiddleware struct {
		IsLocked bool
		location *time.Location
		ctx      *gin.Context
	}

	LockOption func(m *LockApiMiddleware)
)

func (m Middleware) LockAPI(msg string, opts ...LockOption) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		location, _ := time.LoadLocation("Asia/Jakarta")
		lockApiMiddleware := LockApiMiddleware{
			IsLocked: false,
			location: location,
			ctx:      ctx,
		}

		if len(opts) == 0 {
			lockApiMiddleware.IsLocked = true
		}

		for _, opt := range opts {
			opt(&lockApiMiddleware)
		}

		// TODO: Ganti pesannya
		if lockApiMiddleware.IsLocked {
			response.NewFailed("Di luar masa pengumuman. Sabar ya!", myerror.New(msg, http.StatusForbidden)).SendWithAbort(ctx)
			return
		}

		ctx.Next()
	}
}

func (m Middleware) NotBefore(t string) LockOption {
	return func(m *LockApiMiddleware) {
		parsedTime, err := time.ParseInLocation("01-02-2006 15:04:05", t, m.location)
		if err != nil {
			return
		}

		now := time.Now().In(m.location)
		if now.Before(parsedTime) {
			m.IsLocked = true
		}
	}
}

func (m Middleware) NotAfter(t string) LockOption {
	return func(m *LockApiMiddleware) {
		parsedTime, err := time.ParseInLocation("01-02-2006 15:04:05", t, m.location)
		if err != nil {
			return
		}
		now := time.Now().In(m.location)
		if now.After(parsedTime) {
			m.IsLocked = true
		}
	}
}

func (m Middleware) NotInRange(start, end string) LockOption {
	return func(m *LockApiMiddleware) {
		startTime, err1 := time.ParseInLocation("01-02-2006 15:04:05", start, m.location)
		endTime, err2 := time.ParseInLocation("01-02-2006 15:04:05", end, m.location)

		if err1 != nil || err2 != nil {
			return
		}

		now := time.Now().In(m.location)
		if now.Before(startTime) || now.After(endTime) {
			m.IsLocked = true
		}
	}
}

// func (m Middleware) WithTimeEvent(eventId string) LockOption {
// 	return func(ml *LockApiMiddleware) {
// 		var event entity.Event
// 		if err := m.db.Where("id = ?", eventId).Take(&event).Error; err != nil {
// 			logger.Errorln("invalid id or event with this id not found")
// 			return
// 		}

// 		startTime, err := utils.ParseDateEvent("2006-01-02 15:04:05", event.StartDate, ml.location)
// 		if err != nil {
// 			logger.Errorln(err)
// 			return
// 		}
// 		endTime, err := utils.ParseDateEvent("2006-01-02 15:04:05", event.EndDate, ml.location)
// 		if err != nil {
// 			logger.Errorln(err)
// 			return
// 		}

// 		now := time.Now().In(ml.location)
// 		if now.Before(startTime) || now.After(endTime) {
// 			ml.IsLocked = true
// 		}
// 	}
// }

// func (m Middleware) WithActiveEvent(eventId string) LockOption {
// 	return func(ml *LockApiMiddleware) {
// 		var event entity.Event

// 		if err := m.db.Where("id = ?", eventId).Take(&event).Error; err != nil {
// 			logger.Errorln("invalid id or event with this id not found")
// 			return
// 		}

// 		if !event.IsActive {
// 			ml.IsLocked = true
// 		}
// 	}
// }

