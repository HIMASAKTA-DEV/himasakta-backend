package routes

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/gin-gonic/gin"
)

func AcademicCalendar(app *gin.Engine, c controller.AcademicCalendarController) {
	app.GET("/api/v1/kalender-akademik", c.GetCalendar)
}
