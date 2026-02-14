package migrations

import (
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	mylog.Infoln("\n=========== Start Migrate ===========")
	mylog.Infof("Migrating Tables...")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	//migrate table
	if err := db.AutoMigrate(
		&entity.Gallery{},
		&entity.Department{},
		&entity.CabinetInfo{},
		&entity.Member{},
		&entity.Progenda{},
		&entity.ProgendaTimeline{},
		&entity.MonthlyEvent{},
		&entity.News{},
		&entity.NrpWhitelist{},
	); err != nil {
		return err
	}

	return nil
}
