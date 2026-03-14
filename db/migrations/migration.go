package migrations

import (
	"fmt"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println(mylog.ColorizeInfo("\n=========== Start Migrate ==========="))

	mylog.Infof("Migrating Tables...")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	// migrate table
	if err := db.AutoMigrate(
		&entity.Gallery{},
		&entity.Department{},
		&entity.CabinetInfo{},
		&entity.Member{},
		&entity.Progenda{},
		&entity.MonthlyEvent{},
		&entity.News{},
		&entity.NrpWhitelist{},
		&entity.Timeline{},
		&entity.Role{},
		&entity.Visitor{},
		&entity.GlobalSetting{},
		&entity.Tag{},
		&entity.NewsTag{},
	); err != nil {
		return err
	}

	mylog.Infof("Migration completed successfully")

	return nil
}
