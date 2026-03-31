package migrations

import (
	"encoding/json"
	"log"
	"os"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"golang.org/x/crypto/bcrypt"
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

	db.Config.DisableForeignKeyConstraintWhenMigrating = true
	if err := db.AutoMigrate(allEntities...); err != nil {
		return err
	}

	db.Config.DisableForeignKeyConstraintWhenMigrating = false
	if err := db.AutoMigrate(allEntities...); err != nil {
		return err
	}

	if err := seedAdmin(db); err != nil {
		mylog.Errorf("Failed to seed admin: %v", err)
	}

	mylog.Infof("Migration completed successfully")

	return nil
}

func seedAdmin(db *gorm.DB) error {
	var existing entity.GlobalSetting
	result := db.Where("key = ?", "auth").First(&existing)
	if result.Error == nil {
		mylog.Infof("Admin credentials already exist, skipping seed")
		return nil
	}

	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	if username == "" || password == "" {
		mylog.Infof("ADMIN_USERNAME or ADMIN_PASSWORD not set, skipping admin seed")
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	authJSON, _ := json.Marshal(map[string]string{
		"username": username,
		"password": string(hash),
	})

	setting := entity.GlobalSetting{
		Key:   "auth",
		Value: string(authJSON),
	}

	if err := db.Create(&setting).Error; err != nil {
		return err
	}

	mylog.Infof("Initial admin credentials seeded (user: %s)", username)
	return nil
}
