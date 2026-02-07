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

	//migrate table
	if err := db.AutoMigrate(
		// &entity.User{},
		// &entity.RefreshToken{},
		&entity.Task{},
	); err != nil {
		return err
	}

	// if err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email) WHERE deleted_at IS NULL;`).Error; err != nil {
	// 	return err
	// }

	return nil
}
