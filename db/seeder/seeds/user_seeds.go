package seeds

import (
	"encoding/json"

	"os"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/entity"
	mylog "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/logger"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/utils"
	"gorm.io/gorm"
)

func SeederUser(db *gorm.DB) error {
	mylog.Infof("[PROCESS] Seeding users...")
	jsonFile, err := os.Open("./db/seeder/data/user_data.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	var listEntity []entity.User
	if err := json.NewDecoder(jsonFile).Decode(&listEntity); err != nil {
		return err
	}

	for _, entity := range listEntity {
		entity.Password, _ = utils.HashPassword(entity.Password)
		if err := db.Save(&entity).Error; err != nil {
			return err
		}
	}

	mylog.Infof("[COMPLETE] Seeding users completed")
	return nil
}

