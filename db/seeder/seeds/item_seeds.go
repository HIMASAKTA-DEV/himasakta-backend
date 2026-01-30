package seeds

import (
	"encoding/json"
	"os"

	"github.com/Flexoo-Academy/Golang-Template/internal/entity"
	mylog "github.com/Flexoo-Academy/Golang-Template/internal/pkg/logger"
	"gorm.io/gorm"
)

func SeederItem(db *gorm.DB) error {
	mylog.Infof("[PROCESS] Seeding items...")

	jsonFile, err := os.Open("./db/seeder/data/item_data.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	var listEntity []entity.Item
	if err := json.NewDecoder(jsonFile).Decode(&listEntity); err != nil {
		return err
	}

	for _, item := range listEntity {
		if err := db.Create(&item).Error; err != nil {
			mylog.Errorf("[ERROR] Failed to seed item %s: %v", item.CategoryCode, err)
			return err
		}
	}

	mylog.Infof("[COMPLETE] Seeding items completed")
	return nil
}

