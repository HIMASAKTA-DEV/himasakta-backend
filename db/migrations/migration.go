package migrations

import (
	"log"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	//"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println(mylog.ColorizeInfo("\n=========== Start Migrate ==========="))

	mylog.Infof("Migrating Tables...")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	allEntities := []interface{}{
		&entity.Role{},
		&entity.Department{},
		&entity.CabinetInfo{},
		&entity.GlobalSetting{},
		&entity.NrpWhitelist{},
		&entity.Visitor{},
		&entity.MonthlyEvent{},
		&entity.Tag{},
		&entity.Gallery{},
		&entity.Progenda{},
		&entity.Timeline{},
		&entity.Member{},
		&entity.News{},
		&entity.NewsTag{},
	}

	// Phase 1: Create all tables without FK constraints (handles circular Gallery↔Progenda)
	db.Config.DisableForeignKeyConstraintWhenMigrating = true
	if err := db.AutoMigrate(allEntities...); err != nil {
		return err
	}

	// Phase 2: Re-run with FKs enabled to add foreign key constraints
	db.Config.DisableForeignKeyConstraintWhenMigrating = false
	if err := db.AutoMigrate(allEntities...); err != nil {
		return err
	}

	mylog.Infof("Migration completed successfully")

	return nil
}
